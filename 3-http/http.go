package main

import (
	"fmt"
	"net/http"
)

// handleGetFoo handles GET requests to return "FOO".
func handleGetGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method Not Allowed")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Jay Shree Ram"))
}

func main() {
	http.HandleFunc("/greet", handleGetGreet)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
