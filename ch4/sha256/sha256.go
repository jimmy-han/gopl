// 数组
// * 如果数组大小为省略号,则以数组初始化的大小为准
// * 数组的类型包含数组的长度 var r [2]int, var q [3]int, 数组类型分别为[2]int、[3]int

// 1. 键值对数组
// 2. 简写声明数组 [9: -1] 声明长度为10的数组,除最后一个为-1,其余全部为默认值0
// 3. 数组的比较 (如果数组元素是可比较的,那么数组类型也是可比较的)
package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)


func main() {
	// 1. 键值对数组
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	// 2. 简写声明数组
	r := [...]int{9: -1}
	for i, v := range r {
		fmt.Printf("%d %d\n", i, v)
	}

	// 3. 数组的比较
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}