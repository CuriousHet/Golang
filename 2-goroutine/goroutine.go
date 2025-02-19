package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	responseChannel := make(chan any, 4)
	wg := &sync.WaitGroup{}
	wg.Add(4) // Increased to match the number of goroutines

	// Launch multiple goroutines to fetch different user data
	go fetchUserLikes(userName, responseChannel, wg)
	go fetchUserMatch(userName, responseChannel, wg)
	go fetchUserPosts(userName, responseChannel, wg)
	go fetchUserFollowers(userName, responseChannel, wg)

	wg.Wait() // Block until all goroutines call wg.Done()
	close(responseChannel)

	fmt.Println("\nCollected Responses:")
	for resp := range responseChannel {
		fmt.Println("Response:", resp)
	}

	totalTimeWithGoroutines := time.Since(start)
	fmt.Println("\nTime taken with Goroutines:", totalTimeWithGoroutines)

	// Running without goroutines to compare execution time
	singleThreadedStart := time.Now()
	likes := fetchUserLikesSync(userName)
	match := fetchUserMatchSync(userName)
	posts := fetchUserPostsSync(userName)
	followers := fetchUserFollowersSync(userName)
	totalTimeWithoutGoroutines := time.Since(singleThreadedStart)

	fmt.Println("\nSingle-threaded Responses:")
	fmt.Println("Likes:", likes)
	fmt.Println("Match:", match)
	fmt.Println("Posts:", posts)
	fmt.Println("Followers:", followers)

	fmt.Println("\nTime taken without Goroutines:", totalTimeWithoutGoroutines)
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

func fetchUserPosts(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 120)
	responseChannel <- []string{"Post1", "Post2", "Post3"}
	wg.Done()
}

func fetchUserFollowers(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 130)
	responseChannel <- 250
	wg.Done()
}

// Synchronous functions for comparison
func fetchUserLikesSync(userName string) int {
	time.Sleep(time.Millisecond * 150)
	return 11
}

func fetchUserMatchSync(userName string) string {
	time.Sleep(time.Millisecond * 100)
	return "RAM"
}

func fetchUserPostsSync(userName string) []string {
	time.Sleep(time.Millisecond * 120)
	return []string{"Post1", "Post2", "Post3"}
}

func fetchUserFollowersSync(userName string) int {
	time.Sleep(time.Millisecond * 130)
	return 250
}
