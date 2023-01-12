package main

import (
	"bytes"
	"fmt"
	"github.com/klauspost/compress/zstd"
	"io"
)

func main() {
	// Compress some data
	data := []byte("This is some data that we want to compress")
	compressedData := new(bytes.Buffer)
	w, _ := zstd.NewWriter(compressedData)
	w.Write(data)
	w.Close()

	decompressedData := new(bytes.Buffer)

	// Decompress the data
	io.Copy(decompressedData, zstd.NewReader(compressedData))

	fmt.Println(string(decompressedData))
}
