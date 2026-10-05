BINARY := container-runtime
CMD    ?= /bin/bash
GO := /usr/local/go/bin/go

.PHONY: all build run clean

all: build

build:
	$(GO) build -o $(BINARY) .

run: build
	sudo ./$(BINARY) $(CMD)

clean:
	rm -f $(BINARY)