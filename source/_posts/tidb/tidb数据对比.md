---
categories:
- tidb


tags:
- tidb
---

## 配置文件

```yaml
# Diff Configuration.

######################### Global config #########################

# 日志级别，可以设置为 info、debug
log-level = "info"

# sync-diff-inspector 根据主键／唯一键／索引将数据划分为多个 chunk，
# 对每一个 chunk 的数据进行对比。使用 chunk-size 设置 chunk 的大小
chunk-size = 1000

# 检查数据的线程数量
check-thread-count = 4

# 抽样检查的比例，如果设置为 100 则检查全部数据
sample-percent = 100

# 通过计算 chunk 的 checksum 来对比数据，如果不开启则逐行对比数据
use-checksum = true

# 如果设置为 true 则只会通过计算 checksum 来校验数据，如果上下游的 checksum 不一致也不会查出数据再进行校验
only-use-checksum = false

# 是否使用上次校验的 checkpoint，如果开启，则只校验上次未校验以及校验失败的 chunk
use-checkpoint = true

# 不对比数据
ignore-data-check = false

# 不对比表结构
ignore-struct-check = false

# 保存用于修复数据的 sql 的文件名称
fix-sql-file = "fix.sql"

######################### Tables config #########################

# 如果需要对比大量的不同库名或者表名的表的数据，或者用于校验上游多个分表与下游总表的数据，可以通过 table-rule 来设置映射关系
# 可以只配置 schema 或者 table 的映射关系，也可以都配置
#[[table-rules]]
    # schema-pattern 和 table-pattern 支持通配符 *?
    #schema-pattern = "test_*"
    #table-pattern = "t_*"
    #target-schema = "test"
    #target-table = "t"

# 配置需要对比的*目标数据库*中的表
[[check-tables]]
    # 目标库中数据库的名称
    schema = "pricing"

    # 需要检查的表
    tables = ["filter"]

    # 支持使用正则表达式配置检查的表，需要以‘~’开始，
    # 下面的配置会检查所有表名以‘test’为前缀的表
    # tables = ["~^test.*"]
    # 下面的配置会检查配置库中所有的表
    # tables = ["~^"]

# 对部分表进行特殊的配置，配置的表必须包含在 check-tables 中

######################### Databases config #########################

# 源数据库实例的配置
[[source-db]]
    host = "192.168.49.2"
    port = 32316
    user = "root"
    password = "p@ss52Dnb"
    # 源数据库实例的 id，唯一标识一个数据库实例
    instance-id = "source-1"
    # 使用 TiDB 的 snapshot 功能，如果开启的话会使用历史数据进行对比
    # snapshot = "2016-10-08 16:45:26"
    # 设置数据库的 sql-mode，用于解析表结构
    # sql-mode = ""

# 目标数据库实例的配置
[target-db]
    host = "192.168.49.2"
    port = 30492
    user = "root"
    password = ""
    # 使用 TiDB 的 snapshot 功能，如果开启的话会使用历史数据进行对比
    # snapshot = "2016-10-08 16:45:26"
    # 设置数据库的 sql-mode，用于解析表结构
    # sql-mode = ""
```





## 命令

docker run --network=host    -v `pwd`:/home/tidb pingcap/tidb-enterprise-tools  ./sync_diff_inspector --config=/home/tidb/diff.yaml

