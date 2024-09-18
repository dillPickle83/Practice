package main
import "fmt"

// Defining a struct is like defining your own type using values and its corresponding types
type nonEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
	owner	
}

type owner struct{
	name string
}

type gasEngine struct{
	mpg uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

// An interface can be defined if we have similar types and we want to make a function accept multiple types. The condition is that the types mentioned in the interface must have the same methods, return types, etc.
type engine interface{
	milesLeft() uint8
}
// We are defining a milesLeft() method for the struct which is returning a uint8 value when the milesLeft() method is called which is a multiplication of the mpg and gallons values defined
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh*e.kwh
}

// We can also call the method inside another method. We can specify the type that needs to be accepted or we can make an interface which will allow the function to accept similar types defined under the interface
func canMakeIt(e engine, miles uint8) {
	if miles<=e.milesLeft(){
		fmt.Println("You can make it")
	}else{
		fmt.Println("Can't make it")
	}
}


func main() {
	// We can use the gasEngine as a type since the struct has been defined
	/* 
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons) */
	// The above definition will default to the default value of that type [uint8 default value is 0] since no value is set
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 44}
	// fmt.Println(myEngine.mpg, myEngine.gallons)
	// Can also assign values, the below ways
	// var myEng gasEngine = gasEngine{24, 44}
	// myEngine.mpg = 20 			---- This will overwite the value that we set in the prev init

	// We can also define another struct into another struct like above definition
	// Can skip defining the ownerInfo and can directly mention the type "owner" and the value can be accessed using:
	// myEngine.name
	// Like using a custom type we can also add in the default types like int, etc.
	var tempEngine nonEngine = nonEngine{25, 45, owner{"Alex"}, owner{"name"}}
	fmt.Println(tempEngine.mpg, tempEngine.gallons, tempEngine.ownerInfo.name, tempEngine.name)
	// Can declare a struct and the values at the same time, but these are not reusable
	var myStruct = struct{
		mpg int32
		gallons uint8
	}{25, 15}
	fmt.Println("Defining as a var")
	fmt.Println(myStruct.mpg, myStruct.gallons)

	// We can assign a method to the type as well like a function
	var myEngine gasEngine = gasEngine{25, 23}
	fmt.Printf("The amount of miles left in the engine is: %v\n", myEngine.milesLeft())
	fmt.Printf("Can you make 100 miles?\n")
	var miles uint8 = 100
	canMakeIt(myEngine, miles)

	var elEngine electricEngine = electricEngine{12, 56}
	canMakeIt(elEngine, miles)
}
