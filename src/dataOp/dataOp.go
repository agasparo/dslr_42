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