package main

import (
	"fmt"
	"errors"
)

func main(){
    var printVal string = "Hello world!"
    printMe(printVal)

    var numer int = 11
    denom := 2
    var divResult, err = intDivision(numer, denom)
    if err!=nil{
        fmt.Printf(err.Error())
    }else if divResult == 1{
	fmt.Println("The numerator and denominator were the same")
    }else{
        fmt.Printf("%v/%v = %v\n", numer, denom, divResult)
    }

    /* Can also write the above if else statements using switchcase
    switch{
        case err != Nil:
	    fmt.Printf(err.Error())
	case divResult == 1:
	    fmt.Println("The numerator and denominator were the same")
	default:
            fmt.Println("%v/%v = %v", numer, denom, divResult)
    unlike in C the break statement isn't necessary in all the cases */

    /* We can also put a variable in the swtich statement to evaluate with that variable 
    switch number{
	case 1:
            fmt.Printf("1 was close")
	case 2, 3:
	    fmt.Println("It was closer")
	default:
            ft.Println("Not even close")
    */

    var num1 = 46
    var num2 = 23
    var add, sub int = addSub(num1, num2)
    fmt.Printf("Addition: %v, Subraction: %v", add, sub)
}

func printMe(printValue string){
    fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, error) {
    // If the function returns a value, the return value type needs to mentioned
    var err error
    if denominator == 0{
	err = errors.New("Can't divide by zero")
	return 0, err
    }
    var result int = numerator/denominator
    return result, err
}

func addSub(num1 int, num2 int) (int, int) {
    // Can return multiple vales by declaring the return types
    var add int = num1 + num2
    var sub int = num1 - num2
    return add, sub
}
