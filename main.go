package main

import "fmt"

//go:generate go build -tags "main evm java" -x -o java/libshare_de_main.so -ldflags "-s -w" -buildmode=c-shared  dechain-go-sdk/mobile
//go:generate go build -tags "local evm java" -x -o java/libshare_signer_test.so -ldflags "-s -w" -buildmode=c-shared  dechain-go-sdk//mobile
//go:generate go.exe build -tags "main evm java" -x -o java/share_signer_main.dll -ldflags "-s -w" -buildmode=c-shared dechain-go-sdk//mobile
//go:generate go.exe build -tags "local evm java" -x -o java/share_signer_test.dll -ldflags "-s -w" -buildmode=c-shared dechain-go-sdk//mobile
func main() {
	fmt.Println("")
}

