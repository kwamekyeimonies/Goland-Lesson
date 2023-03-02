package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	// "io/ioutil"
)

func main() {

	payload := []byte("Hello high value software engineer")
	// hashAndBroadCast(bytes.NewReader(payload))
	hashAndBroadCast(NewHashReader(payload))

	var buffer_test bytes.Buffer
	buffer_test.WriteString("I Love you")
	buffer_test.WriteString("From the rising of the sun, till the ....")
	fmt.Println("Treasure of my heart :",buffer_test)

}

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: *bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

// func hash(h *hashReader) string{
// 	return hex.EncodeToString(h.buf.Bytes())
// }

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadCast(r HashReader) error {
	// broad_cast, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	return err
	// }

	// hash := sha1.Sum(broad_cast)

	// hash := r.(*hashReader).hash()
	hash := r.hash()
	// fmt.Println(hex.EncodeToString(hash[:]))
	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	broad_cast, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("String of the bytes: ", string(broad_cast))
	return nil
}
