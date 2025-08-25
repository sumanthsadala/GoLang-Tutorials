package main

import (
	"GoLang-Tutorials/simplecalc"
	"GoLang-Tutorials/mapsexamples"
	"fmt"
	"time"
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

	// --- maps ---
	products := mapsexamples.Appleproducts()
	fmt.Println("Price of Iphone:", products["Iphone"])

	products["Ipad"] = 499

	if val, ok := products["Airtag"]; ok{
		fmt.Println("Airtag is available", val)
	} else {
		fmt.Println("Airtag is not available")
	}

	for i, v := range products {
		fmt.Printf("key: %s, value: %d\n", i,v)
	}

    // deleting a key
	delete(products, "Airpods")



    fmt.Println("After deleting a key from the products")
	for i, v := range products {
		fmt.Printf("key: %s, value: %d\n", i,v)
	}

	x := 5
	fmt.Println("Before:", x)
	updateValue(&x)
	fmt.Println("After:", x)


    fmt.Println("--- Method Concept ---")
	tri := Triangle{5, 6}
	fmt.Println("Using method:", tri.areamethod())

	fmt.Println("--- function concept ---")
	fmt.Println("Using function:", areafunc(3.0,4.0))

	//--- go routines concept ---
	fmt.Println("--- Goroutines concept ---")
	go birthdayWishes("Happy Birthday Sumanth")
	go birthdayWishes("Feliz cumpleaños Sumanth")
	time.Sleep(50 * time.Millisecond)

	//--- Palindrome in go ---
	fmt.Println("--- Palindrome in go ---")
	nums := [] int {121, 220, 343, 556, 12321}
	for _, num := range nums {
		if isPalindrome(num){
			fmt.Println(num, "is a Palindrome")
		} else {
			fmt.Println(num, "is not a Palindrome")
		}
	}

	//--- Factorial in go ---
	fmt.Println("--- Factorial in go ---")
	fmt.Printf("Factorial of %d is %d:\n", 3, factorialNum(3))
	fmt.Printf("Factorial of %d is %d:\n", 5, factorialNum(5))

	// --- fibonacci series in go ---
	fmt.Println("--- First 10 Fibonacci Series ---")
	fibonacciSeries(10)

	//--- Wordcount in go ---
	fmt.Println("--- WordCount in go ---")
	words := []string{"go","hello","world","hello","go","hello"}
	wordCount := make(map[string]int)
	for i:=0; i<len(words); i++{
		wordCount[words[i]]++
	}
	for word, count := range wordCount{
		fmt.Printf("%s: %d\n", word, count)
	}

}

// function for Factorial
func factorialNum(n int) int {
	factNum := 1
	for i:=1; i<=n; i++ {
		factNum = factNum * i
	}
	return factNum
}

// function for Fibonacci Series
func fibonacciSeries(series int){
	a, b := 0,1
	for i := 0; i < series; i++{
		fmt.Println(a)
		a, b = b, a+b
	}
}
// function for Palindrome
func isPalindrome(num int) bool{
	numStr := fmt.Sprintf("%d", num)
	for i := 0; i < len(numStr)/2; i++ {
		if numStr[i] != numStr[len(numStr)-1-i]{
			return false
		}
	}
	return true
}


// function for using goroutines
func birthdayWishes(wishes string) {
	for i :=1; i <= 2; i++ {
		fmt.Println(wishes, i)
		time.Sleep(5 * time.Millisecond)
	}

}


func updateValue(val *int){
	*val = *val + 10
}
type Triangle struct {
	base float64
	height float64
}
func (t Triangle) areamethod() float64 {
	return 0.5 * t.base * t.height
}


func areafunc (b, h float64) float64 {
	return 0.5 * b * h
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
