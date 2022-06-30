package main

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	//s := "hello"
	//for i, v := range s {
	//	fmt.Println(i, string(v))
	//}
	var res []int
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 3)
	res = append(res, 4)
	res = res[:len(res)-1]
	fmt.Println(res)
}
