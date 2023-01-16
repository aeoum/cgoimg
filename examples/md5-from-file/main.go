package main

import (
	"fmt"

	"github.com/aeoum/cgoimg"
)

func main() {
	var size float64 = 224
	hash := cgoimg.MD5FromFile("../../testdata/0120001c599dc0af44615e7fe8c41b255a0d288d", size)
	fmt.Printf("Hash: %s\n", hash)
}
