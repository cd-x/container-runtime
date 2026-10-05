BINARY := container-runtime
CMD    ?= /bin/bash

.PHONY: all build run clean

all: build

build:
	go build -o $(BINARY) .

run: build
	sudo ./$(BINARY) $(CMD)

clean:
	rm -f $(BINARY)