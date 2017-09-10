APP?=cider-server
RELEASE?=1.0.0
GOOS?=linux

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: check
check: prepare_metalinter
	gometalinter --vendor ./...

.PHONY: build
build: clean
	CGO_ENABLED=0 GOOS=${GOOS} go build \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}" \
		-o build/${GOOS}/${APP} cmd/cider-server/main.go

.PHONY: clean
clean:
	@rm -f build/${GOOS}/${APP}

HAS_SWAGGER := $(shell command -v swagger;)
HAS_DEP := $(shell command -v dep;)
HAS_METALINTER := $(shell command -v gometalinter;)

.PHONY: gen
gen: prepare_swagger
	swagger generate server -f spec/swagger.json -A cider -t src/api --exclude-main

.PHONY: vendor
vendor: prepare_dep
	dep ensure

.PHONY: prepare_swagger
prepare_swagger:
ifndef HAS_SWAGGER
	go get -u -v -d github.com/go-swagger/go-swagger/cmd/swagger && \
	go install -v github.com/go-swagger/go-swagger/cmd/swagger
endif

.PHONY: prepare_dep
prepare_dep:
ifndef HAS_DEP
	go get -u -v -d github.com/golang/dep/cmd/dep && \
	go install -v github.com/golang/dep/cmd/dep
endif

.PHONY: prepare_metalinter
prepare_metalinter:
ifndef HAS_METALINTER
	go get -u -v -d github.com/alecthomas/gometalinter && \
	go install -v github.com/alecthomas/gometalinter && \
	gometalinter --install --update
endif