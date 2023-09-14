---
title: golangci-lint

categories:
- go


tags:
- goland激活
---
# 静态代码检查
静态代码检查是一个老生常态的问题，它能很大程度上保证代码质量。Go 语言自带套件为我们提供了静态代码分析工具 vet，它能用于检查 go 项目中
可以通过编译但仍可能存在错误的代码。静态代码检查是一个老生常态的问题，它能很大程度上保证代码质量。Go 语言自带套件为我们提供了静态代码分
析工具 vet，它能用于检查 go 项目中可以通过编译但仍可能存在错误的代码。
<!--more-->


以上谈到的工具，我们可以称之为 linter。在维基百科是如下定义 lint 的：
> 在计算机科学中，lint 是一种工具程序的名称，它用来标记源代码中，某些可疑的、不具结构性（可能造成 bug）的段落。它是一种静态程序分析工具，
> 最早适用于 C 语言，在 UNIX 平台上开发出来。后来它成为通用术语，可用于描述在任何一种计算机程序语言中，用来标记源代码中有疑义段落的工具。

而在 Go 语言领域， golangci-lint 是一个集大成者的 linter 框架。它集成了非常多的 linter，包括了上文提到的几个，合理使用它可以帮助我们
更全面地分析与检查 Go 代码。golangci-lint 所支持的 linter 项可以查看页面 https://golangci-lint.run/usage/linters/#golint



# 使用 golangci-lint

## 下载
```shell
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

golangci-lint version

```
可以看出，golangci-lint 框架支持的 linter 非常全面，它包括了 bugs、error、format、unused、module 等常见类别的分析 linter。

## 使用
golangci-lint run 它等效于 golangci-lint run ./...  

golangci-lint 可以通过 -E/--enable 去开启指定 linter，或者 -D/--disable 禁止指定 linter。  
```shell
# 除了 errcheck 的 linter，禁止其他所有的 linter 生效。
golangci-lint run --disable-all -E errcheck

# golangci-lint 还可以通过 -p/--preset 指定一系列 linter 开启。
golangci-lint run -p bugs -p error
```

## 配置文件
当然，如果我们要为项目配置 golangci-lint，最好的方式还是配置文件。golangci-lint 在当前工作目录按如下顺序搜索配置文件。  
* .golangci.yml
* .golangci.yaml
* .golangci.toml
* .golangci.json

在 golangci-lint 官方文档 https://golangci-lint.run/usage/configuration/#config-file 中，提供了一个示例配置文件，非常地详细，
在这其中包含了所有支持的选项、描述和默认值。
```yaml
linters-settings:
    errcheck:
        check-type-assertions: true

goconst:
    5    min-len: 2
    6    min-occurrences: 3

    gocritic:
        enabled-tags:
          - diagnostic
          - experimental
          - opinionated
          - performance
          - style
    gove:
        check-shadowing: true
    nolintlint:
        require-explanation: true
        require-specific: true

linters:

  disable-all: true
  enable:
      - bodyclose
      - deadcode
      - depguard
      - dogsled
      - dupl
      - errcheck
      - exportloopref
      - exhaustive
      - goconst
      - gocritic
      - gofmt
      - goimports
      - gomnd
      - gocyclo
      - gosec
      - gosimple
      - govet
      - ineffassign
      - misspell
      - nolintlint
      - nakedret
      - prealloc
      - predeclared
      - revive
      - staticcheck
      - structcheck
      - stylecheck
      - thelper
      - tparallel
      - typecheck
      - unconvert
      - unparam
      - varcheck
      - whitespace
      - wsl


run:
  issues-exit-code: 1
```

## goland中配置golangci-lint
1. go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  （本地已下载则跳过）
1. Plugins 中安装 Go Linter
2. Tools 中点击Go Linter 设置执行路径


# 使用 pre-commit hook
在项目开发中，我们都会使用到 git，因此我们可以将代码静态检查放在一个 git 触发点上，而不用每次写完代码手动去执行 golangci-lint run 命令。
这里，我们就需要用到 git hooks。

## git hooks
git hooks 是 git 的一种钩子机制，可以让用户在 git 操作的各个阶段执行自定义的逻辑。git hooks 在项目根目录的 .git/hooks 下面配置，
配置文件的名称是固定的，实质上就是一个个 shell 脚本。根据 git 执行体，钩子被分为客户端钩子和服务端钩子两类。


客户端钩子包括：pre-commit、prepare-commit-msg、commit-msg、post-commit 等，主要用于控制客户端 git 的提交工作流。
服务端钩子：pre-receive、post-receive、update，主要在服务端接收提交对象时、推送到服务器之前调用。

> 注意，以 .sample 结尾的文件名是官方示例，这些示例脚本是不会执行的，只有重命名后才会生效（去除 .sample 后缀）。

而 pre-commit 正如其名一样，它在 git add 提交之后，运行 git commit 时执行，脚本执行没报错就继续提交，反之就驳回提交的操作。

## pre-commit
试想，如果我们同时开发多个项目，也许项目的所采用的的编程语言并不一样，那么它们所需要的 git hooks 将不一致，此时我们是否要手动给每个项目
都配置一个单独的 pre-commit 脚本呢，或者我们是否要去手动下载每一个钩子脚本呢。

实际上，并不需要这么麻烦。这里就引出了 pre-commit 框架，它是一个与语言无关的用于管理 git hooks 钩子脚本的工具

### 安装
```shell
pip install pre-commit
#或者
curl https://pre-commit.com/install-local.py | python -
#或者
brew install pre-commit
```
测试安装结果  
```shell
pre-commit --version
pre-commit 2.20.0
```

### 编写配置文件
首先我们在项目根目录下新建一个 .pre-commit-config.yaml 文件，这个文件我们可以通过 pre-commit sample-config 得到最基本的配置模板，
通过 pre-commit 支持的 hooks 列表 https://pre-commit.com/hooks.html 中，我们找到了 golangci-lint。  

因此，使用 golangci-lint 的 .pre-commit-config.yaml 配置内容如下  
```yaml
repos:
  - repo: https://ghproxy.com/https://github.com/golangci/golangci-lint
    rev: v1.50.1 # the current latest version
    hooks:
      - id: golangci-lint
```

初始化工具   
```shell
pre-commit run -a
```