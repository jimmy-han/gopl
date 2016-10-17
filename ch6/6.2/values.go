package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item")) // 获取第一个
	fmt.Println(m["item"])     // 直接访问map

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")
}
