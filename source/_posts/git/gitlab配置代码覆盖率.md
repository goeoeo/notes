---
title: gitlab配置代码覆盖率
categories:
- git

tags:
- 代码覆盖率
---

gitlab 中配置golang单元测试已经代码覆盖率

<!--more-->


## .gitlab-ci.yml
```yaml
variables:
  CI_REGISTRY_IMAGE: '$CI_REGISTRY/$CI_REGISTRY_NAMESPACE/$CI_PROJECT_NAME'
  GOPROXY: https://goproxy.cn
  TAG: "$CI_COMMIT_TAG"

stages:
  - fmt
  - unit-test

fmt:
  stage: fmt
  tags:
    - pd
  script:
    - make fmt-check



unit-test:
  stage: unit-test
  tags:
    - pd
  script:
    - make unit-test
```

## gitlab页面配置  
Project>Settings>CI/CD>General pipelines

### Test coverage parsing: 
1.单包测试覆盖率统计:  
```
coverage: \d+.\d+% of statements  
``` 
2.多包总体测试覆盖率:  
```
\d+.\d+%
```


Pipeline status 中有显示徽标的代码配置到md文件中即可  


## golang 代码覆盖率

### 参数解释
* -coverkpg参数会自动开启 -cover, 可省略-cover
* -coverprofile 将结果保存，后面html显示可以用到


### 获取单个包pkg的测试覆盖率
```
go test  -v  -count 1 -cover -coverpkg ./pkg/... ./pkg
```
### 获取单个包pkg的测试覆盖率并输出到网页
```yaml
go test  -v  -count 1 -cover -coverpkg ./pkg/... ./pkg  -coverprofile=coverprofile.cov && go tool cover -html=coverprofile.cov
```

### 获取多个包pkg及其子包的测试覆盖率
```
go test  -v  -count 1 -cover -coverpkg ./pkg/... ./pkg/...
```
### 获取多个包pkg及其子包的测试覆盖率 并统计总体覆盖率
对应同时测试多个包的时候，输出的是每个包的测试覆盖率 并没有合并 ，需要借助go tool 工具来获得整个项目的测试覆盖率
```
go  test  -count 1  -coverpkg ./pkg/...  ./pkg/...  -coverprofile=coverprofile.cov && go tool cover -func=coverprofile.cov
```
> 提取测试覆盖率 coverage=$(go tool cover -func=cover.out | tail -1  | grep -P '\d+\.\d+(?=\%)' -o)

可以实现一个脚本用来判断测试覆盖率是否小于master ,小于则报错
参考：https://zhuanlan.zhihu.com/p/143535541