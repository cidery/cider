FROM golang

WORKDIR /go/src/github.com/cidery/cider

COPY . .

RUN go build -o build/server cmd/server/main.go
RUN chmod +x build/server

ENTRYPOINT ["/go/src/github.com/cidery/cider/build/server"]
