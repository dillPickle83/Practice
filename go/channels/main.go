// Channels hold data, are thread safe (doesn't curropt the memory when reading/writing data) and listen for when data is removed or added into a channel

package main

import "fmt"

func main(){
  var c = make(chan int)
  // In the following 2 lines of code, the first line takes the value 1 and stores that in the channel c and the below is the syac to point the value to occupy the resp channel. The channel at this point can be thought of as an array. The next line takes the value in the channel (the 1) and then assigns it to the variable i.
  // c <- 1
  // var i = <- c
  // fmt.Println(i)
  // When we try printing the value of i after it has taken the value from the channel, the code will fail with a deadlock. This is because when we're trying to write to the unbuffered channel, the program will write to the channel and then keep waiting for something else to read from it. Hence the code will wait in the c <- 1 line forever. Go is smart enough to realize this and throw the error. Channels aren't meant to be used like this and need to be used in conjunction with goroutines.
  // To use the channels along with a goroutine, we need to create a funciton and call that function as a goroutine and then call the value
  go single_process(c)
  fmt.Println("Channels with single value")
  fmt.Println(<-c)
  go process(c)
  fmt.Println("Channels with multiple value in a for loop")
  // When we call the process function, the function for loop will assign the first value 0 and then the code will move to the below for loop since the process function is called as part of a goroutine. In the for loop it will wait for the value to becopme available to print and then print the value. SO the value 0 will be printed right as the goroutine processes the next loop in the function and adds the value 1 to the channel as part of loop 2, where the for loop in the main block will read and print 1 and so forth.
  for i:= range c{
        fmt.Println(i)
  }
}

func single_process(c chan int){
    c <- 1234
}

func process(c chan int){
    // The below line is needed since if we didn't add this line, the program will finish the for loop and assign and print till value 4 in the for loops and then it will end with a deadlock since the for loop in the main program will be expecting more output. Calling the close(c) (at the end of the needed code block) will tell the program that the channel will not produce anymore values. Hence, the main for loop will print the last value of the for loop and then move on to the next block. The defer keyword is just a way to tell the program that after the code block in this function is done, that will mark the end of the channel.
    defer close(c)
    for i:=0; i<5; i++{
        c <- i
    }
}
