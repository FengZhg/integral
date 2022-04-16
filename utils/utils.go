package utils

import (
	"hash/fnv"
)

// @Author: Feng
// @Date: 2022/4/1 14:12

//GetDBIndex 根据
func GetDBIndex(id string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return h.Sum32() % 1000
}
