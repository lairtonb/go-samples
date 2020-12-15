package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type inference
	v := 42.0 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Consts
	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("Hello", World)
}

/*
Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

/*
The int, uint, and uintptr types are usually 32 bits wide on 32-bit
systems and 64 bits wide on 64-bit systems. When you need an integer
value you should use int unless you have a specific reason to use a
sized or unsigned integer type.

But why?
*/

/*
Zero values
Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/
