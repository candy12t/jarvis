OBJ = ./bin/jarvis

all: clean build

.PHONY: build
build: statik
	go build -o $(OBJ) -v ./cmd/jarvis

.PHONY: run
run: build
	$(OBJ)

.PHONY: lambda
lambda: clean statik
	GOOS=linux go build -o $(OBJ) -v ./cmd/jarvis
	zip -r dist/lambda.zip ./bin

.PHONY: statik
statik:
	statik -f -dest=./internal -src=./config -include=\*.yaml

.PHONY: test
test:
	statik -f -dest=./internal -src=./config/sample -include=\*.yaml
	go test ./... -count=1

.PHONY: clean
clean:
	rm -rf bin/*
