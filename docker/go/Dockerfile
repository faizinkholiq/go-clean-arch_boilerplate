# golang
FROM golang:1.17.3-buster

# working dir
WORKDIR /app

COPY . .

# auto build & run
RUN go mod vendor
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -u github.com/swaggo/swag/cmd/swag
ENTRYPOINT CompileDaemon --build="go build -o /go/bin ./cmd/main.go" --command="/go/bin/main"
