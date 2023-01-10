#include "image.hpp"
#include "image.go.h"

#include <cstdio>
#include <cstdlib>

const char *CIMGLIB_MD5(char *path, double size)
{
    Image img = Image(path);
    if (img.Open() != Status::ImageOpenSuccess) {
	    return nullptr;
    }

    if (img.Resize(Interpolation::Bilinear, size) != Status::ImageResizeSuccess) {
	    return nullptr;
    }

    std::string hash = img.Hash();
    char *res = static_cast<char *>(std::malloc(hash.size()));
    sprintf(res, "%s", hash.c_str());
    return res;
}
