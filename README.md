## Platform introduction

Dechain go SDK is a SDK development package developed in go language and connected to dechain. This SDK can be provided to go clients or IOS clients.

## build for ios clients
Before running the following command, you need to support the environment variable of go language, and then run it on the MAC machine. When it runs successfully, it can be generated A and H files and give them to IOS client developers

export PATH=$PATH:/usr/local/go/bin  

`CFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path) CGO_ENABLED=1 GOARCH=arm64 CC="clang $CFLAGS" go build -tags "main evm ios"  -o java/lib_dechain.a -ldflags "-s -w" -buildmode=c-archive dechain-go-sdk/mobile`

