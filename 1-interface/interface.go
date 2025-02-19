package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

// HashReader interface combines io.Reader with a hash function
// that returns a hex-encoded hash of the data.
type HashReader interface {
	io.Reader
	hash() string
}

// hashReader struct wraps a bytes.Reader and stores the original buffer
// to allow both reading and hashing the same data.
type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

// NewHashReader initializes a hashReader with the given byte slice.
func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

// hash computes and returns the hex-encoded hash of the original data.
func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

// hashAndBroadcast calculates the hash of the data and then broadcasts it.
func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println("Computed Hash:", hash)

	return broadcast(r)
}

// broadcast reads all data from the reader and prints it as a string.
func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("Broadcasting Data:", string(b))
	return nil
}

func main() {
	// Sample data to be hashed and broadcasted
	payload := []byte("I don't wanna live forever")

	// Create a new hashReader instance and process it
	err := hashAndBroadcast(NewHashReader(payload))
	if err != nil {
		fmt.Println("Error:", err)
	}
}
