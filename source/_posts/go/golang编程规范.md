---
categories:
- go

tags:
- golang规范
---

# go fmt

大部分的格式问题可以通过 gofmt 解决，gofmt 自动格式化代码，保证所有的Go代码与官方推荐的格式保持一致，于是所有格式有关问题，都以 gofmt 的结果为准。   
代码提交前，必须执行gofmt进行格式化。

# go vet

vet工具可以帮我们静态分析我们的源码存在的各种问题，例如多余的代码，提前return的逻辑，struct的tag是否符合标准等。 代码提交前，必须执行go ver进行静态检查。
<!--more-->

# 长度约定

* 代码块长度，比如约定超过10行就需要考虑优化。
* 行代码长度控制，太长的需换行，提高代码可读性。

# 注释

* 对外（Public）的结构体、函数和包必须进行注释
* 结构体注释格式：

```
// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create.
type ObjectMeta struct {
}
```

* 函数注释格式

```
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
  。。。
}
```

* 包注释格式

```
// Package path implements utility routines for
// manipulating slash-separated filename paths.
path
```

# 命名

* 关键或者复杂的代码需要进行注释：
* 注释内容必须是可读的完整句子，简要并且突出重点。
* 需要注释来补充的命名就不算是好命名。
* 使用可搜索的名称：单字母名称和数字常量很难从一大堆文字中搜索出来。单字母名称仅适用于短方法中的本地变量，名称长短应与其作用域相对应。若变量或常量可能在代码中多处使用，则应赋其以便于搜索的名称。
* 做有意义的区分：Product 和 ProductInfo 和 ProductData 没有区别，NameString 和 Name 没有区别，要区分名称，就要以读者能鉴别不同之处的方式来区分 。
* 函数命名规则：驼峰式命名，名字可以长但是得把功能，必要的参数描述清楚，函数名名应当是动词或动词短语，不可导出的函数以小写开头。
* 如 postPayment、deletePage、save。并依 Javabean 标准加上 get、set、is 前缀。例如：xxx + With + 需要的参数名 + And + 需要的参数名 + …..
* 结构体命名规则：结构体名应该是名词或名词短语，如 Custome、WikiPage、Account、AddressParser

* 接口命名规则： 单个函数的接口名以"er"作为后缀，如Reader,Writer 接口的实现则去掉“er”

```
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

两个函数的接口名综合两个函数名

```
type WriteFlusher interface {
    Write([]byte) (int, error)
    Flush() error
}
```

三个以上函数的接口名，类似于结构体名

```
type Car interface {
    Start([]byte)
    Stop() error
    Recover()
}
```

# 包

* 包名命名规则：包名应该为小写单词，不要使用下划线或者混合大小写。
* 文件夹命名规则：小写单词，使用横杠连接
* 文件命名 规则：小写单词，使用下划线连接，测试文件_test.go结束
* 一般情况下包名和文件夹命名是一致的，不过也可以不一样

# 常量

常量均需使用全部大写字母组成，并使用下划线分词：

```
const APP_VER = "1.0"
```

如果是枚举类型的常量，需要先创建相应类型：

```
type Scheme string
const (
    HTTP  Scheme = "http"
    HTTPS Scheme = "https"
)
```

如果模块的功能较为复杂、常量名称容易混淆的情况下，为了更好地区分枚举类型，可以使用完整的前缀：

```
type PullRequestStatus int
const (
    PULL_REQUEST_STATUS_CONFLICT PullRequestStatus = iota
    PULL_REQUEST_STATUS_CHECKING
    PULL_REQUEST_STATUS_MERGEABLE
)

```

# 变量

变量名称一般遵循驼峰法，但遇到特有名词时，需要遵循以下规则：

* 单词本身为缩写，使用全大写
* 如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient
* 其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID
* 错误示例：UrlArray，应该写成 urlArray 或者 URLArray

若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头

```
var isExist boolvar hasConflict bool
var canManage bool
var allowGitHook bool
```

# struct

声明和初始化采用多行，初始化结构体使用带有标签的语法

```
type User struct{
    Username  string
    Email     string
}
 
u := User{
    Username: "yourname",
    Email:    "yourname@gmail.com",
}
```

修改对象属性不能直接使用赋值，要写成方法且必须加锁

# map

非线程安全，并发读写map的情况下必须加锁，不然会产生panic    
可使用sync.Map 替代。

# 函数

函数采用命名的多值返回，传入变量和返回变量以小写字母开头

```
func nextInt(b []byte, pos int) (value, nextPos int)
```

函数返回值可能为空或零值时，最好加一个逻辑判断的返回值

```
func Foo(a int, b int) (string, bool)
```

# init

* 在同一个文件中，可以重复定义init方法
* 在同一个文件中，多个init方法按照在代码中编写的顺序依次执行
* 在同一个package中，可以多个文件中定义init方法
* 在同一个package中，不同文件中的init方法的执行按照文件名先后执行各个文件中的init方法

建议同一个文件中只定义一个init方法，同一个package中init尽量合并

# 错误处理

* error作为函数的值返回,必须对error进行处理
* 错误描述如果是英文必须为小写，不需要标点结尾
* 采用独立的错误流进行处理

采用下面的写法

```
if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code
```

使用函数的返回值时，则采用下面的方式

```
var (
  x string
  err error
)

if x, err = f();err != nil {
    // error handling
    return
}
// use x
```

# 控制结构

# panic

在逻辑处理中不要使用panic，且业务逻辑中要有recover机制

# import

对 import 的包进行分组管理，用换行符分割，而且标准库作为分组的第一组。如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，建议采用如下方式进行组织你的包。

```
 package main
 
import (
    "fmt"
    "os"
    
    "code.google.com/a"
    "github.com/b"
 
    "kmg/a"
    "kmg/b"
)


```

# 参数传递

* 对于少量数据，不要传递指针
* 对于大量数据的 struct 可以考虑使用指针
* 传入的参数是 map，slice，chan 不要传递指针，因为 map，slice，chan 是引用类型，不需要传递指针的指针

# 单元测试

# 日志

* 为了方便日志分析，记录日志统一使用
* 记录有意义的日志，日志里记录一些比较有意义的状态数据：程序启动，退出的时间点；程序运行消耗时间；耗时程序的执行进度；重要变量的状态变化。
* 日志内容必须是可读的英文语句，第一个单词首字母大写，合理标点符号

# 依赖包

go module