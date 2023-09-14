---
title: tidb安装

categories:
- tidb


tags:
- tidb
---

# TiDB 安装

[k8s官网安装地址](https://docs.pingcap.com/zh/tidb-in-kubernetes/stable)

## 1.安装 TiDB Operator CRDs

```shell
kubectl apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/manifests/crd.yaml -n beatflow-data
```



## 2.安装 TiDB Operator

```shell
helm repo add pingcap https://charts.pingcap.org/
```

```shell
helm install --namespace beatflow-data tidb-operator pingcap/tidb-operator --version v1.2.3
```

## 3.安装 TiDB 集群

```shell
kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/examples/basic/tidb-cluster.yaml -n beatflow-data
```

## 4. 安装 TiDB 集群监控

```shell
kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/examples/basic/tidb-monitor.yaml -n beatflow-data
```

## 5. 期望输出

```
NAME                              READY   STATUS    RESTARTS   AGE
basic-discovery-6bb656bfd-xl5pb   1/1     Running   0          9m9s
basic-monitor-5fc8589c89-gvgjj    3/3     Running   0          8m58s
basic-pd-0                        1/1     Running   0          9m8s
basic-tidb-0                      2/2     Running   0          7m14s
basic-tikv-0                      1/1     Running   0          8m13s
```

