# 查询所有账号信息
```
SELECT DISTINCT a.`User`,a.`Host`,a.password_expired,a.password_last_changed,a.password_lifetime,a.* FROM mysql.user a;
```

# mysql8 创建帐号
1. 创建用户  
```
create user 'slave'@'%' identified by '密码';
# 刷新权限
flush privileges;
```
2. 授权
```

grant all privileges on *.* to 'slave'@'%' with grant option;
# all privileges 可换成select,update,insert,delete,drop,create等操作 
# 如：grant select on *.* to 'slave'@'%';
# *.* 第一个*表示通配数据库,第二个*表示通配表
# with gran option表示该用户可给其它用户赋予权限，但不可能超过该用户已有的权限

```
3. 查看用户授权信息  
```
show grants for 'slave'@'%';
```
4. 撤销权限
```
revoke all privileges on *.* from 'slave'@'%';
```
5. 删除用户  
```
drop user 'slave'@'slave';
```

# mysql导入导出

## 导出
mysqldump  -uroot -pp@ss52Dnb -h192.168.49.2 -P 32316 subscription bill_0 > ./subscription_bill_0.sql

## 导入
mysql -uroot -pp@ss52Dnb -h192.168.49.2 -P 32316   
use subscription;  
source ./subscription_bill.sql  
