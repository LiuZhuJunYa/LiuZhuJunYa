package main

import (
	"PersonalProjrct/src/utils"
	"fmt"
	"os"
	"runtime/pprof"
)

//在goland中测试的主函数
func main() {
	cpuProfile, _ := os.Create("cpu_profile.prof")
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	//创建两个变量str1、str2用来存放读取的文本
	address := ""
	fmt.Println("请输入第一个文本所在地址：")
	address = "C:\\Users\\66389\\Desktop\\测试文本\\orig.txt"
	str1 := utils.TxtReading(address)
	fmt.Println("请输入第二个文本所在地址：")
	address = "C:\\Users\\66389\\Desktop\\测试文本\\orig_0.8_add.txt"
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