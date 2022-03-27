FROM golang:alpine

MAINTAINER TY tianyuanyuans@163.com

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux\
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    GOPRIVATE=git.domob-inc.cn

WORKDIR /app/golang_blog
COPY . /app/golang_blog
RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]