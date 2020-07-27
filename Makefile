NAME := main

BUILD_DATE := $(shell LANG=c date)
GIT_REVISION := $(shell git rev-list -1 HEAD)

build:
	go build -o $(NAME) -ldflags=" -X 'main.BUILD_DATE=$(BUILD_DATE)' -X 'main.GIT_REVISION=$(GIT_REVISION)'"

all: build

clean-build:
	@rm -rf saved $(NAME) *.log *.wav

clean: clean-build
