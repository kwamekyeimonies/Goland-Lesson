package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// The func Send sends integers to the ch channel

func receive(ch <-chan int) {
	for value := range ch{
		fmt.Println(value)
	}
}

// The receive fxn listen on t

func main() {

	ch := make(chan int)
	go send(ch)
	time.Sleep(2*time.Second)
	go receive(ch)
	fmt.Println(time.Now(),"Hello")

	// go Greeting()
	//Create a channel
	// mychannel := make(chan string)

	// go func ()  {
	// 	fmt.Println(time.Now(), "Taking a Break")
	// 	time.Sleep(2*time.Second)
	// 	mychannel<-"hello"
	// }()

	// fmt.Println(time.Now(),"Waiting for Message")
	// v := <- mychannel

	// fmt.Println(time.Now(),"Message Received",v)

	// chx := make(chan int)

	// go func(){
	// 	for i:=0; i<5; i++{
	// 		fmt.Println(time.Now(),i,"Sedning")
	// 		chx<-i
	// 		fmt.Println(time.Now(),i,"Sent")

	// 		fmt.Println(time.Now(),"all completed")
	// 	}
	// }()

	// // time.Sleep(2 * time.Second)
	// fmt.Println(time.Now(),"Waiting for Messages")
	// fmt.Println(time.Now(),"received",<-chx)
	// fmt.Println(time.Now(),"received",<-chx)
	// fmt.Println(time.Now(),"received",<-chx)

	// fmt.Println(time.Now(), "Exciting")
}

func Greeting() {
	fmt.Println("Goroutines not parallelism")
}
