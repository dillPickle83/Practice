package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "résumé"
	fmt.Println(myString)
	fmt.Printf("The length of the string: %v\n",  len(myString))
	// When we try to index and print the string, we get strange values. This is because in go, the string is stored as bits of the singular characters as per the UTF-8 encoding. The UTF-8 encoding sometimes stores special characters as 2 bits instead of 1 due to its design. This will cause the length of the string to be different compared to the actual character size.
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	// The following for loop will print the index and the value of the string array, the index will not be a straight 0 to 5 since there are special characters. Each special characters will take 2 or more bits and wil increase the overall length of the string, so a few indexes will be skipped over. The value will also be different, since it's actually the value of the bit.
	for i, v := range myString{
		fmt.Println(i, v)
	}

	// We can declare the above with rune types and it will represent the unicode values and the proper length of the list and such
	var myString2 = []rune("résumé")
	fmt.Printf("The length of the rune str: %v\n", len(myString2))
	var index2 = myString2[1]
	fmt.Printf("The value and type for rune str: %v, %T\n", index2, index2)

	for i, v := range myString2{
		fmt.Println(i, v)
	}

	var myRune = "a"
	fmt.Printf("myRune = %v\n", myRune)

	var stringSlice = []string{"H", "e", "l", "l", "o"}
	// The below method of creating a second str using concatenation is slow and a new string is created each time in the loop which is inefficient
	/*
	var catStr = ""
	for i := range stringSlice{
		catStr += stringSlice[i]
	}
	fmt.Println("After concatenating stringSlice to catStr: %v", catStr)
	*/
	// Instead using the strings.builder from the go strings package will make it more efficient. In the for loop this wiill allocate an array internally which will have parts of the string written to it using the WriteString function and then in the end, the string is actually created using the .string() function.
	var strBuilder strings.Builder
	for i := range stringSlice{
		strBuilder.WriteString(stringSlice[i])
	}
	var BuildStr = strBuilder.String()
	fmt.Printf("StrBuild string: %v", BuildStr)
}
