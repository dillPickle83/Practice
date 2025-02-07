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
  fmt.Println(<-c)
}

func single_process(c chan int){
    c <- 1234
}

func process(c chan int){
    defer close(c)
    for i:=0; i<5; i++{
        c <- 1
    }
}
