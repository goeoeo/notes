---
categories:
- tidb


tags:
- tidb
---
# tidb 数据迁移

k8s上部署 tidb 数据迁移工具dm  

方案: 全量迁移+增量迁移

前置条件:  

- TiDB Operator [部署](https://docs.pingcap.com/zh/tidb-in-kubernetes/dev/deploy-tidb-operator)完成。

- 要求 TiDB Operator 版本 >= 1.2.0。

  



## 1.dm-cluster.yaml

参考 DMCluster [示例](https://github.com/pingcap/tidb-operator/blob/master/examples/dm/dm-cluster.yaml)

创建文件: dm-cluster.yaml 内容:

```yaml
apiVersion: pingcap.com/v1alpha1
kind: DMCluster
metadata:
  name: basic
spec:
  version: v2.0.6
  pvReclaimPolicy: Retain
  discovery: {}
  master:
    baseImage: pingcap/dm
    replicas: 1
    # if storageClassName is not set, the default Storage Class of the Kubernetes cluster will be used
    # storageClassName: local-storage
    storageSize: "1Gi"
    requests: {}
    config: {}
  worker:
    baseImage: pingcap/dm
    replicas: 1
    # if storageClassName is not set, the default Storage Class of the Kubernetes cluster will be used
    # storageClassName: local-storage
    storageSize: "1Gi"
    requests: {}
    config: {}
```



## 2.部署 DM 集群

```shell
kubectl apply -f dm_cluster.yaml -n beatflow-data
```

## 3.期望输出

 kubectl get pods -n beatflow-data

```
NAME                                      READY   STATUS    RESTARTS   AGE
basic-dm-discovery-6b99d57d5c-w9qr7       1/1     Running   0          18h
basic-dm-master-0                         1/1     Running   0          18h
basic-dm-worker-0                         1/1     Running   0          18h

```

## 4.启动 DM 同步任务

通过进入 DM-master 或 DM-worker pod 使用 image 内置 dmctl 进行操作。  

```shell
kubectl exec -it basic-dm-master-0 -n beatflow-data -- sh
```

### 1.创建数据源（mysql）

 vi source1.yaml  

```yaml
# MySQL1 Configuration.

source-id: "mysql-replica-01"

# DM-worker 是否使用全局事务标识符 (GTID) 拉取 binlog。使用前提是在上游 MySQL 已开启 GTID 模式。
enable-gtid: false

from:
  host: "mysql.beatflow-data.svc"
  user: "root"
  password: "p@ss52Dnb"
  port: 3306
```



```
./dmctl --master-addr 127.0.0.1:8261 operate-source create source1.yaml
```



### 2.配置同步任务

vi task.yaml

```yaml
# 任务名，多个同时运行的任务不能重名。
name: "test"
# 全量+增量 (all) 迁移模式。
task-mode: "all"
# 下游 TiDB 配置信息。
target-database:
  host: "basic-tidb.beatflow-data.svc"
  port: 4000
  user: "root"
  password: ""

# 当前数据迁移任务需要的全部上游 MySQL 实例配置。
mysql-instances:
-
  # 上游实例或者复制组 ID，参考 `inventory.ini` 的 `source_id` 或者 `dm-master.toml` 的 `source-id 配置`。
  source-id: "mysql-replica-01"
  # 需要迁移的库名或表名的黑白名单的配置项名称，用于引用全局的黑白名单配置，全局配置见下面的 `block-allow-list` 的配置。
  block-allow-list: "global"          # 如果 DM 版本早于 v2.0.0-beta.2 则使用 black-white-list。
  # dump 处理单元的配置项名称，用于引用全局的 dump 处理单元配置。
  mydumper-config-name: "global"


# 黑白名单全局配置，各实例通过配置项名引用。
block-allow-list:                     # 如果 DM 版本早于 v2.0.0-beta.2 则使用 black-white-list。
  global:
    do-dbs: ["*"]
    ignore-tables: 
    - db-name: "*"
      tbl-name: "PDMAN_DB_VERSION"
# dump 处理单元全局配置，各实例通过配置项名引用。
mydumpers:
  global:
    extra-args: ""

```

启动任务: 

```shell
./dmctl --master-addr 127.0.0.1:8261 start-task task.yaml
```

查询任务:

```shell
./dmctl --master-addr 127.0.0.1:8261 query-status test
```





## 遇到的问题

* block-allow-list 变更后需要重新命名任务名称，否则任务不会生效

* 由于包含全量复制，任务中途退出需要清空相关表，否则会遇到主键重复，task can't auto resume 的错误

* 查看同步日志 

  ```shell
  kubectl logs -f basic-dm-worker-0 -n beatflow-data
  ```

  