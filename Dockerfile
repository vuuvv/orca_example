FROM registry.aliyuncs.com/vuuvv/orca:0.0.9 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

RUN mkdir publish && cp main publish && cp -r resources publish

FROM registry.aliyuncs.com/vuuvv/base-debian11:latest

COPY --from=builder /app/publish /

ENV GIN_MODE=release

ENTRYPOINT ["./main"]