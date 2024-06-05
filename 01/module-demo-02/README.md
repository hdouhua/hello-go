## Go Module 的一些常规操作

### 从一个简单示例开始

在上一个 module 的基础上修改代码

```shell
-> cp -r module-demo module-demo-02
```

添加包 "github.com/google/uuid" 的引用，并且打印 uuid

```go
import (
	"github.com/google/uuid" // <==
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello module, go module mode")
	logrus.Println(uuid.NewString()) // <==
}
```

编译

```shell
-> go build

# 得到如下错误
main.go:4:2: no required module provides package github.com/google/uuid; to add it:
        go get github.com/google/uuid
```

执行 go get 修复

```shell
-> go get "github.com/google/uuid"
# 
go: downloading github.com/google/uuid v1.6.0
go get: added github.com/google/uuid v1.6.0
```

可以看到 go.mod 文件内已经插入了依赖包导入路径

```go
require (
	github.com/google/uuid  // <==
    // ...
)
```

再次编译、执行代码。

```shell
-> go build -a && ./mymodule01
```

#### 小结

这是一个简单的示例，手动执行 go get 新增依赖项和执行 go mod tidy 自动分析并下载依赖项的最终效果，是等价的。
但对于复杂的项目变更而言，显然 go mod tidy 是更佳的选择。

### 升级/降级依赖的版本

还是以上面提到过的 logrus 为例， logrus 有多少发布版本？

```shell
-> go list -m -versions github.com/sirupsen/logrus
#
github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1 v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4 v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1 v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3 v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.4.2 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.8.1 v1.8.2 v1.8.3 v1.9.0 v1.9.1 v1.9.2 v1.9.3
```

1. go mod tidy 命令默认选择了 v1.9.3 （查看 go.mod 文件知道的），现在想降级到某个之前发布的兼容版本，比如 v1.8.0

   ```shell
   -> go get github.com/sirupsen/logrus@v1.8.0
   #
   go: downloading github.com/sirupsen/logrus v1.8.0
   go: downloading github.com/stretchr/testify v1.2.2
   go: downgraded github.com/sirupsen/logrus v1.9.3 => v1.8.0
   ```

   查看 go.mod 文件，也已经更新了。
   重新编译、执行。

1. 使用编辑 go.mod 和 go mod tidy 命令来完成降级

   ```shell
   -> go mod edit -require=github.com/sirupsen/logrus@v1.7.0
   -> go mod tidy
   #
   go: downloading github.com/sirupsen/logrus v1.7.0
   ```

   重新编译、执行，看看结果，perfect ！

1. 升级至某个补丁，仍然可以使用上面两种方式 go get 或 go mod tidy

   ```shell
   -> go get github.com/sirupsen/logrus@v1.7.1
   #
   go: downloading github.com/sirupsen/logrus v1.7.1
   go: upgraded github.com/sirupsen/logrus v1.7.0 => v1.7.1
   ```

### 添加版本号 > 1 的依赖

按照语义版本规范，如果要为项目引入主版本号大于 1 的依赖，比如 v2.0.0，那么由于这个版本与 v1、v0 开头的包版本都不兼容，在导入 v2.0.0 包时，不能再直接使用 github.com/user/repo，而要使用如下包导入路径：

```go
import github.com/user/repo/v2/xxx // <== 带版本号的包导入路径
```

举例，导入 redis v7 包

```go
import (
	_ "github.com/go-redis/redis/v7" // <==
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)
```

```shell
-> go mod tidy # <== 也可以用 go get
# ...
-> go build -a && ./mymodule01
#
INFO[0000] hello module, go module mode
INFO[0000] 84fe7dea-3d85-4729-b98b-605552a10237
```

可以得到正确的结果。

### 再升级依赖到一个不兼容的版本

```go
import (
	_ "github.com/go-redis/redis/v8" // <==
   //...
)
```

```shell
-> go mod tidy
# ...
-> go build -a && ./mymodule01
#
INFO[0000] hello module, go module mode
INFO[0000] a47d78da-cf33-4f57-a13f-552dc9d7cd90
```

### 移除一个依赖

编辑源代码 main.go ，移除 redis 依赖；
然后用 go list 命令列出当前模块的所有依赖

```shell
-> go list -m all
mymodule01
# ...
github.com/go-redis/redis/v8 v8.11.5
```

发现 redis 依赖仍然在列？！正确的做法是使用 go mod tidy 命令，将这个依赖彻底从 module 构建上下文中清除。

```shell
-> go mod tidy
-> go list -m all
#
mymodule01
github.com/davecgh/go-spew v1.1.1
github.com/google/uuid v1.6.0
github.com/magefile/mage v1.10.0
github.com/pmezard/go-difflib v1.0.0
github.com/sirupsen/logrus v1.7.1
github.com/stretchr/testify v1.2.2
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
```

可以看到 redis 依赖被清除了。

### 特殊应用 vendor

虽然 vendor 机制诞生于 GOPATH 构建模式主导的年代，但在 Go Module 构建模式下，它依旧被保留了下来，并且成为了 Go Module 构建机制的一个很好的补充。**特别是在一些不方便访问外部网络，并且对 Go 应用构建性能敏感的环境，比如在一些内部的持续集成或持续交付环境（CI/CD）中，使用 vendor 机制可以实现与 Go Module 等价的构建。**

与 GOPATH 构建模式不同，**Go Module 构建模式下，再也无需手动维护 vendor 目录下的依赖包了**，Go 提供了可以快速建立和更新 vendor 的命令

```shell
-> go mod vendor # <== 为项目建立 vendor
-> tree -LF 2 vendor
vendor/
├── github.com/
│   ├── google/
│   ├── magefile/
│   └── sirupsen/
├── golang.org/
│   └── x/
└── modules.txt
```

上面 go mod vendor 命令创建了 *vendor 目录*和一份*项目的依赖包的副本*，并且通过 *vendor/modules.txt* 记录 vendor 下的 module 以及版本。

如果要基于 vendor 构建（而不是基于本地缓存的 Go Module 构建），需要在 go build 后面加上 -mod=vendor 参数。

>*在 Go 1.14 及以后版本中，如果 Go 项目的顶层目录下存在 vendor 目录，那么 go build 默认也会优先基于 vendor 构建，除非已指定 go build -mod=mod 。

```shell
-> rm mymodule01 && go build -mod=vendor
```

## 小结

- go get

   可以升级/降级某依赖的版本
- go mod tidy

   自动分析 go 源码的依赖变化——依赖的新增、删除及版本的变更，并更新 go.mod 中的依赖信息
- go mod vendor

   让 go module 构建模式依然支持 vendor 机制，并且可以自动管理 vendor 目录里的缓存的依赖包
