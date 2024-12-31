package utils

import "math/rand"

func GenRandSn() int64 {
	min := 100000
	max := 999999
	userSn := int64(rand.Intn(max-min+1) + min)
	return userSn
}
