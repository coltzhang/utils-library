package http_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/utils-library/http_request/types"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	method types.RequestMethod
	url    string
	params url.Values
	header map[string][]string
	body   io.Reader
}

func MakeRequest(url string, method types.RequestMethod) *Request {
	return &Request{
		url:    url,
		method: method,
		params: make(map[string][]string),
		header: make(map[string][]string),
		body:   nil,
	}
}

func (r *Request) AddHeader(key, val string) {
	if _, ok := r.header[key]; !ok {
		r.header[key] = make([]string, 0)
	}
	r.header[key] = append(r.header[key], val)
}

func (r *Request) SetHeader(key string, val []string) {
	r.header[key] = val
}

func (r *Request) AddParam(key, val string) {
	r.params.Set(key, val)
}

// form表单数据body参数传入字符串
func (r *Request) SetBody(body interface{}, bodyType types.RequestBodyType) {
	if body == nil {
		return
	}
	if bodyType == types.REQUEST_BODY_TYPE__JSON {
		r.AddHeader("Content-Type", "application/json")
		data, _ := json.Marshal(body)
		r.body = bytes.NewReader(data)
	} else if bodyType == types.REQUEST_BODY_TYPE__FORM {
		r.AddHeader("Content-Type", "application/x-www-form-urlencoded")
		r.body = strings.NewReader(body.(string))
	}
}

func (r *Request) Do() (resp []byte, err error) {
	// 处理链接、参数
	method := r.method.String()
	reqUrl := r.makeUrl()
	// 创建请求
	var req *http.Request
	if r.body == nil {
		req, err = http.NewRequest(method, reqUrl, nil)
	} else {
		req, err = http.NewRequest(method, reqUrl, r.body)
	}
	if err != nil {
		logrus.Errorf("创建请求失败:%v, method:%s, url:%s, reqBody:%+v", err, method,
			reqUrl, r.body)
		return nil, fmt.Errorf("创建请求失败")
	}
	// 写入header
	r.makeHeader(req)
	// 请求
	now := time.Now()
	client := &http.Client{}
	result, err := client.Do(req)
	if err != nil {
		logrus.Errorf("请求失败:%v, method:%s, url:%s, reqBody:%+v, header:%+v", err,
			method, reqUrl, r.body, r.header)
		return nil, fmt.Errorf("请求失败")
	}
	defer func() {
		err = result.Body.Close()
		if err != nil {
			logrus.Errorf("关闭请求失败:%v, method:%s, url:%s, reqBody:%+v, header:%+v",
				err, method, reqUrl, r.body, r.header)
		}
	}()
	// 读取返回内容
	resp, err = ioutil.ReadAll(result.Body)
	if err != nil {
		logrus.Errorf("读取请求结果失败:%v, method:%s, url:%s, reqBody:%+v, header:%+v",
			err, method, reqUrl, r.body, r.header)
		return nil, fmt.Errorf("读取请求结果失败")
	}

	logrus.Infof("Method: %s, Url: %s, Header: %+v, Body: %+v, Resp: %s, Cost: %v", method, reqUrl,
		r.header, r.body, string(resp), time.Since(now))
	return
}

func (r *Request) makeUrl() string {
	if len(r.params) == 0 {
		return r.url
	}
	u, _ := url.Parse(r.url)
	u.RawQuery = r.params.Encode()
	return u.String()
}

func (r *Request) makeHeader(req *http.Request) {
	if len(r.header) == 0 {
		return
	}
	for key, vals := range r.header {
		for i := range vals {
			req.Header.Add(key, vals[i])
		}
	}
}
