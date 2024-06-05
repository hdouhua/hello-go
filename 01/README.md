# GO 语言学习

Go语言是一种静态类型、编译型语言；
内置了并发支持和函数式编程特性。

## 安装

- 特别简单，下载安装包，然后安装，或者下载二进制包，然后解压

   以 Apple Silicon 为例

   ```shell
   wget https://dl.google.com/go/go1.22.3.darwin-arm64.tar.gz
   tar -xf go1.22.3.darwin-arm64.tar.gz -C ~/Documents/tools
   ```

- 配置 环境变量

   ```shell
   # 这些是可选的
   #export GO111MODULE # Go 1.11
   export GOROOT=~/Documents/tools/go
   export GOPATH=~/workspace/hello-go
   #export GOPROXY=https://goproxy.cn # 国内需要使用

   # 在 PATH 中添加 go 可执行程序路径
   export PATH=$PATH:$GOROOT/bin
   ```

## 第一个 Go 程序

以一个简单的程序来看看 Go 程序的基本结构：

- 包是 Go 语言的基本组成单元，通常使用单个的小写单词命名，一个 Go 程序本质上就是一组包的集合。
- main 包在 Go 中是一个特殊的包，整个 Go 程序中仅允许存在一个名为 main 的包。
- main 包中的主要代码是一个名为 main 的函数。这里的 main 函数会比较特殊：当运行一个可执行的 Go 程序的时候，所有的代码都会从这个入口函数开始运行。

[源文件](./src/01/helloworld.go)

执行

```shell
# 编译
go build helloworld.go
# 执行
./helloworld

# 或者直接运行源码
go run helloworld.go
```

源文件命名原则：

- Go 源文件总是用全小写字母形式的短小单词命名，并且以.go 扩展名结尾。
- 如果要在源文件的名字中使用多个单词，通常直接是将多个单词连接起来作为源文件名，而不是使用其他分隔符，比如下划线。
- 总体来说，尽量不要用两个以上的单词组合作为文件名，否则就很难分辨了。


标准 Go 代码风格
- 使用 Tab 而不是空格来实现缩进

   Gofmt 是 Go 语言在解决规模化（scale）问题上的一个最佳实践，提交代码前使用 Gofmt 格式化 Go 源码。

- 在 Go 语言中，只有首字母为大写的标识符才是导出的（Exported），才能对包外的代码可见；如果首字母是小写的，那么就说明这个标识符仅限于在声明它的包内可见。
- main 包是不可以像标准库 fmt 包那样被导入（Import）的，如果导入 main 包，在代码编译阶段会收到一个 Go 编译器错误：import “xx/main” is a program, not an importable package。
- 不需要用分号“;” 来标识语句的结束？

   虽然 Go 语言的正式语法规范使用分号“;”作为语句的结束，但在实践中，大多数分号是可选的，并且会被编译器自动插入，因此通常在代码中省略不写。


## 第一个 Go 模块

```shell
mkdir hellomodule
cd hellomodule
vi main.go
#...

go build main.go
# 抛出错误
main.go:4:2: no required module provides package github.com/valyala/fasthttp: go.mod file not found in current directory or any parent directory; see 'go help modules'
main.go:5:2: no required module provides package go.uber.org/zap: go.mod file not found in current directory or any parent directory; see 'go help modules'

# 初始化 module
go mod init hellomodule
# 查看 module 目录下多了 go.mod 文件

# 执行 go mod tidy
go mod tidy
# 查看 module 目录下还有 go.sum 文件

# 再执行 go build
# ...
# 执行程序
# ...
```

Go module 构建模式是在 Go 1.11 版本正式引入的，为的是彻底解决 Go 项目复杂版本依赖的问题，现在 Go module 已经成为了 Go 默认的包依赖管理机制和 Go 源码构建机制。

>*从 Go 1.11 到 Go 1.16 版本，不同的 Go 版本在 GO111MODULE 为不同值的情况下，构建模式几经变化，直到 Go 1.16 版本，Go Module 构建模式成为了默认模式。*

