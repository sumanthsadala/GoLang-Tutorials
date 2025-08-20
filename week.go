package main

import "fmt"

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
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1)

	a1 := Address {
		Address1: "1525 Amherst dr",
		City: "Rochester",
		State: "NY",
		Zipcode: 14221,
	}
	fmt.Println(a1)

	// --- Arrays ---
	var a [5] string
	fmt.Println("--- Arrays Concept ---")
	fmt.Println(a)
	a[0] = "Go"
	a[1] = "hello"
	a[2] = "Sumanth"
	a[3] = "Sadala"
	a[4] = "Bye"
	fmt.Println(a)

	scores := [5] int {65,75,85}
	fmt.Println(scores)
	scores[3] = 70
	scores[4] = 95
	fmt.Println(scores)

	// -- for in arrays
	for i, v := range scores {
		fmt.Println(i, v)
	}

    fmt.Println("--- for loop with three iterations ---")
	for i, v := range scores {
		if i <3 {
			fmt.Println(i, v)
		}
	}

	// --- Slices --- 
	fmt.Println("--- Slices concept ---")
	s := [] string {"GO", "Hello", "World"}
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)
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
	fmt.Println(slice)

}