PROJECT=algotrivia

printProject:
	echo ${PROJECT}

init:
	go get ./...
	go get github.com/google/wire/cmd/wire

generate:
	go generate ./...

unittest:
	go test -short  ./...

build:
	go build -o ${PROJECT}

run:
	make build
	./${PROJECT}

run_all:
	go run main.go injection.go

build_test_run:
	make build
	go test ./.../test  -v
	make run

docker:
	docker build -t algotrivia .