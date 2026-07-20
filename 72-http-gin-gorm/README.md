# gomock

go install go.uber.org/mock/mockgen@latest


## generate mocks

mockgen -source=internal/repositories/repository.go \
-destination=internal/repositories/mocks/product_db_mock.go \
> -package=mocks

## docker build 

```bash
docker build . -t my-go-app:v0.0.1

docker run -d --name=demo-app2 -p 8085:8080 --network my-app-network -e DSN_URL="host=db user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai" my-go-app:v0.0.1
```

## build and push docker iamge

```bash
docker build . -t jpalaparthi/demo-go-app:v0.1.0

docker push jpalaparthi/demo-go-app:v0.1.0
```

## docker and docker compose

```bash
# docker compose commands
   docker compose up --build -d
   clear
   docker compose logs

   docker compose logs app
   clear
   # test the application
   curl localhost:8080/v1/api/public/ping
   curl localhost:8080/v1/api/public/health

   docker images
```

```curl
# create a product
curl --location 'localhost:8080/v1/api/product' \
--header 'username: admin' \
--header 'password: admin123' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Macbook Pro",
    "description": "Apple Macbook Pro M1 chip set",
    "price": 100123.12,
    "stock":100
}'
```


## Protocol Buffer and gRPC

```bash
# install protoc compiler 
https://protobuf.dev/installation/ 
# 1. Install the Go code generation plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 2. Install the gRPC service generation plugin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Compile protoc and generate Go code

protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       internal/protos/product.proto
