---
title: python3高级核心技术97讲
categories:
- python

 
tags:
- python高级
---

# python一切皆对象

## python一切皆对象
python的面向对象更加彻底  
函数和类也是对象，属于python的一等公民

## type,object和class的关系
1. type->class->obj
2. object是顶层基类
3. type的基类也是object

![](python3高级核心技术97讲/img.png)

## python常见内置类型
对象的三个特征：身份（内存地址），类型，值  

* None 全局只有一个
* 数值 int,float,complex,bool
* 迭代类型
* 序列类型：list,bytes(bytearray,memoryview),range,tuple,str,array
* 映射 dict
* 集合 set,frozenset
* 上下文管理类型 with
* 其它
    * 模块类型
    * class和实例
    * 函数类型
    * 方法类型
    * 代码类型
    * object对象
    * type类型
    * ellipsis类型（省略号）
    * notimplemented类型


# 魔法函数
## 什么是魔法函数
魔法函数用于扩展类的特性  

## 魔法函数一览
### 非数学运算
* 字符串表示 __repr__,__str__
* 集合序列相关 __len__,__getitem__,__setitem__,__delitem__,__contains__  
* 迭代相关 __iter__,__next__
* 可调用 __call__
* with上下文管理器  __enter__,__exit__
* 数值转换 __abs__,__bool__,__int__,__float__,__hash__,__index__
* 元类相关 __new__,__init__
* 属性相关 __getattr__,__setattr__,__getattribute__,__setattribute__,__dir__
* 属性描述符 __get__,__set__,__delete__
* 协程 __await__,__aiter__,__anext__,__aenter__,__aexit__

### 数学运算
。。。


# 深入python类和对象
## 鸭子类型和多态
一个东西，看起来像鸭子，比如会游泳，那就可以这个东西认为是鸭子。  
python 本身是属于动态语言，一个变量可以存储不同的类型，即一个变量的类型只有在执行到当前位置的时候才能知道他的类型。  

## 抽象基类
抽象基类的两个作用：  
1. 使用isinstance进行判定
2. 强制限制子类实现某些方法   

抽象基类在python中主要不是用来集成的，我们在使用的时候可以使用python的鸭子类型来替换抽象基类的使用场景。  

相关模块：abc

## isinstance和type的区别  
* isinstance 用于判定类型，包括继承关系
* type 用于判定类型，不包括继承关系

关键词 is 和 ==的区别， is 用于取类型的id（python中，每个对象都有对应的id，class这些类型属于全局对象,只有一个id）, == 用于判定值是否相等  

> 注意使用isinstance 判定类型，而不是type

## 类变量和实例变量  
实例共享类变量

## 类属性和实例属性以及查找顺序
C3算法  
* 非菱形继承，深度优先
* 菱形继承，广度优先

通过 类.__mro__ 可以查询查找中顺序

> python2中如果类不显示继承object，那么就不会继承，而python3中会默认的继承object, 分别叫做经典类和新式类


## 静态方法，类方法，实例方法
* staticmethod 静态方法，无需类参数
* classmethod 类方法，第一个参数是类
* 实例方法的第一个参数是实例


## 数据封装和私有属性
私有属性通过__开头   
python中没有绝对的私有属性，可以通过类似以下的方式访问到私有属性  
```python
class User:
    def __init__(self):
        self.__a=1

user=User()
print(user._User__a)
```

## python对象的自省机制  
自省是通过一定的机制查询到对象的内部结构  
* __dict__ 以k,v的方式展示对象的属性以及属性值,不包括方法
* dir 函数返回对象的所有属性，包含方法。

## super函数
super函数用于调用父类的函数  
super函数的调用顺序为当前类的__mro__顺序   

## django rest framework 中多继承的使用经验  

### mixin 模式
1. mixin类功能单一
2. 不和基类关联，可以和任意组合关联
3. 在mixin中不要使用super这种用法

## python with语句

## contextlib简化上下文管理器
```python
import contextlib

@contextlib.contextmanager
def file_open(name):
  print("a")
  yield 
  print("b")

with file_open("") as f:
  print("c")
```

# 自定义序列类
## python中的序列类型
### 是否能够存储不同数据类型
* 容器序列：list,tuple,deque
* 扁平序列：str,bytes,bytearray,array.array

### 序列是否可变
* 可变序列：list,deque,bytearray,array
* 不可变序列：str,tuple,bytes

## 序列的abc继承关系
abc模块定义了可变序列以及不可变序列的协议（类必须实现的方法）

## 序列的+，+=和extend的区别
* + 号的两边都必须是list
* extend 接受一个迭代类型，将其一一放入列表中
* += 本质上是嗲用extend 

> append 是加入一个元素到到列表，和extend不同


## 可切片的对象

## bisect维护已排序序列