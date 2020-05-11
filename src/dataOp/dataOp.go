package dataOp

import (
)

func InArray(data map[int]string, s string) (bool) {

	for i := 0; i < len(data); i++ {
		if data[i] == s {
			return (true)
		}
	}
	return (false)
}

func GetMat(cat map[int]string, mat map[int]string) {

	for i := 6; i < len(cat); i++ {
		mat[len(mat)] = cat[i]
	}
}