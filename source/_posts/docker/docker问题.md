# docker本地空间足够，但容器日志提示空间不足

## 情况描述
项目中，es报空间不足，es容器 已使用30G左右的空间
默认情况下docker 对每个容器使用空间有限制  
```
Elasticsearch：high disk watermark [90%] exceeded
```

## 处理
修改配置   
vim /etc/docker/daemon.json   
```
{
    "storage-opt": [ "dm.basesize=80G" ]
} 
```
重启 
```
sudo systemctl daemon-reload &&  sudo systemctl restart docker
```

# 参考
* [收藏！24 个 Docker 疑难杂症处理技巧](https://www.bilibili.com/read/cv16472262)
