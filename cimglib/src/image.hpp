#ifndef CIMGLIB_IMAGE_HPP
#define CIMGLIB_IMAGE_HPP

#include "status.hpp"

#include <string>
#include <cstdint>
#include <opencv2/opencv.hpp>
#include <opencv2/imgcodecs.hpp>

#define IMG_SIZE 224

enum class Interpolation { Bilinear = 0 };

class Image {
public:
    explicit Image(const std::string& path)
	: path(path)
    {
    }

    Image(std::vector<std::uint8_t>& data, int flag)
    {
        img = cv::imdecode(data, flag);
    }

    ~Image()
    {
    }

    Status Open();
    Status Resize(Interpolation interpolation, double size=IMG_SIZE);
    const std::vector<double> PixelMean();
    const std::string Hash();
private:
    std::string path;
    cv::Mat img;
};

#endif
