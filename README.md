
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

```