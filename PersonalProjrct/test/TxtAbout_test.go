package test

import (
	"PersonalProjrct/src/utils"
	"fmt"
	"testing"
)

func TestTxtReading(t *testing.T) {

	address := ""

	//测试正确的输入TxtReading
	fmt.Println("请输入第一个文本所在地址：")
	address = "C:\\Users\\Administrator\\Desktop\\测试文本\\orig.txt"
	_ = utils.TxtReading(address)

	//测试正确的输出TxtWriting
	utils.TxtWriting("C:\\Users\\Administrator\\Desktop\\测试文本\\ans.txt", 0.89)

	//测试错误的输入TxtReading
	fmt.Println("请输入第一个文本所在地址：")
	address = "C:\\Users\\Administrator\\Desktop\\测试文本\\orig1.txt"
	_ = utils.TxtReading(address)

	//测试错误的输出TxtWriting
	utils.TxtWriting("C:\\Users\\Administrator\\Desktop\\测试文本", 0.89)
}