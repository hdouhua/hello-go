## GOPATH

```shell
» go build main.go

main.go:3:8: no required module provides package github.com/sirupsen/logrus: go.mod file not found in current directory or any parent directory; see 'go help modules'

```

要解决这个问题， `go get` 该登场了。

```shell
» go get github.com/sirupsen/logrus
# 
go: go.mod file not found in current directory or any parent directory.
        'go get' is no longer supported outside a module.
        To build and install a command, use 'go install' with a version,
        like 'go install example.com/cmd@latest'
        For more information, see https://golang.org/doc/go-get-install-deprecation
        or run 'go help get' or 'go help install'.

» go install github.com/sirupsen/logrus@latest
#
package github.com/sirupsen/logrus is not a main package
```

目前本机的 go 版本是 1.23 ，看起来是无法使用 GOPATH 模块了。
详细，请移步[扩展部分](#扩展--go-get-vs-go-install)。

### vendor 机制

为了支持可重现构建， Go 1.5 版本引入了 vendor 机制，开发者可以在项目目录下缓存项目的所有依赖，实现可重现构建。但 vendor 机制依旧不够完善，开发者还需要手工管理 vendor 下的依赖包。

Go Module 机制下的vendor，[详见](../module-demo-02/README.md#特殊应用-vendor)

## 扩展 —— go get vs go install

`go get` ：主要用于管理依赖项，修改 go.mod 文件，并下载依赖到模块缓存中。

在 Go 1.16 及以后的版本中，go get 主要用于修改 go.mod 文件，即管理依赖项。下载所指定的依赖项到本地模块缓存中（位于 $GOPATH/pkg/mod），不会在 $GOPATH/bin 或 ./bin 目录下安装任何可执行文件。

`go install` ：用于编译并安装 Go 程序到 bin 目录下。

当运行 go install 时，它会构建并安装指定的包或命令（如果指定的是模块路径下的命令）到 bin 目录下。默认情况下，bin 目录是 $GOPATH/bin，但如果在模块模式下，它可以是模块下的 ./bin 目录。
