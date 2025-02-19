package main

import (
	"io",
	"time",
	"sync"
)

func main() {
	
	start := time.Now()
	userName := fetchUser()
	responseChannel := make(chan any, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go fetchUserLikes(userName, responseChannel, wg)
	go fetchUserMatch(userName, responseChannel. wg)

	wg.Wait() // block until we have 2 wg.Done()
	close(responseChannel)

	for resp := range responseChannel {

		fmt.Println("reponse: ",resp)
	}
	// fmt.Println("likes", likes)
	// fmt.Println("match", match)
	fmt.Println("time takeb: ",time.Since(start))
}

func fetchUser() string {
	
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func fetchUserLikes(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	responseChannel <- 11
	wg.Done()
}

func fetchUserMatch(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	
	time.Sleep(time.Millisecond * 100)

	responseChannel <- "RAM"
	wg.Done()
}
