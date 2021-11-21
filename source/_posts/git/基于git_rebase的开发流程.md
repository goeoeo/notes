---
categories:
- git

tags:
- git-rebase
---



## 分支
* master  #不能直接推送
* staging  #开发主流程分支
* self_branch #自己的开发分支，从staging 打出


<!--more-->



## 开发与代码合并
1. git add .  #在self_branch完成代码开发，使用goimports 格式化代码
2. git commit -m "备注" # 提交到本地仓库
3. git checkout staging && git pull && git checkout self_branch && git rebase staging # 解决冲突
4. git commit -m "冲突备注" # 有冲突执行这一步
5. git push -f origin self_branch #将自己最新代码推送置远程仓库，
6. 将 self_branch 合并到 staging 生成一个mr(merge request)
7. 将mr 连接发送给同组人员进行code review, 同组人员合并代码


## 发布
1. 将 staging 代码合并到 master
2. 基于 master 打出版本tag  v0.1.1-release

## 线上hotfix

### 打出hotfix 分支
1. git fetch origin --prune # 拉取远程仓库信息
2. git checkout -b hotfix/bugname v0.1.1-release #基于最新发布，打出hotfix 分支

### 进行hotfix
3. 进行bugfix
4. git commit -m "bug原因描述"
5. git log # 记录commit id :545a1a4ba417d
6. git checkout self_branch # 切换到自己的分支
7. git cherry-pick 545a1a4ba417d #将bugfix代码合并到自己的分支
8. 对staging 进行rebase
9. git push -f origin self_branch #将自己最新代码推送置远程仓库，
10. 在staging 验证bug是否修复

### 发布
11. 基于hotfix/bugname 分支 打出tag v0.1.2-release
12. 发布 线上验证
13. git push origin --delete hotfix/bugname #删除远程分支
14. git branch -D hotfix/bugname #删除本地分支

