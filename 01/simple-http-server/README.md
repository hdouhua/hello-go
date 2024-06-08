## 创建一个简单的 web 服务应用

```shell
» mkdir simple-http-server
» cd simple-http-server
» go mod init simple-http-server
```

Go 提供了完善的工具链和“自带电池”的标准库，大大减少了对第三方库的依赖。以开发 web 服务为例，基于 net/http 包就可以轻松开发了。

[源代码](./main.go)

```shell
» go run main.go
» curl -i localhost:8080
HTTP/1.1 200 OK
Date: Fri, 07 Jun 2024 08:41:16 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello, World%
```
