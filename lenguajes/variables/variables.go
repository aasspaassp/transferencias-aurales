package main

import (
	"fmt"
	"math/cmplx"
)

const CONSTANTE = 666

var cadena, texto string = "go lang", "dos variables"

var numero int16

var decimal float64

var a = "initial"

var b, c int = 1, 2

var d = true

var e int

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	f := "apple"
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
}
