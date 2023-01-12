package image

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
	hash := C.CIMGLIB_MD5FromBytes((*C.uchar)(&data[0]), C.ulong(len(data)), C.double(size))
	if validHashPtr(hash) {
		// We only free the hash variable if memory has been allocated!
		defer C.free(unsafe.Pointer(hash))
	}
	return C.GoString(hash)
}

func MD5FromFile(path string, size float64) string {
	if path == "" {
		log.Println("failed to calculate md5 from file: path cannot be empty")
		return ""
	}
	pathParam := C.CString(path)
	defer C.free(unsafe.Pointer(pathParam))
	hash := C.CIMGLIB_MD5FromFile(pathParam, C.double(size))
	if validHashPtr(hash) {
		// We only free the hash variable if memory has been allocated!
		defer C.free(unsafe.Pointer(hash))
	}
	return C.GoString(hash)
}

func validHashPtr(hash *C.char) bool {
	emptyStr := C.CString("")
	defer C.free(unsafe.Pointer(emptyStr))
	if hash == nil {
		return false
	}
	return true
}
