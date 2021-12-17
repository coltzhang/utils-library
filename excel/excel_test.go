package excel

import (
	"testing"
)

type UserInfo struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Gender string `json:"gender"`
}

func TestExcel(t *testing.T) {
	header := []string{"id", "name", "age", "gender"}
	rows := make([]UserInfo, 0)
	rows = append(rows, UserInfo{
		ID:     1,
		Name:   "zhangsan",
		Age:    27,
		Gender: "male",
	})
	rows = append(rows, UserInfo{
		ID:     2,
		Name:   "lisi",
		Age:    27,
		Gender: "famale",
	})
	_, err := DataToExcelFileData(header, rows, "Sheet1")
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	// http附件返回
	//attachment := httpx.NewAttachment("投保单管理.xlsx", "application/octet-stream")
	//_, err = attachment.Write(data)
}
