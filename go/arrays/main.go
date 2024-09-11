package main

import(
    "fmt"
    "time"
)

func main(){
    // Arrays are fixed length, indexable, with the same type elements and contiguous in memory
    // The array elements can be changed wrt index but the arr size can't change
    // The below snippet will init an array of length 3 with all elements of type int32; by default all the elements will be 0
    var intArr [3]int32
    intArr[1] = 123
    fmt.Println(intArr[0]) //Will print the first element
    fmt.Println(intArr[1:3]) // Will print elems in index 1 upto but not including 3

    // Since we mentioned that an arr of size 3 needs to be created with int32 elems, go will assign 12 bytes of memory; each int32 value is 4 bytes of memory
    // The memory address can be accessed using the & operator
    fmt.Printf("The address of the 2nd element of fthe array is %v", &intArr[1])
    // When trying to access the memory, it stores the value of the first element and accesses the subsequent values by incrementing by 4 [since int32 requires 4 bytes of memory]

    // Can also init the arr using the below syntaxes
    var intArr1 [3]int32 = [3]int32{1,2,3}
    fmt.Println(intArr1)

    // We can either specify the size of the arr the same way as before or give ...[spread operator]  so that the compiler will take the number of elements in the array assigned and use that as the length
    intArr2 := [...]int32{16,25,32}
    fmt.Println(intArr2)


    // Slices are used when you want to append new elements to an array; But since the arr size is already fixed, the slice will create a new array for the required size and then assign that
    var intSlice []int32 = []int32{213, 3453, 88}
    fmt.Printf("The size of the array %v is %v, with capacity %v\n", intSlice, len(intSlice), cap(intSlice))
    intSlice = append(intSlice, 56)
    fmt.Printf("After appending, the size of array %v is %v, with capacity %v\n", intSlice, len(intSlice), cap(intSlice))
    // The capacity of the new arr after append may be higher, but the indexes which are bigger than the capacity can't be accessed and will throw an index out of range error
    // [213, 3453, 88] became [213, 3453, 88, 56, *, *] after the append; Length=4, capacity=6; intSlice[4] will throw an err

    // Can append multiple values to the slice using the spread operator [...]
    intArr4 := []int32{18, 24}
    intSlice = append(intSlice, intArr4...)
    fmt.Printf("After appending multiple values: %v, length=%v, capacity=%v\n", intSlice, len(intSlice), cap(intSlice))

    // Another way to create a slice is using the make function [syntax: make(type, length, capacity(optional)]
    var intSlice3 []int32 = make([]int32, 3, 8)
    fmt.Println(intSlice3)


    // Maps are data structures that are stored as {"key": "value"} and value can be accessed using the key
    // Can be assigned using the make function or can be assigned directly with the values
    // When assigning the value, the type of the key and then the type of the value needs to be specified
    var defaultMap map[string]uint8 = make(map[string]uint8)
    fmt.Println(defaultMap)

    var myMap = map[string]uint8{"Adam": 23, "Cory": 43}
    fmt.Printf("Cory's age: %v\n", myMap["Cory"])
    // If we try to access the value of a key that doesn't exist, we get the default value of the type assigned for values, which is uint8 in our case [the default value for uint8 is 0]
    fmt.Printf("Accessing non-existent key: %v\n", myMap["Jason"])
    // Go also returns a second value when accessing a value in a map; Which is whether the value exists or not [as a Boolean]
    var val, isExists = myMap["Jason Voorhees"]
    if isExists{
        fmt.Printf("The value for \"Jason voorhees\" exists: %v\n", val)
    }else{
        fmt.Println("The value \"Jason voorhees\" doesn't exist")
    }

    // Can iterate through the map using the range keyword in a for loop
    // No order is preserved while iterating through a for loop and every run will have a different order of the keys/values
    for name, age := range myMap{
        fmt.Printf("Name: %v, Age: %v\n", name, age)
    }
    // Can also iterate through an array similarly
    for index, value := range intArr2{
        fmt.Printf("Array index: %v, Value: %v\n", index, value)
    }



    // There is no while loop in go and we need to use a for loop
    var i int = 0
    for i<10{
        fmt.Println(i)
	i = i+1
    }
    /* can also ommit the condition and use the break keyword to end at 10
    for {
        if i >= 10{break}
	fmt.Printf(i)
	i = i+1
    } */

    // Another way to define a for loop is using the conditions initializing, condition and post respectively. The post is executed everytime after the loop is done
    for i=0; i<10; i++{
        fmt.Println(i)
    }


    // Timed performance test for the slicing with/without preallocation
    var n int = 1000000
    var testSlice1 = []int{}
    var testSlice2 = make([]int, 0, n)

    fmt.Printf("Total time without preallocation: %v\n", timeloop(testSlice1, n))
    fmt.Printf("Total time with preallocation: %v", timeloop(testSlice2, n))
}

func timeloop(slice []int, n int) time.Duration{
    var t0 = time.Now()
    for len(slice)<n{
        slice = append(slice, 1)
    }
    return time.Since(t0)
}
