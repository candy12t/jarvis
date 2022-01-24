EXEC = ./bin/jarvis

.PHONY: install
install:
	go mod download
	go install github.com/rakyll/statik

.PHONY: debug
debug: clean statik
	go build -tags debug -o $(EXEC) -v ./cmd/jarvis
	$(EXEC)

.PHONY: statik
statik:
	statik -f -dest=./internal -src=./config -include=\*.yaml

.PHONY: build-JarvisFunction
build-JarvisFunction: statik
	GOOS=linux go build -o $(ARTIFACTS_DIR) -v ./cmd/jarvis

.PHONY: build
build:
	sam build JarvisFunction

.PHONY: deploy
deploy:
	sam deploy

.PHONY: test
test:
	ls config/sample/*.yaml.sample | sed 's/\.sample//' | xargs -I{} cp {}.sample {}
	statik -f -dest=./internal -src=./config/sample -include=\*.yaml
	go test ./... -count=1
	rm -rf config/sample/*.yaml

.PHONY: clean
clean:
	rm -rf bin/*
