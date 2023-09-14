---
title: docker-machine

categories: 
- docker

tags:
- docker-install-ware
---

## 下载
```
base=https://github.com/docker/machine/releases/download/v0.16.0 &&
  curl -L $base/docker-machine-$(uname -s)-$(uname -m) >/tmp/docker-machine &&
  sudo mv /tmp/docker-machine /usr/local/bin/docker-machine &&
  chmod +x /usr/local/bin/docker-mach
```

### 网速慢
使用xx_net代理
```
sudo  apt install privoxy
export https_proxy=127.0.0.1:8087
export http_proxy=127.0.0.1:8087 
```

### x509问题
curl -k 参数跳过证书验证


## 使用
```
docker-machine create -d  virtualbox manager1

输出：
Running pre-create checks...
(manager1) No default Boot2Docker ISO found locally, downloading the latest release...
Error with pre-create check: "Get https://api.github.com/repos/boot2docker/boot2docker/releases/latest: x509: certificate signed by unknown authority"

```
### 直接去下载默认的iso文件
```
curl -Lo ~/.docker/machine/cache/boot2docker.iso https://github.com/boot2docker/boot2docker/releases/download/v19.03.5/boot2docker.iso

输出：
curl: (60) SSL certificate problem: unable to get local issuer certificate
More details here: https://curl.haxx.se/docs/sslcerts.html

```


### curl: (60) SSL certificate problem: unable to get local issuer certificate问题
```
curl -k --remote-name --time-cond cacert.pem https://curl.haxx.se/ca/cacert.pem
# 注意执行此命令后需要重新开终端
```

###　终端下载实在太慢问题
1. [浏览器下载](https://github.com/boot2docker/boot2docker/releases/download/v19.03.5/boot2docker.iso)
2. 将iso文件移动到　 ~/.docker/machine/cache/boot2docker.iso