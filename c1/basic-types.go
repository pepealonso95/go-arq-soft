// 3

package main

import (
	"fmt"
	"math/cmplx"
)

// En go los tipos basicos son:
// bool (false por defecto)

// string ("" por defecto)

// int  int8  int16  int32  int64 (0 por defecto)
// uint uint8 uint16 uint32 uint64 uintptr

// byte  (alias de uint8)

// rune  (alias de int32)
//      representa un punto en codigo Unicode

// float32 float64

// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
