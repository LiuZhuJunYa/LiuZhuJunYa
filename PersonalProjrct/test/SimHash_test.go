package test

import (
	"PersonalProjrct/src/utils"
	"fmt"
	"testing"
)

func TestSimHash(t *testing.T) {

	//对simhash算法的无文本输入测试
	str := "张老师真帅，请多给我一点点分，可以吗?"
	s := utils.Params()
	//正确输入测试
	hash := s.Simhash(str)
	fmt.Println(hash)
	//异常输入测试
	hash = s.Simhash("")
	fmt.Println(hash)

	//对simhash算法的文本输入测试
	address1 := "C:\\Users\\Administrator\\Desktop\\测试文本\\orig.txt"
	address2 := "C:\\Users\\Administrator\\Desktop\\测试文本\\ans.txt"
	str = utils.TxtReading(address1)
	hash = s.Simhash(str)
	fmt.Println(hash)
	utils.TxtWriting(address2, float64(hash.Int64()))
}