---
layout: post
title: Trae AI 日常使用指南
date: 2023-10-24 10:00:00
categories: linux
tags: [trae, ide, 快捷键, python, golang]
---

# 快捷键
* 选中当前单词后，按 Command + D 可依次选中文档中下一个相同的单词
* 要一次性选中所有相同单词，按 Command + Shift + L
* 全局搜索的快捷键是 Command + Shift + F
* 搜索文件快捷键是 Command + P
* AI侧边栏 Command + U
* 跳转到函数定义 Ctrl + 鼠标左键
* 重命名重构 F2
* 面板快捷键 Command + J
* 左边面板快捷键 Command + B
* 跳到指定的行 Ctrl + G
* 打开最近项目 Ctrl + R
* 单词转大写 Command + Shift + U
* 打开设置 Command + ,
* 跳转至文件中的下一个错误 F8


# python
## 配置解释器 
使用.pyenv 添加的新环境不能直接被trae感知到需要
Command + Shift + P -> Python: Select Interpreter 
添加后，后面其它工作区就可以选择到新的解释器了

## 配置PYTHONPATH
配置终端设置 -> Env:Osx -> PYTHONPATH


# golang
## 环境配置
command + shift + p -> go: install/upgrade tools   
选中需要的工具  

### Go 开发工具列表

1. **gotests（未安装）**
   - 作用：自动生成 Go 代码的测试用例（比如为函数、方法生成 `*_test.go` 文件），节省手动编写测试的时间
   - 举例：执行 `gotests -all -w myfile.go`，会为 `myfile.go` 中的所有函数生成对应的单元测试代码

2. **impl（未安装）**
   - 作用：自动生成 Go 接口的实现代码
   - 举例：定义接口 `type Reader interface { Read() string }` 后，运行 `impl 'type MyStruct struct{}' Reader`，会自动生成 `MyStruct` 对 `Reader` 接口的实现方法

3. **goplay（未安装）**
   - 作用：本地启动 Go 代码的 "playground" 环境，支持快速运行/测试小段 Go 代码（类似在线 Go Playground，但在本地运行）

4. **dlv（已安装，版本 v1.26.0）**
   - 作用：Golang 官方调试工具（全称 Delve），是 Go 程序断点调试的核心工具（Trae/GoLand 等 IDE 的调试功能底层基于 dlv）
   - 功能：支持断点、单步执行、查看变量、查看调用栈等调试操作

5. **gopls（已安装，版本 v0.21.0）**
   - 作用：Golang 语言服务器（全称 Go Language Server），是 IDE（如 Trae、VS Code）实现代码补全、语法检查、跳转定义、重构等功能的核心依赖
   - 说明：你在 Trae 中看到的 Go 代码语法高亮、包导入提示等，都是 gopls 在后台提供的支持


## 配置golang项目的运行与调试
点击运行与调试 -> 点击创建launch.json文件 -> 选择go -> 配置launch.json文件    
会在.vscode目录下生成launch.json文件 
```json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "gateway",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/gateway/main.go"
        },
        {
            "name": "product",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/product/main.go"
        },
        {
            "name": "cooperate",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/cooperate/main.go",
            "args": []
        }

    ]
}
```