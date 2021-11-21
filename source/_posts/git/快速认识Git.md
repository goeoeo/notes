---
categories:
- git 

tags:
- git
---

## 概述

分布式版本控制系统，由Linux创始人Linus使用C语言实现
<!--more-->

## 基本使用

### 创建版本库

初始化一个git 仓库，命令执行后，会在当前目录下生成一个隐藏目录.git

```
git init
```

### git本地架构

![](快速认识Git/0.jpeg)

#### git组成

* 工作区（所有的操作，在git add 之前都位于工作区域）
* 版本库（整个.git目录）
    * 暂存区
    * 若干分支

#### 文件到版本库的过程

1. 文件位于工作区
2. 使用git add 命令将文件从工作区提交到暂存区
3. 使用git commit 命令将文件从暂存区提交到某个分支（默认提交到当前分支）

### 撤销修改、版本回退

* 文件位于工作区，丢弃工作区修改，直接从版本库中检出文件覆盖

```
git checkout -- file
```

* 文件位于暂存区，暂存区的修改撤销掉，重新放回工作区

```
git reset HEAD <file>
```

* 文件已经提交到分支
    * 回滚到上一个版本
        ```
        git reset --hard HEAD^ #HEAD代表当前版本， HEAD^代表上一个版本
        ```
    * 回滚到指定版本（向后回滚）
        ```
        git log #查看版本id
        git reset --hard id #执行回滚
        ```
    * 回滚到指定版本（向前回滚）
      ```
      git reflog #记录了你的每一次操作，通过此命令查看版本id
      git reset --hard id #执行回滚
      ```

### 远程仓库

远程仓库存在的意义在于同步代码

#### 本地已有仓库，远程仓库为空

1.将本地仓库关联到远程仓库

```
git remote add origin git@server-name:path/repo-name.git
```

> 添加后，远程库的名字就是origin，这是Git默认的叫法，也可以改成别的，但是origin这个名字一看就知道是远程库。

2.第一次提交

```
git push -u origin master
```

> 由于远程库是空的，我们第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来，在以后的推送或者拉取时就可以简化命令

#### 本地已有仓库，远程仓库非空

1.将本地分支和远程分支关联

```
git branch --set-upstream-to=origin/master master
```

2.使用git pull整合远程仓库和本地仓库，使用这个命令后你需要解决本地与远程的冲突，非常麻烦，谨慎使用

```bash
git pull --allow-unrelated-histories 忽略版本不同造成的影响
```

#### 本地无仓库，直接从远程仓库克隆

```
git clone git@server-name:path/repo-name.git
```

#### git配置过程中fatal:拒绝合并无关的历史

```
首先将远程仓库和本地仓库关联起来：

git branch --set-upstream-to=origin/master master

然后使用git pull整合远程仓库和本地仓库，

git pull --allow-unrelated-histories    (忽略版本不同造成的影响)

```

### 分支合并策略

通常，合并分支时，如果可能，Git会用Fast forward模式，但这种模式下，删除分支后，会丢掉分支信息。

如果要强制禁用Fast forward模式，Git就会在merge时生成一个新的commit，这样，从分支历史上就可以看出分支信息。

使用--no-ff 禁用Fast forward模式

```
git merge --no-ff -m "merge with no-ff" dev
```

合并分支时，加上--no-ff参数就可以用普通模式合并，合并后的历史有分支，能看出来曾经做过合并，而fast forward合并就看不出来曾经做过合并。

### 标签

tag 是某个commit id的别名

#### 创建标签

* git tag v1.0 # git默认将标签打在当前commit id 上
* git tag v0.9 f52c633 #指定commit id 打标签
* git tag -a v0.1 -m "version 0.1 released" 1094adb #打出带有说明的标签
* git tag #查看标签

> 标签总是和某个commit挂钩。如果这个commit既出现在master分支，又出现在dev分支，那么在这两个分支上都可以看到这个标签。

#### 操作标签

* git tag -d v0.1 #删除标签
* git push origin v1.0 #推送标签
* git push origin --tags #一次性推送全部尚未推送到远程的本地标签
* 删除远程标签
    * git tag -d v0.9 #先删除本地标签
    * git push origin :refs/tags/v0.9 #通知远程删除标签

## 常用命令

### 设置

* git config --global credential.helper store #记住用户名和密码
* git config --global core.filemode false #忽略文件权限的修改

### 撤销修改

* git checkout -- file #撤销git add 操作
* git reset HEAD <file> #撤销commit 操作

### 版本回滚

* git reset --hard origin/master #回滚到远程分支，有冲突的时候，放弃本地修改，使用这个
* git reset --hard HEAD^ #回到上一个版本
* git reset --hard id #回滚到指定版本，配合git log 或者git reflog 使用

### 忽略文件

* .gitignore

> 需要注意的是已经提交到版本库中的文件忽略会无效，需要同时使用下面这个命令才会生效,其次.gitignore 文件中如果不指定目录，则会递归的去忽略文件或者目录

* git update-index --assume-unchanged config/pay.php #忽略已经提交到版本库的文件

### 远程仓库

* git remote add origin git@server-name:path/repo-name.git #添加远程仓库
* git remote set-url origin git@server-name:path/repo-name.git #修改远程仓库地址
* git push -u origin master #第一次推送，实际上这条命令是将当前的master分支推送到远程分支,git 会自动关联远程的分支和本地的分支

### 分支

* git checkout -b dev #创建并切换到dev分支
* git branch --set-upstream-to=origin/dev dev #将远程的dev分支与本地的dev分支进行关联
* git checkout -b dev origin/dev #创建并切换到dev分支并与远程dev分支关联(直接拉取远程分支)
* git checkout master #切换到master分支
* git merge --no-ff -m '合并dev分支' dev # 禁用快进方式合并分支
* git branch -d feature-vulcan #删除一个已经合并过的分支
* git branch -D feature-vulcan #删除一个没有合并国的分支
* git push origin --delete branchName # 删除远程分支
* git remote prune origin #删除所有远端已经删除本地仍然存在的分支,本地未推送到远端的分支不会被删除

### 标签

* git tag -a <tagname> -m "blablabla..." #为某一次提交打一个标签
* git push origin v1.01 #将v1.01标签推送到远程
* git push origin :refs/tags/v1.00 #删除远程v1.00标签

### 日志

* git log --graph --pretty=oneline --abbrev-commit #友好的显示日志信息，在linux下可以为这个长命令起一个别名

## 问题

### error: RPC 失败。HTTP 411 curl 22 The requested URL returned error: 411

把本地发送的body限制改到500m

```
git config http.postBuffer 524288000
```source/_posts/go/golang项目结构规范.md