package utils

import "math/big"

//字符串哈希算法
func (s *SimHash) Hash(token string) *big.Int {
	if token == "" {                                          //假如在划分字符串的时候没有得到内容，马上结束
		return big.NewInt(0)
	} else {
		bigIntToken := big.NewInt(int64(token[0]))
		funit := new(big.Int)
		x := funit.Lsh(bigIntToken, 7)
		m := big.NewInt(1000003)
		mslB := big.NewInt(1)
		mask := mslB.Lsh(mslB, uint(s.HashBits))
		tsk_b := mask.Sub(mask, big.NewInt(1))
		for i := 0; i < len(token); i++ {
			tokens := big.NewInt(int64(token[i]))
			x.Mul(x, m)
			x.Xor(x, tokens)
			x.And(x, tsk_b)
		}
		x = x.Xor(x, big.NewInt(int64(len(token))))
		if x == big.NewInt(-1) {
			x = big.NewInt(-2)
		}
		return x
	}
}

