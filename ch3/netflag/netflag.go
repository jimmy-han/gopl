package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLookback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	fmt.Printf("FlagUp = %b\n", FlagUp)
	fmt.Printf("FlagUp = %b\n", FlagBroadcast)
	fmt.Printf("FlagLookback = %b\n", FlagLookback)
	fmt.Printf("FlagPointToPoint = %b\n", FlagPointToPoint)
	fmt.Printf("FlagMulticast = %b\n", FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}