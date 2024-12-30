package main

import (
	"fmt"
	"time"
	"sync"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	// Simulate DB call delay with sleep
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The results from the database is:", dbData[i])
	// This lock is needed since if we run this without any lock, the results could be unpredictble since the goroutines will all be accessing the same memory to write in the results in the splice. When we put the lock, each of the goroutines will place a lock on the address before actually writing on that address and then unlock it once the call is done.
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	// The sync package also has a read lock and unlock in addition to the lock and unlock. This just gives more flexibility since the lock and unlock would completely lock off the address from other goroutines.
	// The read lock will also wait for the full lock to complete before placing the full lock
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
