# nop-kratos-demo

# 了解一些kratos基础


# clone下代码

# 编译服务

## 从根目录进入vite server
cd server
make build
./bin/server -conf ./configs/config.yaml
## 从根目录进入blog
cd blog
make build
./bin/blog -conf ./configs/config.yaml
#

# 命令行调用
`
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

