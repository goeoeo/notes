## 开发环境：ubuntu18.04
#### 安装
1.下载docker并安装
```bash
sudo apt install docker.io #ubuntu18.04 下可以用这个，我用的这个

sudo wget -qO- https://get.docker.com/ | sh #下载docker并安装
```

2.将cy用户加入到docker组中，默认情况下只有root可以运行docker
```bash
#用下面这个两个命令
sudo gpasswd -a $USER docker     #将登陆用户加入到docker用户组中
newgrp docker     #更新用户组 ,这个命令是暂时的，重启电脑，每个终端都能生效

```

3.安装docker-compose
```bash
sudo curl -L --fail https://github.com/docker/compose/releases/download/1.23.1/run.sh -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version

```
遇到一个警告：

WARNING: Error loading config file: /home/cy/.docker/config.json: stat /home/cy/.docker/config.json: permission denied

解决警告：
```bash
sudo chown "$USER":"$USER" /home/"$USER"/.docker -R
sudo chmod g+rwx "/home/$USER/.docker" -R
```
#### 更换镜像
1.编写配置文件
```bash
sudo vim /etc/docker/daemon.json

{
  "registry-mirrors": ["https://9lrfffi7.mirror.aliyuncs.com"]
}

sudo systemctl daemon-reload
sudo systemctl restart docker
```

2.重启docker
```bash
sudo systemctl daemon-reload #重新加载配置文件
sudo service docker restart #重启docker
#这里建议使用sudo service docker stop 先将docker服务彻底关闭，因为我在使用的过程中发现，不彻底关闭，docker search 不能用
```

#### docker 命令
1.查看docker信息（version、info）
```bash
docker version            #查看docker版本
docker info               #显示docker系统的信息
```
2.对image的操作（search、pull、images、rmi、history）
```bash
docker search image_name  #检索image
docker pull image_name    #下载image
docker images             #列出镜像列表
docker rmi image_name     #删除一个或者多个镜像
docker history image_name #显示一个镜像的历史
```
3.容器（run）
```bash
docker run image_name echo "hello word"           #在容器中运行"echo"命令，输出"hello word"
docker run -i -t image_name /bin/bash             #交互式进入容器中
docker run image_name apt-get install -y app_name #在容器中安装新的程序
docker run -p 80:80 -d -v $PWD/html:/usr/share/nginx/html nginx-fun 共享文件
docker run 
```
4.查看容器（ps）
```bash
docker ps    #列出当前所有正在运行的container
docker ps -a #列出所有的container
docker ps -l #列出最近一次启动的container
```
5.保存对容器的修改（commit）
```bash
# 保存对容器的修改; -a, --author="" Author; -m, --message="" Commit message
docker commit ID new_image_name
```


6.删除命令 
```bash
docker rm $(docker container ls -f "status=exited" -q) #删除所有已经退出的容器
docker rmi $(docker images| grep '<none>' | awk '{print $3}') #删除有<none>的镜像
```
7.停止
```bash
 docker stop $(docker ps | sed -n "2, 1p" |awk '{print $1}') #停止第一个docker 容器

```


#### Dockerfile语法


| 命令 | 用途 | 技巧 |
| :------: | :------: | :------ |
| FROM | bash image | 
| LABEL | 描述标签 |
| RUN | 执行命令 | 避免无用分层,使用\,将多条命令合并成一行执行
| ADD | 添加文件 | ADD可以自动解压
| COPY | 拷贝文件 | 大部分情况,COPY优于ADD,添加远程文件/目录,请使用curl/wget
| CMD | 设置容器启动后默认执行的命令和参数 | 1.如果docker run 指定了其他命令,CMD会被忽略<br>2.如果定义了多个CMD,只有最后一个会执行
| ENTRYPOINT | 设置容器启动时运行的命令 | 1.让容器以应用程序或者服务的形式执行<br>2.不会被忽略,一定会执行<br>3.写一个shell脚本作为entrypoint
| EXPOSE | 暴露端口 |
| WORKDIR | 设定当前工作路径 | 用WORKDIR来改变路径,不要用 RUN cd,尽量使用绝对目录
| MAINTAINER | 维护者 |
| ENV | 设定环境变量 | 尽量使用ENV,提高可维护性
| USER | 指定用户 |
| VOLUME | mount point |




#### Registry相关
1.登录阿里云Docker Registry
```bash
sudo docker login --username=chenyu977564830 registry.cn-hangzhou.aliyuncs.com
```
2.从Registry中拉取镜像
```bash
sudo docker pull registry.cn-hangzhou.aliyuncs.com/phpdi/lnmp:[镜像版本号]
```

3.将镜像推送到Registry
```bash
sudo docker login --username=chenyu977564830 registry.cn-hangzhou.aliyuncs.com
sudo docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/phpdi/lnmp:[镜像版本号]
sudo docker push registry.cn-hangzhou.aliyuncs.com/phpdi/lnmp:[镜像版本号]
```
#### 安装elasticsearch
1.第一步拉取elasticsearch6.5.0
```bash
docker pull elasticsearch:6.5.0
```
2.安装ik插件
* 原本我是想用Dockerfile 来构建这个插件的,但是/home/cy/S.md/docker/docker.md 要询问权限,所以只能进入容器里面去安装好,自己制作这个镜像.
1.由于elasticsearch指定了entryponit 所以使用docker run -it 命令进不去,需要加--entrypoint 参数
```bash
docker run -it --entrypoint /bin/bash elasticsearch:6.5.0
```
2.进入容器后,手动安装ik插件
```bash
elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v6.5.0/elasticsearch-analysis-ik-6.5.0.zip
```
此时提示错误:max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144],应该是内存给小了,
我本来想直接修改容器的vm.max_map_count,可是提示:setting key "vm.max_map_count": Read-only file system,docker系统是只读的.
3.退出容器,修改系统参数x
```bash
sudo vim /etc/sysctl.conf
#在最后一行添加
vm.max_map_count=262144
```
4.保存镜像,注意在commit 的时候需要指定entrypoint ,否则直接启动新镜像,不会执行原镜像的entryponit
```bash
docker commit -m 'ik,entrypoint' --change='ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]' 58c793849522 elasticsearch-ik1:6.5.0

#/usr/local/bin/docker-entrypoint.sh 这个是原镜像的entryponit ,是通过下面这个命令查到的
docker history elasticsearch:6.5.0 --no-trunc=true

```
5.启动镜像
```bash
docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:ik1
```
6.进入启动后的容器
注意：这里不能使用docker attach， 这个命进入的是enpryponit的位置 所以通过这个命令进入后是elasticsearch的进程交互界面
```bash
docker exec -it 21088ed3ba2f /bin/bash 
```


#### docker-compose

1.docker-compose.yml 常用命令

| 命令 | 用途 |
| :------: | :------:|
| build | 本地镜像创建 值为目录 | 
| command | 覆盖缺省命令 |
| depends_on | 连接容器 |
| ports | 暴露端口 |
| volumes | 挂载卷 |
| image | pull镜像 |

2.docker-compose 命令

| 命令 | 用途 |
| :------: | :------:|
| up | 启动服务 | 
| stop | 停止服务 |
| rm | 删除服务中的各个容器,可加-f参数,全部删除 |
| logs | 观察各个容器的日志 |
| ps | 列出服务相关的容器 |


### 容器时间
**问题描述:**docker 容器默认使用的时区为UTC ,与宿主机CST时区相差8个小时
**解决方案:**挂载宿主机时间到容器
-v /etc/timezone:/etc/timezone -v /etc/localtime:/etc/localtime 

### 容器日志
docker logs -f -t --tail 行数 容器名