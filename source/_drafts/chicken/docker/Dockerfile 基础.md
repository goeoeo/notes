### 底层技术支持
* Namespaces #做隔离pid,net,ipc,mnt,uts
* Control groups #做资源限制
* Union file systems  #Container和image的分层

###  语法
1.FROM #声明一个基本镜像
* FROM scratch #制作base image
* FROM ubuntu:18.04 #使用base image
>尽量使用官方的image作为base image

2.LABEL #标注
* LABEL version="1.0"
>Metadata不可少

3.RUN #执行命令
* 每运行一次就会生成一层container 
>* 为了美观，复杂的RUN 请用反斜线换行
>* 避免无用分层，合并多条命令成一行

4.WORKDIR #设定工作目录，如果没有会自动创建
>* 用WORKDIR,不要使用RUN cd 
>* 尽量使用绝对目录

5.ADD ,COPY #添加和保存文件到容器中
* 区别,ADD 添加压缩文件到容器会自动解压缩
>* 大部分情况，COPY优于ADD 
>* ADD除了COPY 还有额外的解压功能
>* 添加远程文件、目录，使用curl或者wget

6.ENV #设置常量
* ENV MYSQL_VERSION 5.6


7.CMD #设置容器启动后默认执行的命令和参数
* 容器启动时默认执行的命令
* 如果docker run 指定了其他命令，CMD 命令会被忽略
* 如果定义了多个CMD ,只有最后一个会执行


8.ENTRYPOINT #设置容器启动时运行的命令
* 让容器以应用程序或者服务的形式运行
* 不会被忽略，一定会执行
* 最佳实践: 写一个shell脚本作为entrypoint


### 例子
1.命令参数的传入
```bash
FROM ubuntu
RUN apt-get update && apt-get install -y stress
ENTRYPOINT ["/usr/bin/stress"]
CMD []

```
2.容器资源限制
1. --memary #限制容器内存
2. --cpu-shares #cpu 占用权重
3. --cpu #指定占用cpu的个数





