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

```python
class A:
  pass
class B:
  def __init__(self):
    # python2语法
    super(B,self).__init__()
    
    # python3中可以直接调用super()
    super().__init__()
```

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
用于向列表插入数据，并且维护列表的顺序性

## 什么时候我们不应该使用列表
* array和list的区别，array只能存储指定类型的值
* dqueue

## 列表推导式，生成器表达式，字典推导式
### 列表推导式
```python
# 提取1-20之间的奇数
odd_list=[i for i in range(21) if i%2==1]

# 复杂情况
def handle_item(item):
    return item*item

odd_list=[handle_item(i) for i in range(21) if i%2==1]
```
逻辑简单可以用列表生产式，如果逻辑过于复杂不建议用列表生成式  

### 生成器表达式
```python
odd_gen=(i for i in range(21) if i%2==1)
odd_list=list(odd_gen)

```

### 字典推导式
```python
# 字典推导式
my_dict={"a":1,"b":2,"c":3}
reversed_dict={value:key for key,value in my_dict.items()}
```

### 集合推导式
```python
my_dict={"a":1,"b":2,"c":3}
my_set={key for key,value in my_dict.items()}
```


# 深入python的set和dict
## dict 常用方法
* dict.copy 浅拷贝，copy.deepcopy 深拷贝
* dict.fromkeys 将可迭代的对象转换成dict
* dict.get 可以避免KeyError错误
* dict.item 返回，key,value 用于迭代中
* dict.setdefault 获取值，如果不存在则会设置
* dict.update 接受可迭代对象

## dict 子类
* collections.UserDict
* collections.defaultdict

## set和frozenset
set 集合，frozenset 不可变集合    
不重复，无序  
### set
* set.update 合并两个
* set.difference 求差集
* ｜ & -  集合运算
* set 用C语言实现，性能很高，hash实现，查找元素时间复杂度为O(1)
* set.issubset 判定是否为另外一个set的子集
### frozenset 
其不可变性，可以作为dict的key  


## dict和set的实现原理


# 对象引用，可变性和垃圾回收
## python中的变量是什么？
python的变量实质上是一个指针 int str,便利贴

## ==和is 的区别
* == 是用于判定值是否等
* is 判定是否为同一个对象

## del和垃圾回收
cpython中的垃圾回收算法是 引用计数

## 一个经典的错误
* list,dict 为引用类型，作为参数传递时，函数内部对变量的操作，会影响外部
* 当list作为 class的参数时，所有class对于list的默认值会使用同一个list，位于 $class.__init__.__defaults__ 下





# 元类编程
## property 动态属性
* @property 标明动态属性
* @age.setter 设置动态属性

## __getattr__,__getattribute__ 魔法函数
*  __getattr__ 查找不到属性的时候的钩子函数
* __getattribute__  属性存不存在，都会执行这个钩子函数，很少使用。

## 属性描述符和属性的查找过程
* 属性描述符可以限制属性的类型，但感觉在py中没有必要，使用属性描述符的场景不如换强类型语言
* 属性描述符在框架的ORM，Model的位置用的比较多。

## __new__,__init__的区别
* __new__ 是用来控制对象的生成过程
* __init__ 是用来完善对象的，是对象生成之后

子类对象__new__ 中需要 return super().__new__(cls) 否则不会调用__init__   
__init__ 中的参数需要和初始化对象时的参数一致，否则会报错。


## 自定义元类
元类是可以动态创建类的类  
type-> class(对象) -> 对象    

MetaClass元类，本质也是一个类，但和普通类的用法不同，它可以对类内部的定义（包括类属性和类方法）进行动态的修改。可以这么说，使用元类的主要目的就是为了实现在创建类时，能够动态地改变类中定义的属性或者方法。
如果想把一个类设计成 MetaClass 元类，其必须符合以下条件：  
1. 必须显式继承自 type 类；
2. 类中需要定义并实现 __new__() 方法，该方法一定要返回该类的一个实例对象，因为在使用元类创建类时，该 __new__() 方法会自动被执行，用来修改新建的类。

