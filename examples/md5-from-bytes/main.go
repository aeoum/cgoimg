package main

import (
	"fmt"
	"os"

	"github.com/aeoum/cgoimg"
)

func main() {
	var size float64 = 224
	data, err := os.ReadFile("../../testdata/01200073904a56e3da52af6330c6f380fe235503")
	if err != nil {
		panic(err)
	}
	hashBytes := cgoimg.MD5FromBytes(data, size)
	hashFile := cgoimg.MD5FromFile("../../testdata/01200073904a56e3da52af6330c6f380fe235503", size)

	fmt.Printf("Hash from bytes: %v\n", hashBytes)
	fmt.Printf("Hash from file: %v\n", hashFile)
}
