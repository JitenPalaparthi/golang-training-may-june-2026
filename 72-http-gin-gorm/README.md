# gomock

go install go.uber.org/mock/mockgen@latest


## generate mocks

mockgen -source=internal/repositories/repository.go \
-destination=internal/repositories/mocks/product_db_mock.go \
> -package=mocks

