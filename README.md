
## To start, write, run go application

```bash

mkdir demo && cd demo 

go mod init demo 

go run main.go

# Compiles, Assembles and Links --> Binary

# to know where the binary is created 

go run --work main.go

# can also give 

go run .

```

## Build go application

```bash
go build . 

# file name
go build -o app .

# cross compile and build 

GOOS=windows GOARCH=amd64 go build -o build/demo-win-amd64.exe .

```

## go environment variables

```bash
go env 
```

## TO clearn package cache 

```bash
go clean -modcache
```

## go environment variables

- GOROOT --> It is a factory instalation of Go

- GOPATH --> It is a thirdpay library cashe, all thirdparty libraries or custom libraries come from here

- GOBIN  --> all tools, are installed in GOBIN, if GOBIN is empty the GOBIN location is $(GOPATH)/bin

```bash
go tool dist list
```

- GOOS    -> Operating System name

- GOARCH  -> Chipset arch

## Build process

    - Go source --> go compiler --> SSA(Optimizations|Constants|dead code elimation) --> assembler(plan9 assembler asm) --> .o file -> linker -> link all required files, memory mappings --> bin (static binary)

    - java source --> java compiler --> java byte code --> JVM -> loads the byte code --> JIT compillation --> Executes
    - scala source --> scale compiler --> java byte code --> JVM -> loads the byte code --> JIT compillation --> Executes
    - closure source --> closure compiler --> java byte code --> JVM -> loads the byte code --> JIT compillation --> Executes
    - kotlin source --> kotlin compiler --> java bute code --> JVM -> loads the byte code --> JIT compillation --> Executes

    - C# souce --> csharp compiler --> IL code --> .NET Runtime --> loads IL code --> JIT compilation -> Executes
    - F# souce --> fsharp compiler --> IL code --> .NET Runtime --> loads IL code --> JIT compilation -> Executes

    - Rust source --> Rust compiler --> [LLVM [ LLVM FE LLVM BE] --> LLVM IR(FE) --> LLBM BE -> Assemble --> Links] --> Bin


## go compile asn link

```bash
# if the below is not running , run the below command 
# sudo GODEBUG=installgoroot=all go install std

go tool compile -o demo.o main.go  
go tool link -o demo demo.o

```

## debug build vs release build --> nm tool


```bash

go build -o demo main.go ## debug build
# how to know it is debug build

go tool nm  demo # it gives symbol information , symbols are all types, functions, methods information in the binary

## how to do the release build
go build -ldflags="-s -w" -o demo-slim main.go
```
