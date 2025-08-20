package main

import (
	"GoLang-Tutorials/simplecalc"
	"fmt"
)


func main() {
	fmt.Println("--- Basic Hello world code ---")
	fmt.Println("Hello World")

	a, b := 5, 3
	fmt.Println("--- Arithmetic Operations ---")
	fmt.Println(simplecalc.Add(a, b))

	c, d := 10, 5
	fmt.Println(simplecalc.Subb(c, d))

	e, f := 5, 3
	fmt.Println(simplecalc.Mul(e, f))

	g, h := 15, 5
	fmt.Println(simplecalc.Div(g, h))
	// --- Flow Control ---
    // --- For Loop ---
	fmt.Println("--- Flow Control Statements ---")
	fmt.Println("--- for loop ---")
 	for i:=0; i<5; i++ {
		fmt.Println(i)
	}


	// --- If-Else ---
	fmt.Println("--- If Else Statements ---")
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
	fmt.Println("--- Switch/Case Statements ---")
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
	fmt.Println("--- Pointers Concept ---")
	fmt.Println(j)
	fmt.Println(p)
	fmt.Println(*p)

	// --- defer ---
	defer fmt.Println("Sadala")
	defer fmt.Println("Sumanth")
	defer fmt.Println("Hello")
	defer fmt.Println("--- Defer Concept ---")

    struct_demo()
}


