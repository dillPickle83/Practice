package main

import "fmt"
import "unicode/utf8"

func main(){
    var intNum int16 = 32767
    // Can also use int8, int16, int32 and int64 [bitsize variants] depending on the number of bits to be used by the variable
    // The largest int16 that can be stored is 32767; Trying to initialize a var bigger than int16 [>32767] will throw a compiler error [an overflow error]
    // If we initialize the var as 32767 and then in the next line we give "intNum = intNum +1", the compiler won't throw an err but print -32768
    // We also have unsigned ints, uint, uint8, uint16,...... which store only positive integers and therefore can store twice 
    // the amount of positive integers as int type in the same memory bandwidth
    fmt.Println(intNum)

    var floatNum float64 = 12345678.9
    // They have float32 and float64 depending on the mem requirements; 64 takes more memory but can store decimals precisely
    // Storing a decimal number as float32 and then printing it will give a different number which is how the computer stores it [prints 123456789.00000]
    // Meanwhile, float64 gives the actual decimal number
    fmt.Println(floatNum)

    var num1 float32 = 10.1
    var num2 int32 = 12
    var add_res float32 = num1 + float32(num2)
    // Can add the same bitsize values
    fmt.Println(add_res)

    var numer int8 = 3
    var denom int8 = 2
    fmt.Println(numer/denom)		// 1 since the operation is rounded down by default
    fmt.Println(numer%denom)		// 1 since the remainder here is one

    var myString string = "Hello" + " " + "World" 	// Can also use '' to define multiline strings
    fmt.Println(myString)
    fmt.Println(len(myString))
    // Since go uses UTF coding, any character outside of the vanilla ASCII characterset will have more than 1 len
    // Length of "A" is 1, but γ [gammma] is 2, since the len function gets the number of bytes in the string
    fmt.Println(len("test"))
    fmt.Println(utf8.RuneCountInString("γ"))
    // Runes are a data type in go and represents characters

    var myRune rune = 'a'
    fmt.Println(myRune)		// Ouput will be 97

    var defaultInt int
    var defaultString string
    fmt.Println(defaultInt)
    fmt.Println(defaultString)
    // The default int value will be 0 and this applies to uint, uint8,.... , int8,..... , float32, float64 and rune
    // The default string value is ''

    myTxt := "text"
    fmt.Println(myTxt)
    // Here, the type of the variable is inferred using the := shorthand
    // Can also initialize mutiple variables as below

    var1, var2 := 1, 2
    // var var1, var2 int = 1, 2
    // var var1, var2 = 1, 2
    fmt.Println(var1, var2)

    // Good practice to specify types when defining vars when it isn't obvious
    // var funcString string = foo()

    const myConst string = "This is a constant"
    const pi float64 = 3.145
    // Constants can't be left in default state and needs to be initialized with a value
    // Constants values can't be changed later
    fmt.Println(myConst)
    fmt.Println(pi)
}
