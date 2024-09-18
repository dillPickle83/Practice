package main

import (
	"fmt"
)

func main() {
	// The below is the way to define a pointer variable. The new() function will store a free memory address to the pointer. Initializing with just "var p *int32" will store the pointer but will not allocate memory to that var and the new function is needed to allocate the free memory
	var p *int32 = new(int32)
	var i int32
	// The var p will have the location (mem address) of the memory address that has been allocated. *p is the actual place where the value is stored and the way to access the value is using *p which is called dereferencing the pointer
	fmt.Printf("The value of p points to: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
	// We can manipulate the value of the address like below
	*p = 10
	// We can get access to the mem address of a var using the & operator
	// We do need to specify that mem needs to be allocated using the new() function since then the program will fail when trying to use the pointer since no memory will be allocated for the the pointer
	var d *int32 = new(int32)
	d = &i
	*d = 65
	fmt.Printf("The value of d points to: %v\n", *d)
	fmt.Printf("The value of i is: %v\n", i)


	/*
	Slices consist of pointers to the values of the assigned values. So if we copy a slice into a different slice2, modify slice2's values, we will see that both the slice values will be changed.
	var slice = []slice{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)		// [1,2,4]
	fmt.Println(sliceCopy)		// [1,2,4]
	*/

	/* When we use arrays in a function, the function will create a new array and modify that. But doing so will increase the memory consumption like in the below example
	func main() {
		var thing1 = [5]float64{1,2,3,4,5}
		fmt.Printf("The memory allocation of thing1 is: %v", &thing1)
		var result [5]float64 = square(thing1)
		fmt.Printf("The result is: %v", result)
		fmt.Printf("The value of slice 1 is: %v", thing1)
	}

	func square(thing2 [5]float64) [5]float64{
		fmt.Printf("The memory allocation of thing2 is: %v", &thing2)
		for i := range thing2{
			thing2[i] = thing2[i]*thing2[i]
		}
		return thing2
	}
	
	When we print the values of thing1 and thing2, we see that the values of the memory addresses assigned are different, implying that they are different arrays copies of the other array
	In the below example, we are passing the address of thing1 and we're passing that to the function which is accepting the pointer and returns the pointer to the modified array. This will save a lot of memory
	*/
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("The memory allocation of thing1 is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of slice 1 is: %v\n", thing1)

}

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("The memory allocation of thing2 is: %p\n", thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}
