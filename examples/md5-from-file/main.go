package main

import (
	"fmt"

	"github.com/aeoum/cgoimg"
)

func main() {
	var size float64 = 224
	hash := cgoimg.MD5FromFile("img", size)
	fmt.Printf("Hash: %s\n", hash)
}
