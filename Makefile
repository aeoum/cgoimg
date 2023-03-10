OUT=libcgoimg.so

.PHONY: build
build:
	g++ -o $(OUT) ./cimglib/src/*.cpp -std=c++2a -O3 -fPIC -shared `pkg-config --libs --cflags opencv4` 
	strip $(OUT)

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
