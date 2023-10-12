.DEFAULT_GOAL := build
bin=loxilb

build:
	@go build -o ${bin}

run: build
	./$(bin)