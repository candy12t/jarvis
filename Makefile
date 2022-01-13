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
	ls config/sample/*.yaml.sample | sed 's/\.sample//' | xargs -I{} cp {}.sample {}
	statik -f -dest=./internal -src=./config/sample -include=\*.yaml
	go test ./... -count=1
	rm -rf config/sample/*.yaml

.PHONY: clean
clean:
	rm -rf bin/*
