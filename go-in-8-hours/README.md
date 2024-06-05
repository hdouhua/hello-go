# 从其他语言快速转入 GO 语言

## 安装

太简单，略过。

## 第一个 Go 程序

```go
os.Args // 命令行参数
os.Exit(0) // 因为 go 的 main 函数没有返回值，如果需要判断程序是否正常退出，可以使用这个
```

### main 函数

- 无参数、无返回值
- main 方法必须在 main 包里
- 如果包名不是 main ，运行时需要先 `go build` ，否则直接使用 `go run` 即可

### 包的声明

- 由字母和下划线组成
- 可以和文件夹不同名
- 同一文件夹下声明一致
- 引入包重名时，可以使用别名，如 `import alias`
- 如果一个引入的包没有使用，那么会报错
- 匿名引用，使用 _

## go 工程目录

- src
- pkg
- bin

## 基本类型

go 不会默认做类型转换（隐式类型转换），类型不同无法编译！

### string

- 双引号，字符串内部引号使用 \ 转义
- 长语句可以使用 "\`\`"
- 字符的长度（字节数）与编码无关，使用 len(str)
- 字符的数量与编码相关
- 字符拼接使用 +

### rune 类型

- 这在别的语言里没有，但 go 里没有 char ，rune 不是 char ，也不是 number ，也不是 byte ！
- 本质上是 int32 ，是字符，一个 rune 4 个字节
- 不是很常用

### bool, int, unit, float

- int8, int16, int32, int64, int
- uint8 ..., uintptr
- float32, float64

### byte 类型

- 本质上是 int8

## 基本语法

### 变量声明

- 变量声明了，没有使用会报错
- 类型不匹配，会报错

#### var —— 变量声明

- var name type = value
- 变量作用域
   - 局部变量
   - 包变量
   - 块声明
- go 支持类型推断
   - 数字被认为是 int 或者 float64
- 使用 驼峰 命名规范
- 变量名首字符的大小写控制访问属性（作用域，可见性）
   - 大写：包外可以访问

#### := —— 变量声明

- 只用于局部变量

### 常量声明 const

参见 var

## 方法声明

- 关键字 func
- 方法名称首字符大小写控制访问属性
- 参数列表 [name type, ...]
- 返回值列表 [name type, ...]
- 连续相同类型的参数可以省略前面的参数的类型，如 `a, b, c int`
- 不定参数 `name... type` 放最后, `args... string`

推荐的方法声明写法：
- 参数列表含有参数名称
- 返回值列表不指定名称

## 方法调用

- 忽略返回值使用 `_`

## 常用函数

### fmt

- 常用的 %s %d %v %+v %#v ，这些对于 go 严格的类型操作，如 string 的拼接是很有帮助的

## 数组和切片

### 数组

语法 [capacity]type

- 初始化时需要指定容量（长度）
- 直接初始化
- 使用 len 和 cap 操作来获取数组的长度

```go
// 直接初始化
a1 := [3]int{3, 2, 1} // 区别于 a1 := []int{3, 2, 1} ，后者是切片
// 等同于 a1 := [...]int{3, 2, 1} // ... 代表数组长度自动推导
fmt.Printf("a1: %v, len: %d, cap: %d, 1st emlement: %d\n", a1, len(a1), cap(a1), a1[0])

var a2 [3]int 
var a2 [3]int // 等同于 var a2 = [3]int{0, 0, 0} 或 var a2 = [...]int{0, 0, 0}
fmt.Printf("a2: %v, len: %d, cap: %d\n", a2, len(a2), cap(a2))
```

### 切片

语法 []type

- 直接初始化
- make 初始化： make([]type, 0, capacity)  <- 推荐写法
- append 追加元素
- len 获取元素数量
- cap 获取切片容量
- 子切片使用 [start:end] 获取 [start, end) 之间的元素（集合前闭后开）
   - 子切片和切片是否会相互影响？了解一下 ShareSlice
- 切片的底层实现也是数组，也就是不支持随机添加/删除操作

### 区别

切片**可扩容**（append），可使用 **make 初始化**

## 指针类型

不支持指针运算
string 是值类型，默认初始化是空字符，而不是 nil

## for 语句

- for {} ，类似 while （go 没有 while 语句）
- fori
- for range

```go
a1 := [...]int{3, 2, 1}
// while 形式
index := 0
for {
    if index >= len(a1) {
        break
    }
    fmt.Printf("a1[%d]: %d\n", index, a1[index])
    index++
}
// fori 形式
for i := 0; i < len(a1); i++ {
    fmt.Printf("a1[%d]: %d\n", i, a1[i])
}
// for range 形式
for i, v := range a1 {
    fmt.Printf("a1[%d]: %d\n", i, v)
}
for _, v := range a1 {
    println(v)
}
```

## swtich 语句

与其他语言的区别就是不需要 break
