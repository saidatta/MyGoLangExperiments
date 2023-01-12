package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/klauspost/compress/zstd"
)

func main() {
	// Compress some data using a dictionary
	data := []byte("This is some data that we want to compress using a dictionary")
	dictionary := []byte("dictionary")
	compressedData := new(bytes.Buffer)
	w, _ := zstd.NewWriterLevelDict(compressedData, zstd.BestCompression, dictionary)
	w.Write(data)
	w.Close()

	// Decompress the data
	decompressedData, _ := ioutil.ReadAll(zstd.NewReader(compressedData))

	fmt.Println(string(decompressedData))
}
