package utils

import "math/big"

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

