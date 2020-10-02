# 构建：使用golang:1.13版本
FROM golang:1.13 as build

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作区
WORKDIR /go/release

# 把全部文件添加到/go/release目录
ADD . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build 

# 运行：使用scratch作为基础镜像
FROM scratch as prod

# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/mahjim /

# 启动服务
CMD ["/mahjim"]
