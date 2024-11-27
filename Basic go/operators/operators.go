package main

import "fmt"

func main() {
	// Arithmetic operators: +, -, *, /, %, ++, --
	var a, b int = 10, 20

	fmt.Printf("a + b = %d\n", a + b) // 30
	fmt.Printf("a - b = %d\n", a - b) // -10
	fmt.Printf("a * b = %d\n", a * b) // 200
	fmt.Printf("a / b = %d\n", a / b) // 0
	fmt.Printf("a %% b = %d\n", a % b) // 10
	a++
	fmt.Printf("a++ = %d\n", a) // 11
	a--
	fmt.Printf("a-- = %d\n", a) // 10
	fmt.Print("\n")



	// Relational operators: ==, !=, >, <, >=, <=
	fmt.Printf("a == b = %v\n", a == b) // false
	fmt.Printf("a != b = %v\n", a != b) // true
	fmt.Printf("a > b = %t\n", a > b) // false
	fmt.Printf("a < b = %t\n", a < b) // true
	fmt.Printf("a >= b = %t\n", a >= b) // false
	fmt.Printf("a <= b = %t\n", a <= b) // true
	fmt.Print("\n")



	// Logical operators: &&, ||, !
	var c, d bool = true, false

	fmt.Printf("c && d = %t\n", c && d) // false
	fmt.Printf("c || d = %t\n", c || d) // true
	fmt.Printf("!c = %t\n", !c) // false
	fmt.Print("\n")



	// Bitwise operators: &, |, ^, <<, >>
	var e, f int = 60, 13

	fmt.Printf("e & f = %d\n", e & f) // 12
	fmt.Printf("e | f = %d\n", e | f) // 61
	fmt.Printf("e ^ f = %d\n", e ^ f) // 49
	fmt.Printf("e << 2 = %d\n", e << 2) // 240
	fmt.Printf("e >> 2 = %d\n", e >> 2) // 15
	fmt.Print("\n")



	// Assignment operators: =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=
	var g int = 10

	g += 5
	fmt.Printf("g += 5 = %d\n", g) // 15

	g -= 5
	fmt.Printf("g -= 5 = %d\n", g) // 10

	g *= 5
	fmt.Printf("g *= 5 = %d\n", g) // 50

	g /= 5
	fmt.Printf("g /= 5 = %d\n", g) // 10

	g %= 5
	fmt.Printf("g %%= 5 = %d\n", g) // 0

	g <<= 2
	fmt.Printf("g <<= 2 = %d\n", g) // 0

	g >>= 2
	fmt.Printf("g >>= 2 = %d\n", g) // 0

	g &= 5
	fmt.Printf("g &= 5 = %d\n", g) // 0

	g ^= 5
	fmt.Printf("g ^= 5 = %d\n", g) // 5

	g |= 5
	fmt.Printf("g |= 5 = %d\n", g) // 5
}