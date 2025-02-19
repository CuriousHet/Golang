package main

import (
	"io",
	"encoding/hex",
	"fmt",
	"bytes",
	"io/ioutil"
)

type HashReader iterface {
	io.Reader
	hash() string
}

func main() {

	payload := []byte("i dont wanna live forever")
	hashAndBroadcast(NewHashReader(payload))
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err != ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string: ",string(b))

	return nil
}
