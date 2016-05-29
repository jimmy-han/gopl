// Package tempconv performs Celsius and Fahrenheit temperature computations.

package main

import "fmt"

type Celsius float64        // 大写字母开头属全局,别的包也可使用
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC)
	bolingF := CToF(BoilingC)
	fmt.Printf("%g\n", bolingF - CToF(FreezingC))
	//fmt.Printf("%g\n", bolingF - FreezingC)         // 编译错误:类型不匹配

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) // 编译错误:类型不匹配
	fmt.Println(c == Celsius(f))    // 返回true,因为c 和 f 都为0

	sc := FToC(212.0)
	fmt.Println(sc.String())
	fmt.Printf("%v\n", sc)      // 不需要调用String()
	fmt.Printf("%s\n", sc)
	fmt.Println(sc)
	fmt.Printf("%g\n", sc)      // 不需要调用String()
	fmt.Println(float64(sc))    // 不需要调用String()
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}