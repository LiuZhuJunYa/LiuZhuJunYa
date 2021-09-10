package utils

import (
	"math/big"
)

//抽象定义一个海明哈希类型
type SimHash struct {
	IntSimHash *big.Int
	HashBits   int
}

//初始化SimHash
func Params() (s *SimHash) {
	s = &SimHash{}
	s.HashBits = 128
	return s
}

