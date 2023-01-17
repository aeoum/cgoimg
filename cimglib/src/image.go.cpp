#include "image.hpp"
#include "image.go.h"

#include <vector>
#include <cstdio>
#include <cstdlib>

typedef unsigned char uint8;

const char *AllocHash(Image& img)
{
    if (img.Hash().size() == 0) {
        std::fprintf(stderr, "%s failed:%s:%d: hash is empty",
            __func__, __FILE__, __LINE__);
        return nullptr;
    }
    std::string hash = img.Hash();
    char *res = static_cast<char *>(std::malloc(hash.size()));
    sprintf(res, "%s", hash.c_str());
    return res;
}

const char *CIMGLIB_MD5FromBytes(uint8 *data, std::size_t dataSize, double size)
{
    std::vector<uint8> vec(data, data + dataSize);
    Image img = Image(vec, CV_8S);
    if (img.Resize(Interpolation::Bilinear, size) != Status::ImageResizeSuccess) {
	    return nullptr;
    }
    return AllocHash(img);
}

const char *CIMGLIB_MD5FromFile(char *path, double size)
{
    Image img = Image(path);
    // Only open an image if it's loaded from a file.
    if (img.Open() != Status::ImageOpenSuccess) {
	    return nullptr;
    }
    if (img.Resize(Interpolation::Bilinear, size) != Status::ImageResizeSuccess) {
	    return nullptr;
    }
    return AllocHash(img);
}

