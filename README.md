# cgoimg

Minimal image library implemented in C++ with bindings for Go.

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes. See deployment
for notes on deploying the project on a live system.

### Prerequisites

Requirements for the software and other tools to build, test and push
- g++ or clang++
- [Go](https://https://go.dev/doc/install.com)
- [OpenSSL](https://www.openssl.org/)
- [OpenCV](https://opencv.org/)

### Build
Build the cgoimg library by simply running

    make

You can also build the examples in the examples folder the same way.


### Installing
Install the library with 

    make install

Note that this command might require superuser privileges. The installed location is

    /usr/lib/libcgoimg.so

You could also change to a location that is looked up by LD_LIBRARY_PATH.

### Running the tests
Run the tests with the command

    make test

### Usage
You can use <b>cgoimg</b> as any other Go package. Check the examples folder for more information.