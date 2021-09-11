package utils

import "math/big"

//字符串哈希算法(md5)
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

//计算simhash的距离
func (s *SimHash) HammingDistance(hash, other *big.Int) *big.Int {
	hase_v := new(big.Int)
	c_w := hase_v.Xor(hash, other)
	fbIng := big.NewInt(1)

	fbIng.Lsh(fbIng, uint(s.HashBits))
	bit_big := new(big.Int)
	t_mak := bit_big.Sub(fbIng, big.NewInt(1))
	c_result := new(big.Int)
	c_result.And(c_w, t_mak)

	tot := big.NewInt(0)
	for c_result.Cmp(big.NewInt(0)) > 0 {
		tot.Add(tot, big.NewInt(1))
		ts := new(big.Int)
		ts.Sub(c_result, big.NewInt(1))
		c_result.And(c_result, ts)
	}
	return tot
}

//相似度计算
func (s *SimHash) Similarity(hash, other *big.Int) float64 {
	a := new(big.Rat)
	a.SetInt(hash)
	b := new(big.Rat)
	b.SetInt(other)
	val := new(big.Rat)
	if a.Cmp(b) > 0 {
		val.Quo(b, a)
		f, _ := val.Float64()
		return f
	}
	val.Quo(a, b)
	f, _ := val.Float64()
	return f
}

