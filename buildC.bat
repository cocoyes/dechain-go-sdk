set GOARCH=amd64
set GOBIN=E:\soft\GOSDK\bin
set GOCHAR=6
set GOEXE=.exe
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=E:\soft\GOPATH
set GORACE=
set GOROOT=E:\soft\GOSDK\
set GOTOOLDIR=E:\soft\GOSDK\pkg\windows_amd64
set CC=gcc
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0
set CXX=g++
set CGO_ENABLED=1

%GOBIN%/go build -x -v -ldflags "-s -w" -buildmode=c-archive -o main.a main.go