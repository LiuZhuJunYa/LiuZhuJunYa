package utils

import (
	"fmt"
	"io/ioutil"
)

//设计一个方法用于读取txt文本内容
func TxtReading(address string) string {
	data, err := ioutil.ReadFile(address)
	if err != nil {
		fmt.Println(err)
		return "TXT Reading Error..."
	}
	str := string(data)
	return str
}