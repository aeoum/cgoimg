#ifndef CIMGLIB_IMAGE_GO_H
#define CIMGLIB_IMAGE_GO_H

typedef unsigned char uint8;

#ifdef __cplusplus
extern "C" {
#endif

const char *CIMGLIB_MD5FromBytes(uint8 *data, size_t dataSize, double size);
const char *CIMGLIB_MD5FromFile(char *path, double size);

#ifdef __cplusplus
}
#endif

#endif
