FROM golang:1.17 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

RUN mkdir publish && cp main publish && cp -r resources publish

FROM scratch

COPY --from=builder /dist/publish /

ENV GIN_MODE=release

ENTRYPOINT ["./main"]