```python
#定义一个元类
class FirstMetaClass(type):
    # cls代表动态修改的类
    # name代表动态修改的类名
    # bases代表被动态修改的类的所有父类
    # attr代表被动态修改的类的所有属性、方法组成的字典
    def __new__(cls, name, bases, attrs):
        # 动态为该类添加一个name属性
        attrs['name'] = "C语言中文网"
        attrs['say'] = lambda self: print("调用 say() 实例方法")
        return super().__new__(cls,name,bases,attrs)
```

元类相比类继承的方式，侧重于类的创建过程，类的创建有如下方式：  
1. type函数
2. class 关键字
3. meteclass 元类

# python中的迭代协议
## python中的迭代器和迭代协议
迭代是一种重复访问集合元素的方式，而迭代器和迭代协议是实现这一功能的核心机制。   
迭代器是一个可以记住遍历的位置的对象，而迭代协议则是一种约定，规定了迭代器应该如何工作。  


## 迭代器（Iterator）
迭代器是一个实现了迭代器协议的对象，它必须包含__iter__()和__next__()这两个方法。其中，__iter__()方法返回迭代器对象本身，如果类定义了__iter__()，那么它的实例对象就是一个迭代器；__next__()方法返回容器的下一个值，如果容器中没有更多元素了，那么抛出StopIteration异常。

## 可迭代对象（Iterable）
实现了__iter__方法，能够循环取值的对象即为可迭代对象，与迭代器的区别是，迭代器会存储当前迭代的位置，可迭代对象不会，例如 list   
当一个对象很大的时候,所占用的内存需要1G,那要对这个对象进行迭代，不能吧1G的数据直接载入内存中，而是通过迭代器的方式去访问数据  

## iter 函数
iter函数，会尝试将一个对象转化为一个迭代器。如果对象实现了__iter__函数，会有限调用这个，如果不存在，就是尝试调用__getitem__函数，来完成迭代器的创建


## 生成器
生成器函数，函数里只要有yield关键字  
具有yield关键字的函数，返回的是genreate对象   

## 生成器的原理
python的函数栈桢对象是分配在堆内存中的，即不会自动释放内存
每次遇到yield 函数就会停止执行，下一次执行的时候，在停止的地方继续执行  
生成器可以被next函数调用  

## 生成器在UserList中的应用

## 生成器如何读取大文件
```python
def read_big_file(f,splitStr):
    buf=""
    while True:
        while splitStr in buf:
            pos= buf.index(splitStr)
            yield buf[:pos]
            buf=buf[pos+len(splitStr):]
        chunk=f.read(2)

        if not chunk:
            #已经读到了最后
            yield buf
            break
        buf+=chunk



if __name__=="__main__":
    with open('bigfile.txt') as f:
        for chunk in read_big_file(f,","):
            print(chunk)
```


# python socket编程
## HTTP,Socket,TCP 
* HTTP 应用层协议基于TCP
* TCP 传输层协议
* Socket 建立连接的工具

## socket编程中 client与server实现通信
![](python3高级核心技术97讲/img_1.png)


## socket实现聊天和多用户连接

## socket发送http请求


# 多线程，多进程和线程池编程
## python GIL锁 
gil global interpreter lock (cpython)  
python中一个线程对应c语言中的一个线程   
gil 使得同一个时刻只有一个线程在一个cpu上执行字节码，无法将多线程映射到多个cpu上，无法利用多核

### 什么时候 GIL锁会释放
* 根据执行的字节码行数以及时间片到了
* 遇到有io操作

## 多线程编程-threading
多线程编程的几种方式
### import threading
对于io操作来说，多线程和多进程性能差别不大    
* thread.setDaemon 守护线程，设置守护线程后，主线程不会等待线程，直接会退出
* thread.join 等待线程执行完成后，才会继续往下执行

### 通过集成Thread来实现多线程

## 线程间通信
### 共享变量
共享变量通信的方式，是非线程安全的，需要加锁
### Queue
线程间通信推荐方式，类比 golang中的 channel     
golang 中推行的Communicating Sequential Processes（CSP）模型，并发编程的推荐通信方式（1978年推出）  
CSP模型的核心是“不要通过共享内存来通信，而是通过通信来共享内存”  


