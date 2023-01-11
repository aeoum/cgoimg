#include "image.hpp"
#include "image.go.h"

#include <vector>
#include <cstdio>
#include <cstdlib>
#include <cstdint>

static const char *AllocHash(Image& img)
{
    std::string hash = img.Hash();
    char *res = static_cast<char *>(std::malloc(hash.size()));
    sprintf(res, "%s", hash.c_str());
    return res;
}

const char *CIMGLIB_MD5FromBytes(std::uint8_t *data, std::size_t dataSize, double size)
{
    std::vector<std::uint8_t> vec(data, data + dataSize);
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

