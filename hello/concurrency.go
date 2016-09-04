package main

import (
	"fmt"
	"time"
	"sync"
)

func getPara(v int) int{
	fmt.Printf("getPara %v\n", v)
	return v
}

func smallJob(v int) {
	fmt.Printf("Enter smallJob %v\n", v)
}


func concurrency(){
	fmt.Println("---Concurrency")
	
	// basicGoRoutine()
	// producerConsumer()
	// selectTest()
	syncMutex()
}

// What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?
// This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.


func syncMutex(){

	ct := SafeCounter{v:0}
	for i:= 0; i < 10; i++ {
		go func(v *SafeCounter){
			v.mux.Lock()
			v.v++
			v.mux.Unlock()
		}(&ct)
		go func(v *SafeCounter){
			fmt.Println(v.Value())
		}(&ct)
	}
	
	time.Sleep(10000 * time.Millisecond)
	

}

func (c *SafeCounter) Value() int {

	c.mux.Lock()
	// // Lock so only one goroutine at a time can access the map c.v.
	// We can also use defer to ensure the mutex will be unlocked as in the Value method.
	defer c.mux.Unlock()
	return c.v
}

type SafeCounter struct {
	v int
	mux sync.Mutex
}


func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		// The default case in a select is run if no other case is ready.
		default:
			fmt.Println("Default case, not blocking")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func selectTest() {
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	// The select statement lets a goroutine wait on multiple communication operations.

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}


func basicGoRoutine() {

	for i := 0; i < 10; i++{
		// The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
		go smallJob(getPara(i))
	}

	time.Sleep(1000 * time.Millisecond)	
}


func producerConsumer(){
	// unbuffered channels
	// c := make(chan int)

	//Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	c := make(chan int, 5)
	go consumer(c)
	go producer(c)
	time.Sleep(9000 * time.Millisecond)	
}

func producer(c chan int){
	for i := 0; i < 10; i++{
		fmt.Printf("Producer sending %v\n", i)
		c <- i
		fmt.Printf("Producer sent %v\n", i)
	}
	// A sender can close a channel to indicate that no more values will be sent.
	// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
	fmt.Println("producer closed the channel")
}

func consumer(c chan int){
	for {
		fmt.Printf("Consumer reading\n")
		f, ok := <-c
		if  !ok {
			fmt.Println("Channel closed, consumer quit")
			// ok is false if there are no more values to receive and the channel is closed.
			return
		}
		fmt.Printf("Consumer received %v\n", f)
	}
}
