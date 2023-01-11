#include "image.hpp"
#include "status.hpp"

#include "../include/pillow_resize.hpp"

#include <openssl/md5.h>

#include <string>
#include <sstream>
#include <cstdio>
#include <cmath>

Status Image::Open()
{
    img = cv::imread(path, CV_8S);
    if (img.empty()) {
	    std::fprintf(stderr, "fatal error: failed to open image");
	    return Status::ImageOpenFailed;
    }
    return Status::ImageOpenSuccess;
}

Status Image::Resize(Interpolation interpolation, double size)
{
    auto imgSize = img.size();
    auto maxSize = std::max(imgSize.width, imgSize.height);

    if (size != maxSize) {
	auto scale = size / maxSize;
	auto newSize = cv::Size(
		std::max(1, int(std::floor(scale * imgSize.width))),
		std::max(1, int(std::floor(scale * imgSize.height))));

	PillowResize::InterpolationMethods pillowInterpolation;
    	if (interpolation == Interpolation::Bilinear) {
    	    pillowInterpolation = PillowResize::InterpolationMethods::INTERPOLATION_BILINEAR;
    	}
    	else {
	    std::fprintf(stderr, "fatal error: interpolation not supported");
    	    return Status::InterpolationNotSupported;
    	}
    	img = PillowResize::resize(img, newSize, pillowInterpolation);
    }
    return Status::ImageResizeSuccess;
}

const std::vector<double> Image::PixelMean()
{
    std::vector<double> pixelMean;
    for (std::size_t i=0; i<img.rows; i++) {
	for (std::size_t j=0; j<img.cols; j++) {
	    double r = img.at<cv::Vec3b>(i, j)[2];
	    double g = img.at<cv::Vec3b>(i, j)[1];
	    double b = img.at<cv::Vec3b>(i, j)[0];
	    pixelMean.emplace_back(std::floor( (r + g + b) / 3 ));
	}
    }
    return pixelMean;
}

const std::string Image::Hash()
{
    std::vector<double> pixelMean = PixelMean();
    for (std::size_t i=0; i<pixelMean.size(); i++) {
	pixelMean[i] = std::ceil( pixelMean[i] / 4 ) * 4;
    }
    
    std::uint8_t result[MD5_DIGEST_LENGTH];
    MD5((std::uint8_t *)pixelMean.data(), pixelMean.size() * sizeof(double), result);

    std::stringstream ss;
    for (int i=0; i<MD5_DIGEST_LENGTH; i++) {
	ss << std::setfill('0') << std::setw(2) << std::hex << static_cast<int>(result[i]);
    }
    return ss.str();
}
