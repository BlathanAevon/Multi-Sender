package tools

import "math/big"

func FloatToWei(num float64) *big.Int {
	bigFloat := new(big.Float).SetFloat64(num)

	wei := new(big.Float).Mul(bigFloat, new(big.Float).SetFloat64(1e18))

	weiInt, _ := wei.Int(nil)

	return weiInt
}
