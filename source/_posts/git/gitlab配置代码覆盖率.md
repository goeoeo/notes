---
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

Test coverage parsing配置为: coverage: \d+.\d+% of statements  

Pipeline status 中有显示徽标的代码配置到md文件中即可  

