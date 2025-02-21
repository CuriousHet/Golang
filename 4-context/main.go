package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// Create a context with a timeout of 300ms
	ctx := context.WithValue(context.Background(), "source", "fetchUserData")
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	userID := 10
	val, err := fetchUserData(ctx, userID)

	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Fetched Data:", val)
	}

	fmt.Println("Total Time Taken:", time.Since(start))
}

// Response struct to store fetched data
type Response struct {
	value int
	err   error
}

// fetchUserData fetches data from multiple sources concurrently
func fetchUserData(ctx context.Context, userID int) (int, error) {
	fmt.Println("Fetching data for user ID:", userID)
	fmt.Println("Context Source:", ctx.Value("source"))

	responseChannel := make(chan Response, 2)

	// Start goroutines for cache and third-party fetch
	go fetchCachedData(responseChannel)
	go fetchThirdPartyData(responseChannel)

	// Wait for the first successful response within the context deadline
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("timeout: data fetch took too long")
		case resp := <-responseChannel:
			if resp.err == nil {
				return resp.value, nil
			}
		}
	}

	return 0, fmt.Errorf("failed to fetch data")
}

// Simulated cache lookup (faster response)
func fetchCachedData(responseChannel chan Response) {
	time.Sleep(time.Millisecond * 100) // Simulating quick cache response
	responseChannel <- Response{value: 123, err: nil}
}

// Simulated third-party API call (slower response)
func fetchThirdPartyData(responseChannel chan Response) {
	time.Sleep(time.Millisecond * 500) // Simulating slow API response
	responseChannel <- Response{value: 666, err: nil}
}
