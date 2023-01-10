OUT=libcgoimg.so

.PHONY: build
build:
	g++ -o $(OUT) ./cimglib/src/*.cpp -std=c++20 -O3 -fPIC -shared `pkg-config --libs --cflags opencv4` -lcrypto -lssl

.PHONY: install
install:
	cp $(OUT) /usr/lib/$(OUT)

.PHONY: remove
remove:
	rm /usr/lib/$(OUT)

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf $(OUT)
