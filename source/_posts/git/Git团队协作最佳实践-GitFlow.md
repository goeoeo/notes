---
title: Git团队协作最佳实践-GitFlow
categories: 
- git

tags:
- git-flow
---
## Git的优点
* 分布式，不需要远程仓库，也可以正常工作
* 优秀的分支模型，创建和删除方便且高效，因为分支的创建和删除，只是操作HEAD指针。
* 分支合并轻松
<!--more-->

## 版本管理的挑战
虽然有这么优秀的版本管理工具，但是我们面对版本管理的时候，依然有非常大得挑战，我们都知道大家工作在同一个仓库上，那么彼此的代码协作必然带来很多问题和挑战，如下：
* 如何开始一个Feature的开发，而不影响别的Feature？
* 由于很容易创建新分支，分支多了如何管理，时间久了，如何知道每个分支是干什么的？
* 哪些分支已经合并回了主干？
* 如何进行Release的管理？开始一个Release的时候如何冻结Feature, 如何在Prepare Release的时候，开发人员可以继续开发新的功能？
* 线上代码出Bug了，如何快速修复？而且修复的代码要包含到开发人员的分支以及下一个Release?

大部分开发人员现在使用Git就只是用三个甚至两个分支，一个是Master, 一个是Develop, 还有一个是基于Develop打得各种分支。这个在小项目规模的时候还勉强可以支撑，因为很多人做项目就只有一个Release, 但是人员一多，而且项目周期一长就会出现各种问题。

## Git Flow
就像代码需要代码规范一样，代码管理同样需要一个清晰的流程和规范

Vincent Driessen 同学为了解决这个问题提出了 [A Successful Git Branching Model](https://nvie.com/posts/a-successful-git-branching-model/)

下面是Git Flow的流程图
![](Git团队协作最佳实践-GitFlow/git-model@2x.png)

### Git Flow常用分支
* Production分支
>  也就是我们经常使用的Master分支，这个分支最近发布到生产环境的代码，最近发布的Release， 这个分支只能从其他分支合并，不能在这个分支直接修改
* Develop 分支 
> 这个分支是我们是我们的主开发分支，包含所有要发布到下一个Release的代码，这个主要合并与其他分支，比如Feature分支
* Feature 分支 
> 这个分支主要是用来开发一个新的功能，一旦开发完成，我们合并回Develop分支进入下一个Release
* Release分支 
> 当你需要一个发布一个新Release的时候，我们基于Develop分支创建一个Release分支，完成Release后，我们合并到Master和Develop分支
* Hotfix分支 
> 当我们在Production发现新的Bug时候，我们需要创建一个Hotfix, 完成Hotfix后，我们合并回Master和Develop分支，所以Hotfix的改动会进入下一个Release

### Git Flow工作原理
#### 初始分支
所有在Master分支上的Commit应该Tag
![](Git团队协作最佳实践-GitFlow/o_git-workflow-release-cycle-1historical.png)

#### Feature 分支
分支名 feature/*

Feature分支做完后，必须合并回Develop分支, 合并完分支后一般会删点这个Feature分支，但是我们也可以保留
![](Git团队协作最佳实践-GitFlow/o_git-workflow-release-cycle-2feature.png)

#### Release分支
分支名 release/*

Release分支基于Develop分支创建，打完Release分之后，我们可以在这个Release分支上测试，修改Bug等。同时，其它开发人员可以基于开发新的Feature 
(记住：一旦打了Release分支之后不要从Develop分支上合并新的改动到Release分支)

发布Release分支时，合并Release到Master和Develop， 同时在Master分支上打个Tag记住Release版本号，然后可以删除Release分支了。
![](Git团队协作最佳实践-GitFlow/o_git-workflow-release-cycle-3release.png)

#### 维护分支 Hotfix
分支名 hotfix/*

hotfix分支基于Master分支创建，开发完后需要合并回Master和Develop分支，同时在Master上打一个tag
![](Git团队协作最佳实践-GitFlow/o_git-workflow-release-cycle-4maintenance.png)


### Git Flow代码示例

1. 创建develop分支
```
git branch  develop
git push -u origin develop   
```

2. 开始新Feature开发
```
git checkout -b some-feature develop

# 推送功能分支到远程，也可以不推送
git push -u origin some-feature    

# 做一些改动    
git status
git add some-file
git commit 

```

3. 完成Feature
```
git pull origin develop
git checkout develop
git merge --no-ff some-feature
git push origin develop

git branch -d some-feature

# 如果你推送了功能分支到远程，执行这个命令删除远程分支
git push origin --delete some-feature 
```

4. 开始Relase
```
git checkout -b release-0.1.0 develop

# 如果有bug 在此分支上处理
```

5. 完成Release
```
# 合并release到master
git checkout master
git merge --no-ff release-0.1.0
git push

# 合并release到develop 
git checkout develop
git merge --no-ff release-0.1.0
git push

# 删除release分支
git branch -d release-0.1.0

# 如果分支推送到了远程，删除远程分支:
git push origin --delete release-0.1.0 

# 在master分支上打出版本号
git tag -a v0.1.0 master
git push --tags
```

6. 开始Hotfix
```
git checkout -b hotfix-0.1.1 master
```

7. 完成Hotfix
```
git checkout master
git merge --no-ff hotfix-0.1.1
git push

git checkout develop
git merge --no-ff hotfix-0.1.1
git push

git branch -d hotfix-0.1.1

git tag -a v0.1.1 master
git push --tags
```

### Git flow工具
实际上，当你理解了上面的流程后，你完全不用使用工具，但是实际上我们大部分人很多命令就是记不住呀，流程就是记不住呀，肿么办呢？

总有聪明的人创造好的工具给大家用, 那就是Git flow script.

#### 安装
* OS X
```
brew install git-flow
```
* Linux
```
apt-get install git-flow
```
* Windows
```
wget -q -O - --no-check-certificate https://github.com/nvie/gitflow/raw/develop/contrib/gitflow-installer.sh | bash
```

#### 使用

* 初始化: git flow init
* 开始新Feature: git flow feature start MYFEATURE
* Publish一个Feature(也就是push到远程): git flow feature publish MYFEATURE
* 获取Publish的Feature: git flow feature pull origin MYFEATURE
* 完成一个Feature: git flow feature finish MYFEATURE
* 开始一个Release: git flow release start RELEASE [BASE]
* Publish一个Release: git flow release publish RELEASE
* 完成Release: git flow release finish RELEASE
* 推送tag: git push --tags
* 开始一个Hotfix: git flow hotfix start VERSION [BASENAME]
* 发布一个Hotfix: git flow hotfix finish VERSION
![](Git团队协作最佳实践-GitFlow/o_git-flow-commands.png)

#### GUI工具-SourceTree
目前SourceTree支持Mac, Windows, Linux.
使用 自行探索

## 合并指定提交
```
先用git log查看，C2 commit的id，复制下来
git checkout 到master分支下
git cherry-pick <C2_id>

```

## 参考
* [Git 在团队中的最佳实践--如何正确使用Git Flow](https://www.cnblogs.com/wish123/p/9785101.html)

* 声明：本站部分图片或文章内容来源于网络，版权归原作者所有，如有侵权，请与我联系删除。
