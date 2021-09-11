package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

//设计一个方法用于输出txt文本
func TxtWriting(address string, sim float64) {
	str := strconv.FormatFloat(sim, 'f', -1, 64)
	err := ioutil.WriteFile(address, []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
