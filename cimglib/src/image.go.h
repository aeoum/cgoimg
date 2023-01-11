#ifndef CIMGLIB_IMAGE_GO_H
#define CIMGLIB_IMAGE_GO_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

const char *CIMGLIB_MD5FromBytes(uint8_t *data, size_t dataSize, double size);
const char *CIMGLIB_MD5FromFile(char *path, double size);

#ifdef __cplusplus
}
#endif

#endif
