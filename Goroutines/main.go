package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RoundNumbers(name string,limit int, sleep int){
	for i:=0; i <= limit; i++{
		fmt.Println(name,rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))
	}
}

// Using the Waitgroup sync int the goroutines

func RoundNumbers_Sync(wg *sync.WaitGroup,name string, limit int, sleep int){
	defer wg.Done()
	for i:=0; i<=limit; i++{
		fmt.Println(name,rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))
	}
}

func WriteChannel(wg *sync.WaitGroup, limitchannel chan int, stop int){
	defer wg.Done()

	for i:=1; i<= stop; i++{
		limitchannel <- i
	}
}

func main(){

	// var first_channel chan int
	// fmt.Println(first_channel)

	// second_channel := make(chan int)
	// fmt.Println(second_channel)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	limitChannel := make(chan int,2)
	defer close(limitChannel)
	go WriteChannel(wg,limitChannel,2)
	wg.Wait()
	fmt.Println(<-limitChannel)
	fmt.Println(<-limitChannel)

	// start := time.Now()
	// username := fetchUser()
	// // likes := fetchUsersLikes(username)
	// // match := fetchUsersMatch(username)

	// //create a channey
	// respch := make(chan any)
	// go fetchUsersLikes(username, respch)
	// go fetchUsersMatch(username,respch)

	// // fmt.Println("likes: ",likes)
	// // fmt.Println("match: ", match)
	// fmt.Println("took: ",time.Since(start))
}

// func fetchUser() string{
// 	time.Sleep(time.Millisecond * 100)
// 	return "Dexoangle@001"
// }

// func fetchUsersLikes(userName string,respchan chan any){
// 	time.Sleep(time.Millisecond * 120)

// 	return <- 99
// }

// func fetchUsersMatch(username string,respch chan any){
// 	return <- "Tenkorang Daniel"
// }