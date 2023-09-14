---
title: MySQL在什么情况下会使用到临时表

categories: 
- mysql
tags:
- mysql临时表
---

MySQL在执行SQL查询时可能会用到临时表，一般情况下，用到临时表就意味着性能较低。临时表存储，MySQL临时表分为“内存临时表”和“磁盘临时表”，
其中内存临时表使用MySQL的MEMORY存储引擎，磁盘临时表使用MySQL的MyISAM存储引擎；
<!--more-->

 一般情况下，MySQL会先创建内存临时表，但内存临时表超过配置指定的值后，MySQL会将内存临时表导出到磁盘临时表；
 Linux平台上缺省是/tmp目录，/tmp目录小的系统要注意啦。  
	 
### 使用临时表的场景
* ORDER BY子句和GROUP BY子句不同， 例如：ORDERY BY price GROUP BY name；   
* 在JOIN查询中，ORDER BY或者GROUP BY使用了不是第一个表的列 例如：SELECT * from TableA, TableB ORDER BY TableA.price GROUP by TableB.name
（我在测试的时候发现即使是使用第一个表的列，也会使用到临时表）
* ORDER BY中使用了DISTINCT关键字 ORDERY BY DISTINCT(price)  
* SELECT语句中指定了SQL_SMALL_RESULT关键字 SQL_SMALL_RESULT的意思就是告诉MySQL，结果会很小，请直接使用内存临时表，
不需要使用索引排序 SQL_SMALL_RESULT必须和GROUP BY、DISTINCT或DISTINCTROW一起使用 一般情况下，我们没有必要使用这个选项，
让MySQL服务器选择即可。  
* 当group by 索引的时候，当数据量较小的时候不会使用临时表，当数据量（我测试的数据量在30万左右）大的时候会使用临时表
* 当group by 不是索引的时候会使用临时表  

### 直接使用磁盘临时表的场景
* 表包含TEXT或者BLOB列；
* GROUP BY 或者 DISTINCT 子句中包含长度大于512字节的列
* 使用UNION或者UNION ALL时，SELECT子句中包含大于512字节的列；

### 临时表相关配置
* tmp_table_size 指定系统创建的内存临时表最大大小；
* max_heap_table_size 指定用户创建的内存表的最大大小；
> 注意：最终的系统创建的内存临时表大小是取上述两个配置值的最小值。

### 表的设计原则
使用临时表一般都意味着性能比较低，特别是使用磁盘临时表，性能更慢，因此我们在实际应用中应该尽量避免临时表的使用。 
常见的避免临时表的方法有：
* 创建索引：在ORDER BY或者GROUP BY的列上创建索引；

* 分拆很长的列：一般情况下，TEXT、BLOB，大于512字节的字符串，基本上都是为了显示信息，而不会用于查询条件， 因此表设计的时候，
应该将这些列独立到另外一张表。

### SQL优化
如果表的设计已经确定，修改比较困难，那么也可以通过优化SQL语句来减少临时表的大小，以提升SQL执行效率。常见的优化SQL语句方法如下：
* 拆分SQL语句，临时表主要是用于排序和分组，很多业务都是要求排序后再取出详细的分页数据，这种情况下可以将排序和取出详细数据拆分成不同的SQL，
以降低排序或分组时临时表的大小，提升排序和分组的效率，我们的案例就是采用这种方法。
* 优化业务，去掉排序分组等操作，有时候业务其实并不需要排序或分组，仅仅是为了好看或者阅读方便而进行了排序，例如数据导出、数据查询等操作，
这种情况下去掉排序和分组对业务也没有多大影响。

### 如何判断使用了临时表
使用explain查看执行计划，Extra列看到Using temporary就意味着使用了临时表。