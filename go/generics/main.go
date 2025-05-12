package main

import "fmt"

func main(){
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
	// Let's say we want to write this same function but with different types. We will need to define a new function for each type sumIntSlice, sumFloat32Slice, sum,Float64Slice, etc. and in each of those functions (like below sumSlice) we will need to use the same for loop and the same reutrn. The only thing to changer would be the var sum declaration (var sum int, var sum float 32, etc.) To avoid the repetition we use the generics in go where we can delare the function to take any of the types and then declare that as the type for the variable.	

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))

	// We can also use an any type so that we don't have to specify what kinda type the function can accept. The [] in the isEmpty[] is optional and can be ignored since it accepts any kind of value.
	var emptySlice = []int{}
	fmt.Println(isEmpty(emptySlice))

	fmt.Println(isEmpty(intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _,v := range slice{
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}
