This project demonstrates fetching user data from multiple sources (cache and third-party API) using Go's concurrency features. The first successful response is returned, and the execution is managed using Go's `context` package to handle timeouts efficiently.

   ```sh
   go run main.go
   ```

## Expected Output
```
Context Value: bar
Fetching data for user ID: 10
Fetched Data: 123
Time taken: 104.8237ms
```