## 线程同步
### Lock 
互斥锁，golang中为Mutex  
锁带来的问题：  
1. 用锁会影响性能
2. 可能有死锁问题，死锁问题在golang编译过程中会报错
```
# 相互等待，导致的死锁
A acquire(a), acquire(b)
B acquire(b), acquire(a)

# 重复调用导致的锁
A acquire(a), acquire(a)

# A调用B
A acquire(a)
B acquire(a)
```
### RLock
为解决 A调用B,或重复调用导致的锁问题 ,设计出可重入的锁  RLock   
在相同的线程中，可以连续多次调用acquire, acquire的次数要和release的次数相同 （内部应该有一个计数器，以及线程ID）  
golang中没有可重入的锁，原因应该是和GMP模型有关，golang中的协程与线程是M:N的关系，

### Condition 
条件变量，复杂的线程同步  
在调用with cond之后才能调用wait或notify 方法  
condition有两层锁，一把底层RLock锁，在线程调用了wait方法的时候释放   
每次调用wait时会分配一把锁放入到cond的等待队列中，等待notify方法的唤醒  

### Semaphore
用于控制并发数量  
基于Condition


## ThreadPoolExector 线程池
concurrent  
频繁的创建以及销毁线程，会消耗资源，所以需要线程池    
主线程中可以获取某一个线程的状态或者某一个任务的状态，以及返回值  

### Futuer
未来对象，task的容器   
Futuer里面包含了workItem,workItem是线程执行的单元，执行结果会保存到futuer容器中  
ThreadPoolExector 会根据实例化时的线程数量，拉起线程，执行workItem  


## 多进程和多线程对比
python多线程编程由于GIL锁的存在，无法使用多核cpu，所以对于耗cpu的程序来说，使用python，需要使用多进程，才能提高程序性能  

concurrent.futures 提供了多进程与多线程统一抽象


### multiprocessing 多进程编程

### 进程间的通信-Queue,Pipe,Manager
1. 多进程的Queue和线程的Queue不是同一个
2. 共享全局变量不适用多进程
3. multiprocessing中的Queue,不能用于进程池
4. multiprocessing.pool 中的进程间通信需要使用manager中的Queue,manager需要实例化
5. python中有3个Queue,threading.Queue,multiprocessing.Queue,multiprocessing.manager.Queue
6. pipe只能适用于两个进程间通信,pipe性能高于queue

### 进程间共享变量
通过multiprocessing.manager.dict实现  


# 协程和异步IO
## 并发，并行，同步，异步，阻塞，非阻塞
* 并发是指一个时间段内，有几个程序在同一个cpu上运行，但是任意时刻只有一个程序在cpu上运行  
* 并行是指任意时刻点上，有多个程序同时运行在多个cpu上
* 同步是指代码在调用IO操作时，必须等待IO操作完成才返回的调用方式
* 异步是指代码调用IO操作时，不必等待IO操作完成就返回的调用方式
* 阻塞是指调用函数的时候当前线程会被刮起,阻塞不会消耗cpu
* 非阻塞是指调用函数的时候当前线程不会被挂起，而是立即返回

## C10K问题和IO多路复用（select,poll,epoll）

### C10K问题
如何在一个1GHz CPU,2G内存，1gbps网络环境下，让单台服务器同时为1万个客户端提供FTP服务   

### Unix下的五种I/O模型  
#### 阻塞式I/O 主流
![](python3高级核心技术97讲/img_2.png)
#### 非阻塞式I/O 
![](python3高级核心技术97讲/img_3.png)

#### I/O复用 
为解决非阻塞式I/O中，不停等待IO返回，耗费cpu的情况，出现了IO多路复用，当内核的数据准备好后，内核会通知应用程序   
select 是一个阻塞函数，其有两个优点   
1. 阻塞期间是不会消耗cpu
2. 可以同时监听多个socket
![](python3高级核心技术97讲/img_4.png)

#### 信号驱动I/O
很少应用  

#### 异步I/O (POSIX的aio系列函数)
![](python3高级核心技术97讲/img_5.png)

### select,poll,epoll
select,poll,epoll 都是IO多路复用的机制，I/O多路复用就是通过一种机制，一个进程可以监听多个描述符，一旦某个描述符就绪（读就绪，写就绪）,能够
通知程序进行响应的读写操作  
但select,poll,epoll 本质上都是同步IO,因为他们都需要在读写事件就绪后自己负责拷贝数据（内核空间拷贝到用户空间），也就是读写这个过程是阻塞的，而
异步IO则无需自己负责读写，异步IO的实现会负责把数据从内核拷贝到用户空间


