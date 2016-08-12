// 简单实现slice append函数功能
package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// 一次添加多个元素到slice
func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		 z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * zlen {
			 zcap = 2 * zlen
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	fmt.Printf("cap=%d\t%v\n", cap(z), z)
	return z
}

func main() {
	var x []int
	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(x), x)
	}

	var y []int
	y = appendInts(y, 1, 2)
	y = appendInts(y, 3, 4, 5)
	y = appendInts(y, 6, 7, 8)
	y = appendInts(y, 9, 0)
	fmt.Println(y)
}
