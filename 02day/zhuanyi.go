package main

import (
	"fmt"
	"unsafe"
)

func main() {
    num := 5.1234e2
	fmt.Println("num=", num)
	fmt.Println("num的数据类型是", unsafe.Sizeof(num)
}
