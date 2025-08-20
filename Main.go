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
	fmt.Println("Addition of two numbers a and b:", simplecalc.Add(a, b))

	c, d := 10, 5
	fmt.Println("Subtraction of two numbers c and d:", simplecalc.Subb(c, d))

	e, f := 5, 3
	fmt.Println("Multiplication of two numbers e and f:", simplecalc.Mul(e, f))

	g, h := 15, 5
	fmt.Println("Division of two numbers g and h:", simplecalc.Div(g, h))
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
	fmt.Println("Value of j:", j)
	fmt.Println("Address of j through pointer p:", p)
	fmt.Println("Value of pointer p ", *p)

	// --- defer ---
	defer fmt.Println("Sadala")
	defer fmt.Println("Sumanth")
	defer fmt.Println("Hello")
	defer fmt.Println("--- Defer Concept ---")


	// Declaring the func struct_demo
	struct_demo()

    
}

type Person struct {
	Name string
	Age int
}

type Address struct {
	Address1 string
	City string
	State string
	Zipcode int
}

func struct_demo() {
	p1 := Person {
		Name: "Sumanth",
		Age: 24,
	}
	fmt.Println("--- Structs Concept ---")
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Name and Age:", p1)

	a1 := Address {
		Address1: "1525 Amherst dr",
		City: "Rochester",
		State: "NY",
		Zipcode: 14221,
	}
	fmt.Println("Address of the person:", a1)

	// --- Arrays ---
	var a [5] string
	fmt.Println("--- Arrays Concept ---")
	fmt.Println("Array before adding the values:", a)
	a[0] = "Go"
	a[1] = "hello"
	a[2] = "Sumanth"
	a[3] = "Sadala"
	a[4] = "Bye"
	fmt.Println("Array after adding the values:", a)

	scores := [5] int {65,75,85}
	fmt.Println("Batsmens scores before:", scores)
	scores[3] = 70
	scores[4] = 95
	fmt.Println("Batsmens scores after two more innings:", scores)

	// -- for in arrays
	for i, v := range scores {
		fmt.Printf("Innings: %d, scores: %d\n", i, v)
	}

    fmt.Println("--- Batsmens scores for three innings ---")
	for i, v := range scores {
		if i <3 {
			fmt.Printf("Innings: %d, scores: %d\n", i, v)
		}
	}

	// --- Slices --- 
	fmt.Println("--- Slices concept ---")
	s := [] string {"GO", "Hello", "World"}
	fmt.Println("Length before:", len(s))
	fmt.Println("Capacity Before", cap(s))
	fmt.Println("Slice before:", s)
	s = append(s, "Good")
	s = append(s, "bye")
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)

	s1 := make([] int, 3)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)
	s1 = append(s1, 4)
	s1 = append(s1, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)

	s2 := make([] int, 3, 5)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	s2 = append(s2, 6)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)


	// Slice using array 
	arr := [5] int {1,2,3,4,5}
	slice := arr[1:4]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("Slice for the index 1 to 4:", slice)
	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println("Slice after add two more values:", slice)
	fmt.Println("Lenght after add two more values:", len(slice))
	fmt.Println("Capacity after add two more values:", cap(slice))


}
