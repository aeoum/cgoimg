package image

// #cgo LDFLAGS: -L. -lcgoimg
// #include "./cimglib/src/image.go.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

func MD5FromBytes(data []byte, size float64) string {
	if len(data) == 0 {
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
