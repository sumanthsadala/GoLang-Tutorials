package main

import (
	"GoLang-Tutorials/simplecalc"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	a, b := 5, 3
	fmt.Println(simplecalc.Add(a, b))

	c, d := 10, 5
	fmt.Println(simplecalc.Subb(c, d))

	e, f := 5, 3
	fmt.Println(simplecalc.Mul(e, f))

	g, h := 15, 5
	fmt.Println(simplecalc.Div(g, h))
	// --- Flow Control ---
    // --- For Loop ---
 	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	// --- While Loop ---
	// i := 0
	// for i < 5{
		// fmt.Println(i)
		// break
	// }


	// --- If-Else ---
	num := 5
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}



	Num := 35
	if Num > 35 {
		fmt.Println("Num is greater than 35")
	} else if Num == 35 {
		fmt.Println("Num is 35")
	} else {
		fmt.Println("Num is less than 35")
	}

	// --- Switch/Case ---
	day := "wed"
	switch day {
	case "mon":
		fmt.Println("monday")
	case "tue":
		fmt.Println("tuesday")
	case "wed":
		fmt.Println("wednesday")
		fallthrough
	default:
		fmt.Println("weekend")
	}


	// --- Pointers ---
	j := 25
	p := &j
	fmt.Println(j)
	fmt.Println(p)
	fmt.Println(*p)

	// --- defer ---
	defer fmt.Println("Sadala")
	fmt.Println("Hello")
	defer fmt.Println("Sumanth")



}


