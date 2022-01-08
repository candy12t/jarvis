OBJ = bin/jarvis

all: clean build

.PHONY: build
build: statik
	go build -o $(OBJ) -v ./cmd/jarvis

.PHONY: lambda
lambda: clean statik
	GOOS=linux go build -o $(OBJ) -v ./cmd/jarvis
	mkdir dist
	zip -r dist/lambda.zip ./bin

.PHONY: statik
statik:
	statik -f -dest=./internal -src=./config -include=\*.yaml

.PHONY: clean
clean:
	rm -rf bin/*
