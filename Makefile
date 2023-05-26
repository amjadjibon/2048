run:
	go run main.go

build:
	env GOOS=js GOARCH=wasm go build -o 2048.wasm
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
