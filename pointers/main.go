package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    if false {
        interegers();
        structs();
        slices();
    } else {
        useArguments();
    }
}

func useArguments() {
	var thing = [5]float64{1, 2, 3, 4, 5}

	for i := range thing {
        // note: %g prints numbers in a more human-readable format compared to the default scientific notation.
		fmt.Printf("%g\n", thing[i])
	}

	println("=======================")

    // note: call the arguments function and pass the pointer to the "thing" array.
    // by passing a pointer we avoid copying the whole array and end up saving memory
    // the function modifies the array in place (squaring each element) and returns the updated array
	var newThing = arguments(&thing)

	for i := range newThing {
		fmt.Printf("%g\n", newThing[i])
	}
}

// note: arguments receives a pointer to an array of 5 float64 numbers
// it modifies each element of the array in place by squaring its value
// since we're passing a pointer, the original array is accessed directly without a full copy
// this is memory efficient because we avoid duplicating the array data
func arguments(thing *[5]float64) [5]float64 {
	for i := range thing {
		// Square each element of the array.
		thing[i] = thing[i] * thing[i]
	}

    // note: return the array with modified values
	// while the array is copied for the return value
	// the in-place modification using the pointer ensures that the original array was updated.
	return *thing
}

func interegers() {
    // note: allocate memory for an int32 using new. p now points to that memory location
    // at this point, *p is initialized to zero.
    var p *int32 = new(int32)

    // note: declare an int32 variable i; since it is not explicitly initialized
    // it defaults to zero.
	var i int32

    // note: print the pointer value stored in p (which is a memory address)
    // and the initial value of i (which is 0).
    fmt.Printf("The value p points to is: %v", p)
    fmt.Printf("\nThe value of i is: %v", i)

    // note: make p point to the variable i by taking its address
    p = &i
    // note: dereference p and assign 10 to the memory location it points to
    // note: this operation effectively sets i to 10.
    *p = 10

    fmt.Printf("\nThe value p points to is: %v", p)
    fmt.Printf("\nThe value of i is: %v", i)

    // note: declare a new int32 variable k and initialize to 2
    var k int32 = 2
    // note: assign the value of k to i, overwriting the previous value of 10
    i = k
}

func structs() {
    // note: allocate memory for a person using new
    // new(Person) returns a pointer to a Person value initialized with zero values
    // here Name will be an empty string and Age will be 0
    p := new(Person)

    // note: later on in the code when needed you can set the fields
    p.Name = "Alice"
    p.Age = 30

    fmt.Println("Allocated using new:", p)

    // note: we can allocate a person using a struct literal
    // Using the & operator creates and returns a pointer to the Person struct
    // and you can initialize fields immediately
    q := &Person{
        Name: "Bob",
        Age:  25,
    }

    fmt.Println("Allocated using struct literal:", q)
}

func slices() {
    // note: create a slice 'data' with 5 elements
    data := []int{1, 2, 3, 4, 5}

    // note: create two slices from the same underlying array
    // s1 references elements from index 1 to 3 (2, 3, 4)
    s1 := data[1:4]
    // note: s2 references elements from index 2 to 4 (3, 4, 5)
    s2 := data[2:5]

    fmt.Println("Initial data:", data)
    fmt.Println("s1:", s1)
    fmt.Println("s2:", s2)

    // note: modify an element via s1 (change the second element of s1)
    // since s1 and s2 share the same underlying array, this modification affects both slices.
    s1[1] = 42

    fmt.Println("\nAfter modifying s1:")
    fmt.Println("data:", data)
    fmt.Println("s1:", s1)
    fmt.Println("s2:", s2)
}
