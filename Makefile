BUILD_TAGS = netgo
PACKAGES=$(shell go list ./dao ./service ./monitor/.)

build_monitor:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/monitor.exe cmd/greenfield-data-marketplace-monitor/main.go
else
	go build $(BUILD_FLAGS) -o build/monitor cmd/greenfield-data-marketplace-monitor/main.go
endif

#build_monitor_docker:
#	docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -f monitor.dockerfile -t monitor .

build_server:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/server.exe cmd/greenfield-data-marketplace-server/main.go
else
	go build $(BUILD_FLAGS) -o build/server cmd/greenfield-data-marketplace-server/main.go
endif

#build_server_docker:
#	docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -f server.dockerfile -t server .

build: build_server build_monitor

#build_docker: build_server_docker build_monitor_docker

test:
	go test -race -v $(PACKAGES)

lint:
	golangci-lint cache clean
	golangci-lint run

genswagger:
	swagger generate server -f ./swagger.yaml -A greenfield-data-marketplace --default-scheme=http

genabi:
	./script/abigen --abi ./monitor/contracts/marketplace.abi --pkg contracts --out ./monitor/contracts/marketplace.go --type Marketplace


.PHONY: build build_docker test lint genswagger genabi
