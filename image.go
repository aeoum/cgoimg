package image

// #cgo LDFLAGS: -L. -lcgoimg
// #include "./cimglib/src/image.go.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

func MD5(path string, size float64) string {
	pathParam := C.CString(path)
	sizeParam := C.double(size)
	defer C.free(unsafe.Pointer(pathParam))
	hash := C.CIMGLIB_ImageHash(pathParam, sizeParam)
	emptyStr := C.CString("")
	defer C.free(unsafe.Pointer(emptyStr))
	if *hash != *emptyStr {
		// We only free the hash variable if memory has been allocated!
		defer C.free(unsafe.Pointer(hash))
	}
	return C.GoString(hash)
}
