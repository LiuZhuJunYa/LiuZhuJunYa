package utils

import (
	"math/big"
	"strings"
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

//核心算法海明距离的实现
func (s *SimHash) Simhash(str string) *big.Int {
	//核心方法之一，对传入文本进行分割，选择的参数会影响到simhash的计算
	m := strings.Split(str, "，")

	token_int := make([]int, s.HashBits)
	for i := 0; i < len(m); i++ {
		temp := m[i]
		//获取其hash值
		t := s.Hash(temp)
		//加权合并
		for j := 0; j < s.HashBits; j++ {
			fbIng := big.NewInt(1)
			bitMask := fbIng.Lsh(fbIng, uint(j))
			statusBig := new(big.Int)
			statusBig.And(t, bitMask)
			if statusBig.Cmp(big.NewInt(0)) != 0 {
				token_int[j] += 1
			} else {
				token_int[j] -= 1
			}
		}

	}
	fingerprint := big.NewInt(0)
	//降维
	for i := 0; i < s.HashBits; i++ {
		if token_int[i] >= 0 {
			oneBig := big.NewInt(1)
			tbig := big.NewInt(0)
			fingerprint.Add(fingerprint, tbig.Lsh(oneBig, uint(i)))
		}
	}
	return fingerprint
}
