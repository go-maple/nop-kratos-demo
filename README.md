# nop-kratos-demo

# nop中nacos meta 填写下面内容需要注意
`
`
# 了解一些kratos基础

https://go-kratos.dev/

# clone下代码
git clone git@github.com:go-maple/nop-kratos-demo.git

cd nop-kratos-demo

# 编译服务

### 生成协议
`
go generate ./...

### 生成结果
proto: api/blog/v1/blog.proto
proto: api/blog/v1/error.proto
proto: internal/conf/conf.proto
wire: github.com/go-maple/nop-kratos-demo/blog/cmd/blog: wrote wire_gen.go
proto: api/server/v1/greeter.proto
proto: internal/conf/conf.proto
wire: github.com/go-maple/nop-kratos-demo/server/cmd/server: wrote wire_gen.go
`
## 从根目录进入vite server
`
cd server

make build

./bin/server -conf ./configs/config.yaml

启动成功后日志大概是下面内容
DEBUG msg=config loaded: config_local.yaml format: yaml
INFO msg=watcher's ctx cancel : context canceled
2024-01-12T21:35:24.075+0800    INFO    nacos_client/nacos_client.go:79 logDir:</tmp/nacos/log>   cacheDir:</tmp/nacos/cache>
INFO ts=2024-01-12T21:35:25+08:00 caller=grpc/server.go:212 service.id=glennxudeMacBook-Pro.local service.name=server service.version= trace.id= span.id= msg=[gRPC] server listening on: [::]:9001
`
## 从根目录进入blog
`
cd blog

make build

./bin/blog -conf ./configs/config.yaml

启动成功后大概是下面内容
INFO msg=[resolver] update instances: [{"id":"192.168.30.101#9001#DEFAULT#DEFAULT_GROUP@@server.grpc","name":"DEFAULT_GROUP@@server.grpc","version":"","metadata":{"kind":"grpc","version":""},"endpoints":["grpc://192.168.30.101:9001"]}]
Handle connection.
INFO ts=2024-01-12T21:38:20+08:00 caller=grpc/server.go:212 service.id=glennxudeMacBook-Pro.local service.name=blog service.version= trace.id= span.id= msg=[gRPC] server listening on: [::]:9000
INFO ts=2024-01-12T21:38:20+08:00 caller=http/server.go:317 service.id=glennxudeMacBook-Pro.local service.name=blog service.version= trace.id= span.id= msg=[HTTP] server listening on: [::]:8000
`
#

# 命令行调用
`
## blog调用vite server正常调用
curl --location 'http://localhost:8000/v1/article' \
--header 'Content-Type: application/json' \
--data '{
    "title": "my article",
    "content": "my article content"
}'
`
### 返回结果里面出现 Hello my article content 这里的Hello就是来自vite server添加的内容 my article content是原先提交的内容

`
{"Article":{"id":"0","title":"my article","content":"Hello my article content","like":"0"}}
`

## blog调用vite server,但是vite server接口异常
`
curl --location 'http://localhost:8000/v1/article_err' \
--header 'Content-Type: application/json' \
--data '{
    "title": "my article",
    "content": "my article content"
}'
`
### 返回结果
`
{
    "code": 500,
    "reason": "",
    "message": "rpc error: code = Internal desc = SayHelloError",
    "metadata": {}
}
`

## blog内部接口异常
`
curl --location 'http://localhost:8000/v1/error' \
--header 'Content-Type: application/json' \
--data '{
    "title": "my article",
    "content": "my article content"
}'
`
### 返回结果
`
{
    "code": 500,
    "reason": "CreateError",
    "message": "create Error",
    "metadata": {}
}
`