Go Module 的核心是一个名为 go.mod 的文件，在这个文件中存储了这个 module 对第三方依赖的全部信息。

一个 module 就是一个包的集合，这些包和 module 一起打版本、发布和分发。

go.mod 所在的目录—— module 的根目录。

```go
module hellomodule

go 1.22.3
```

这时除了按提示手动添加外，也可以使用 go mod tidy 命令，让 Go 工具自动添加依赖。

go.mod 已经记录了 hellomodule 直接依赖的包的信息。与此同时 hellomodule 目录下还多了一个名为 go.sum 的文件，这个文件记录了 hellomodule 的直接依赖和间接依赖包的相关版本的 hash 值，用来校验本地包的真实性。在构建的时候，如果本地依赖包的 hash 值与 go.sum 文件中记录的不一致，就会被拒绝构建。

## Go 项目的结构布局

### 典型的项目结构

```shell
$tree -F exe-layout 
exe-layout
├── cmd/
│   ├── app1/
│   │   └── main.go
│   └── app2/
│       └── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── pkga/
│   │   └── pkg_a.go
│   └── pkgb/
│       └── pkg_b.go
├── pkg1/
│   └── pkg1.go
├── pkg2/
│   └── pkg2.go
└── vendor/
```

这个项目典型布局就是“脱胎”于 Go 创世项目的最新结构布局
- cmd 目录就是存放项目要编译构建的可执行文件对应的 main 包的源文件。如果你的项目中有多个可执行文件需要构建，每个可执行文件的 main 包单独放在一个子目录中，比如图中的 app1、app2，cmd 目录下的各 app 的 main 包将整个项目的依赖连接在一起。
- 通常来说，main 包应该很简洁。我们在 main 包中会做一些命令行参数解析、资源初始化、日志设施初始化、数据库连接初始化等工作，之后就会将程序的执行权限交给更高级的执行控制对象。
- pkgN 目录，这是一个存放项目自身要使用、同样也是可执行文件对应 main 包所要依赖的库文件，同时这些目录下的包还可以被外部项目引用。
- go.mod 和 go.sum ，它们是 Go 语言包依赖管理使用的配置文件。
- vendor 目录。vendor 是 Go 1.5 版本引入的用于在项目本地缓存特定版本依赖包的机制，在 Go Modules 机制引入前，基于 vendor 可以实现可重现构建，保证基于同一源码构建出的可执行程序是等价的。现在 vendor 目录视为一个可选目录。因为 Go Module 本身就支持可再现构建，而无需使用 vendor。

### 多个 module 的项目结构

```shell
$tree multi-modules
multi-modules
├── go.mod // mainmodule
├── module1
│   └── go.mod // module1
└── module2
    └── go.mod // module2
```

可以通过 git tag 名字来区分不同 module 的版本。
其中 vX.Y.Z 形式的 tag 名字用于代码仓库下的 mainmodule；
而 module1/vX.Y.Z 形式的 tag 名字用于指示 module1 的版本；
同理，module2/vX.Y.Z 形式的 tag 名字用于指示 module2 版本。

### 仅有一个可执行程序的项目结构

```shell
$tree -F -L 1 single-exe-layout
single-exe-layout
├── go.mod
├── internal/
├── main.go
├── pkg1/
├── pkg2/
└── vendor/
```

删除了 cmd 目录，将唯一的可执行程序的 main 包就放置在项目根目录下。

## Go 库项目结构布局

### 典型的库项目结构

```shell
$tree -F lib-layout 
lib-layout
├── go.mod
├── internal/
│   ├── pkga/
│   │   └── pkg_a.go
│   └── pkgb/
│       └── pkg_b.go
├── pkg1/
│   └── pkg1.go
└── pkg2/
    └── pkg2.go
```

