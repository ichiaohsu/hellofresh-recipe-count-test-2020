
.PHONY: mocks setup build count

SETUP_BUILD_PATH=/tmp/setup-recipe
postcode ?= 10120
delivery-from ?= 11
delivery-to ?= 3
filepath ?= hf_test_calculation_fixtures.json

mocks:
	go get github.com/golang/mock/mockgen
	go generate ./...

setup:
	rm -f $(SETUP_BUILD_PATH)
	go build -o $(SETUP_BUILD_PATH) cmd/setup/main.go
	$(SETUP_BUILD_PATH) --filepath ${filepath}

build:
	 docker build --tag recipe-count .

count:
	 @docker run -t recipe-count --postcode ${postcode} --delivery-from ${delivery-from} --delivery-to ${delivery-to}
