INTEGRATION_TEST_PATH?=./cmd/handler

dependencies:
	go mod download

build-cmd:
	go build -o autorizador cmd/main.go

build-mocks:
	go get github.com/golang/mock/gomock@v1.6.0
	go install github.com/golang/mock/mockgen@v1.6.0
	~/go/bin/mockgen -source=usecase/account/interface.go -destination=usecase/account/mock/account.go -package=mock
	~/go/bin/mockgen -source=usecase/transaction/interface.go -destination=usecase/transaction/mock/transaction.go -package=mock

test:
	go test -tags testing -v  ./...

test-cov:
	go test -coverprofile=cover.txt ./... && go tool cover -html=cover.txt -o cover.html

test-integration:
	go test -tags=integration $(INTEGRATION_TEST_PATH) -v -count=1

run:
	go run -race ./cmd/main.go

build: dependencies build-cmd