库类型项目相比于 Go 可执行程序项目的布局要简单一些。因为这类项目不需要构建可执行程序，所以去除了 cmd 目录。
vendor 也不再是可选目录了。对于库类型项目而言，不推荐在项目中放置 vendor 目录去缓存库自身的第三方依赖，库项目仅通过 go.mod 文件明确表述出该项目依赖的 module 或包以及版本要求即可。

### 仅有一个包的 Go 库项目结构

进一步简化为

```shell
$tree -L 1 -F single-pkg-lib-layout
single-pkg-lib-layout
├── feature1.go
├── feature2.go
├── go.mod
└── internal/
```

将这唯一包的所有源文件放置在项目的顶层目录下（比如上面的 feature1.go 和 feature2.go）。

## Go 包依赖管理

### GO 构建模式的演进

- [01](./gopath-demo/)
- [02](./module-demo/)

---

项目所依赖的包有很多版本，Go Module 是如何选出最适合的那个版本的呢？

### 语义导入版本机制 (Semantic Import Versioning)

go.mod 的 require 段中依赖的版本号，都符合 v**X.Y.Z** 的格式。

参考 [语义版本](https://semver.org/)
- 语义版本号分成 3 部分：主版本号（major）、次版本号（minor）和补丁版本号（patch）。
- 按照语义版本规范，主版本号不同的两个版本是相互不兼容的。在主版本号相同的情况下，次版本号大都是向后兼容次版本号小的版本。补丁版本号也不影响兼容性。

Go Module 规定：如果同一个包的新旧版本是兼容的，那么它们的包导入路径应该是相同的。
那么有一天发布新的不兼容的版本时，它的包导入路径就要不同，那要怎么做？ Go Module 创新性地给出了一个方法：将包主版本号引入到包导入路径中。如：

```go
import "github.com/sirupsen/logrus"
// 和
import "github.com/sirupsen/logrus/v2"
```

这就是 Go 的“语义导入版本”机制，也就是*通过在包导入路径中引入主版本号的方式，来区别同一个包的不兼容版本*。
甚至可以同时依赖一个包的两个不兼容版本，如：

```go
import (
    "github.com/sirupsen/logrus"
    logv2 "github.com/sirupsen/logrus/v2"
)
```

### 最小版本选择机制 (Minimal Version Selection) 

![](https://static001.geekbang.org/resource/image/49/1b/49eb7aa0458d8ec6131d9e5661155f1b.jpeg?wh=400:225)

在这张图中，myproject 有两个直接依赖 A 和 B，A 和 B 有一个共同的依赖包 C，但 A 依赖 C 的 v1.1.0 版本，而 B 依赖的是 C 的 v1.3.0 版本，并且此时 C 包的最新发布版为 C v1.7.0。这个时候，Go 命令是如何为 myproject 选出间接依赖包 C 的版本呢？

理想状态下，语义版本控制被正确应用，那么 Go 包依赖管理工具都会选择依赖项的“最新最大 (Latest Greatest) 版本”，对应到图中的例子，这个版本就是 v1.7.0。

但 Go 设计者另辟蹊径，在诸多兼容性版本间，他们不光要考虑最新最大的稳定与安全，还要尊重各个 module 的述求：A 明明说只要求 C v1.1.0 ，B 明明说只要求 C v1.3.0 。所以 **Go 会在该项目依赖项的所有版本中，选出符合项目整体要求的“最小版本”**。

这个例子中，C v1.3.0 是符合项目整体要求的版本集合中的版本最小的那个，于是 Go 命令选择了 C v1.3.0，而不是最新最大的 C v1.7.0。
并且，**Go 团队认为“最小版本选择”为 Go 程序实现<u>持久的和可重现的构建</u>提供了最佳的方案**。

### 总结

Go Module 构建模式将成为 Go 语言唯一的标准构建模式。从现在开始就彻底抛弃 GOPATH 构建模式，全面使用 Go Module 构建模式！

## Go Module 一些常规操作

[详见](./module-demo-02/)

## 参考

- https://github.com/golang-standards/project-layout