#### select
select 函数监视的文件描述符分3类，分别是writefds,readfds,exceptfds，调用后select会阻塞，直到有文件描述符就绪（有数据可读，可写，except）,
或者超时（timeout指定等待时间，如果需要立即返回设置为null即可），函数返回。当select函数返回后，可以通过遍历fdset,来找到就绪的描述符  

select 缺点在于  
1. 单个进程能够监视的文件描述符数量有限制，为1024 （修改这个限制需要重新编译内核）
2. 每次都需要将1024全部遍历完成来获取就绪的fdset

#### poll
poll 相比于select 解决了1024限制的问题   
不同于select使用三个位图来表示三个fdset的方式，poll使用pollfd的指针实现。pollfd结构体其包含了要监视的event和发生的event，不再使用select"参数-值"传递的方式。   
同时，pollfd并没有最大数量限制（但是数量过大后，性能也是会下降）。和select函数一样，poll返回后，需要轮询pollfd来获取就绪的描述符。   
从上面看，select和poll都需要在返回后，通过遍历文件描述符来获取已经就绪的socket。事实上，同一个连接的大量客户端在一个时刻可能只有极少数的socket处于
就绪状态，因此随着监视的描述符数量的增长其效率也会线性下降


#### epoll
epoll 是在linux2.6中提出来的，无描述符限制。epoll使用一个文件描述符管理多个描述符，将用户关系的文件描述符的事件存放到内核的一个事件表中，这样在
用户空间和内核空间的copy只需一次

## 回调之痛
1. 可读性差
2. 共享状态管理差
3. 异常处理困难

## 协程是什么
C10M的问题   
如何利用8核心CPU,64G内存，在10gbps的网络上保持1000万并发连接   

非协程情况下的问题：  
1. 回调模式编码复杂度高 
2. 同步编程的并发性不高
3. 多线程编程需要线程间同步，线程间同步依赖锁

想到达到的效果：  
1. 采用同步的方式去编写异步代码
2. 使用单线程去切换任务
    *  线程的切换是由操作系统切换的，单线程切换意味着我们需要程序员自己调度任务
    *  单线程不需要锁，并发性高，如果单线程切切换函数，性能远高于线程切换，并发性高

协程是可以暂停的函数  


协程是解决cpu与IO运行速度不匹配的问题，当一个线程中存在阻塞时调用的时候，cpu可以去执行那些纯计算的任务，提高cpu的利用率（但是对于python这样的语言来说，只能用1个cpu,提高比较有限） 
一句话就是cpu不能停   
python的协程是运行在单线程上的，也就是说线程:协程为 1:N, 相比于golang的 M:N的调度模型弱了很多，且golang中封装出的协程语法更简单   




## 生成器进阶-send,close和throw方法
这三个函数都是生成器的方法
### send
启动生成成器的方式有两种：next,send  
send方法不仅可以向函数发送数据，同时还会重启生成器   
在send发送非None值之前，必须先启动生成器，两种方式：  
1. gen.send(None)
2. next(gen)

### close
关闭生成器

### throw
向生成器传递异常


## yield from
python3.3 加入的语法  
1. 子生成器生产的值，都是直接传递给调用方的；调用方通过.send()发送的值都是直接传递给子生成器的，如果传递的是None调用子生成器的next方法，否则调用子生成器的send方法 
2. 子生成器退出的时候，最后 return EXPR ,会触发一个StopIteration(EXPR)异常
3. yield from 表达式的值，是子生成器终止时，传递给StopIteration异常的第一个参数
4. 如果调用的时候出现了StopIteration异常，委托生成器会恢复运行，同时其他的异常会向上“冒泡”
5. 传入委托生成器的异常里，除了GeneratorExit外，其他的所有异常全部传递给子生成器的.throw方法，如果调用.throw出现了StopIteration,那么就恢复委托生成器的运行，其他的异常全部向上“冒泡”
6. 如果在委托生成器上调用.close()或者传入GeneratorExit异常，会调用子生成器的.close方法，没有.close方法则不会调用，


## async和await
通过生成器 yield from（python3.3） 可以实现协程，到python3.5后 python提供async和await两个关键词用于支持原生的协程  
