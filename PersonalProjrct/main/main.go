package main

import (
	"PersonalProjrct/utils"
	"fmt"
)

//在goland中测试的主函数
func main() {

	//创建两个变量str1、str2用来存放读取的文本
	address := ""
	fmt.Println("请输入第一个文本所在地址：")
	fmt.Scanln(&address)
	str1 := utils.TxtReading(address)
	fmt.Println("请输入第二个文本所在地址：")
	fmt.Scanln(&address)
	str2 := utils.TxtReading(address)


}

//在cmd命令行中测试的主函数
func cmd(address1 , address2 string) {

}