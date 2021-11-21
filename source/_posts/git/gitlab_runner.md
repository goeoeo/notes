---
categories:
- git

tags:
- gitlab_runner
---


# Gitlab Runner

### GitLab-CI

GitLab-CI就是一套配合GitLab使用的持续集成系统（当然，还有其它的持续集成系统，同样可以配合GitLab使用，比如Jenkins）。而且GitLab8.0以后的版本是默认集成了GitLab-CI并且默认启用的。.gitlab-ci.yml的脚本解析就由它来负责。

<!--more-->


### GitLab-Runner

GitLab-Runner是配合GitLab-CI进行使用的。一般地，GitLab里面的每一个工程都会定义一个属于这个工程的软件集成脚本，用来自动化地完成一些软件集成工作。当这个工程的仓库代码发生变动时，比如有人push了代码，GitLab就会将这个变动通知GitLab-CI。这时GitLab-CI会找出与这个工程相关联的Runner，并通知这些Runner把代码更新到本地（服务器）并执行预定义好的执行脚本。



###  .gitlab-ci.yml

当有新内容push到仓库后，GitLab会查找是否有.gitlab-ci.yml文件，如果文件存在， Runners 将会根据该文件的内容开始build 本次commit。

.gitlab-ci.yml 使用YAML语法， 你需要格外注意缩进格式，要用空格来缩进，不能用tabs来缩进。

基本的`.gitlab-ci.yml`结构如下：

```yaml
stages:
  - build
  - test 
  - deploy
before_script:
  - echo "restoring packages..."
build_jog:
  stage: build
  script:
  - echo "Release build..."
  except:
  - tags
test_job:
  stage: test
  script:
  - echo "Tests run..."
  - cd cds.ci.test
  - gradle test
  tags:
    - wikirunner 
```

> 最上面的stages配置意思是，先构建阶段为build的job，然后再构建阶段为test的job，下面build_job和test_job都是job，如果不配置stages，默认顺序为build - test - deploy。
>
> 通过tags wikirunner 指定使用含有tag wikirunner的runner



### gitlab runner 安装

> gitlab 和 gitlab runner 不需要安装到同一台机器上 ，gitlab runner 通过register 将自身注册到gitlab中，当gitlab项目有推送事件，gitlab-ci会根据.gitlab-ci.yml,查找gitlab runner执行

官网安装地址:https://docs.gitlab.com/runner/install/linux-repository.html

ubuntu 安装

```shell
curl -L "https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh" | sudo bash
sudo apt-get install gitlab-runner

# 启动
sudo gitlab-runner run
```

### 在gitlab代码库中注册gitlab-runner

进入仓库->settings->CI/CD Pipelines  ![image-20211119145209526](/home/yu/code/notes/source/_posts/git/gitlab_runner/image-20211119145209526.png)

2. 执行注册

   ```shell
   root@i-2ns2mtbw:~# gitlab-runner register
   Runtime platform                                    arch=amd64 os=linux pid=1052553 revision=e95f89a0 version=13.4.1
   Running in system-mode.                            
                                                      
   Please enter the gitlab-ci coordinator URL (e.g. https://gitlab.com/):
   https://git.internal.com/
   Please enter the gitlab-ci token for this runner:
   u_-sf4e44oq9M5_QTn2x
   Please enter the gitlab-ci description for this runner:
   [i-2ns2mtbw]: benchmark
   Please enter the gitlab-ci tags for this runner (comma separated):
   benchmark
   Registering runner... succeeded                     runner=u_-sf4e4
   Please enter the executor: docker, docker-ssh, parallels, docker-ssh+machine, docker+machine, kubernetes, custom, shell, ssh, virtualbox:
   shell
   Runner registered successfully. Feel free to start it, but if it's running already the config should be automatically reloaded! 
   
   ```

3.  查看状态

```shell
gitlab-runner list
```



### 依赖 OpeNVPN

1. 单机测试机器上安装openvpn，参考 [Ubuntu下OpenVPN客户端配置教程](https://cwiki.yunify.com/pages/viewpage.action?pageId=52266586)
2. Gitlab Runner在执行前，需要从代码仓库中拉取最新的代码，因此需要有权限执行 git fetch/git pull 命令，由于我们的gitlab仓库需要VPN访问，所以在Runner所在的机器需要配置VPN Client，参考命令：openvpn --cd /etc/openvpn/client --config office.ovpn
3. 由于每次发布构建镜像，都会在本地保存一份，因此Runner所在的机器需要足够的硬盘空间，并且推荐定期清理无用的docker container 和 image
