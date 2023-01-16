package cgoimg

// #cgo LDFLAGS: -L. -lcgoimg
// #include "./cimglib/src/image.go.h"
// #include <stdlib.h>
import "C"
import (
	"log"
	"unsafe"
)

func MD5FromBytes(data []byte, size float64) string {
	if len(data) == 0 {
		log.Println("failed to calculate md5 from bytes: byte array cannot be empty")
		return ""
	}
	pHash := C.CIMGLIB_MD5FromBytes((*C.uchar)(&data[0]), C.ulong(len(data)), C.double(size))
	if validHashPtr(pHash) {
		// We only free the hash variable if memory has been allocated!
		defer C.free(unsafe.Pointer(pHash))
	}
	return C.GoString(pHash)
}

func MD5FromFile(path string, size float64) string {
	if path == "" {
		log.Println("failed to calculate md5 from file: path cannot be empty")
		return ""
	}
	pPath := C.CString(path)
	defer C.free(unsafe.Pointer(pPath))
	pHash := C.CIMGLIB_MD5FromFile(pPath, C.double(size))
	if validHashPtr(pHash) {
		// We only free the hash variable if memory has been allocated!
		defer C.free(unsafe.Pointer(pHash))
	}
	return C.GoString(pHash)
}

func validHashPtr(pHash *C.char) bool {
	pEmptyStr := C.CString("")
	defer C.free(unsafe.Pointer(pEmptyStr))
	if pHash == nil {
		return false
	}
	return true
}
