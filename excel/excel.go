package excel

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

// 将数据转换为excel文件的数据
func DataToExcelFileData(columnNames []string, v interface{}, sheetName string) (data *bytes.Buffer, err error) {
	if v == nil {
		return nil, fmt.Errorf("导出数据不能为空")
	}
	// 判断类型
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Slice {
		return nil, fmt.Errorf("导出数据必须是数组")
	}
	// 判断长度
	vValue := reflect.ValueOf(v)
	vl := vValue.Len()
	logrus.Debugf("cap :%d", vl)
	xlsx := excelize.NewFile()
	sheetIndex := xlsx.NewSheet(sheetName)
	err = AddExcelHeader(xlsx, sheetName, columnNames)
	if err != nil {
		return
	}
	for i := 0; i < vl; i++ {
		childValue := vValue.Index(i)
		if reflect.TypeOf(childValue.Interface()).Kind() == reflect.Struct {
			numField := childValue.NumField()
			for j := 0; j < numField; j++ {
				err = AddExcelValue(xlsx, sheetName, j, i+2, childValue.Field(j).Interface())
				if err != nil {
					return
				}
			}
		} else {
			err = AddExcelValue(xlsx, sheetName, 0, i+2, childValue.Interface())
			if err != nil {
				return
			}
		}
	}
	xlsx.SetActiveSheet(sheetIndex)
	// Save xlsx file by the given path.
	if err := xlsx.SaveAs("Book1.xlsx"); err != nil {
		logrus.Errorf("save file err:%v", err)
		return nil, err
	}
	data, err = xlsx.WriteToBuffer()
	return
}

func AddExcelHeader(xlsx *excelize.File, sheetName string, columnNames []string) (err error) {
	for index, value := range columnNames {
		err = xlsx.SetCellValue(sheetName, GenAxisName(index, 1), value)
		if err != nil {
			return
		}
	}
	return
}

func AddExcelValue(xlsx *excelize.File, sheetName string, index int, rowNumber int, v interface{}) (err error) {
	err = xlsx.SetCellValue(sheetName, GenAxisName(index, rowNumber), v)
	return err
}

const NUMBERLEN = 26

var columnNumbers = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func GenAxisName(index int, rowNumber int) string {
	if index > NUMBERLEN-1 {
		count := index / NUMBERLEN
		newIndex := index % NUMBERLEN
		return fmt.Sprintf("%s%d", strings.Repeat(columnNumbers[newIndex], count+1), rowNumber)
	} else {
		return fmt.Sprintf("%s%d", columnNumbers[index], rowNumber)
	}
}
