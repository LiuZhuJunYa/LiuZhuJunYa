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

	//计算simhash
	s := utils.Params()
	hash1 := s.Simhash(str1)
	fmt.Println(hash1)
	hash2 := s.Simhash(str2)
	fmt.Println(hash2)

	////距离
	distance := s.HammingDistance(hash1, hash2)
	fmt.Println(distance)
	////计算相似度
	similarity := s.Similarity(hash1, hash2)
	fmt.Println(similarity)
}

//在cmd命令行中测试的主函数
//func cmd(address1 , address2 string) {
//
//}