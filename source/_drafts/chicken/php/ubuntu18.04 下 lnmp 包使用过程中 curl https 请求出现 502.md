#### 问题原因
查看phpinfo 发现curl扩展的SSL Version版本为:Openssl1.1.0, 在ubuntu 中,ubuntu 18.04+版本采用的是最新的openssl 1.1，之前的ubuntu基本都是1.0的版本，openssl1.1安装低于php7的版本
会报错.
#### 解决
##### 1.失败的解决过程
最先我想是为curl指定ssl的版本,然后再将此curl编译到php的版本中,可是遇到一下问题.

* 编译curl 指定ssl版本的时候报错,各种错误,最后摸索到下面这条命令,编译成功
```bash
LIBS="-ldl" PKG_CONFIG_PATH=/usr/local/openssl/lib/pkgconfig ./configure --with-ssl --with-libssl-prefix=/usr/local/openssl --disable-shared --prefix=/usr/local/curl
```
* 当我尝试将编译安装好的curl,编译进php5.6的时候,还是报错这条路是不好走.

#####2.成功的解决过程
我重新安装系统,在执行lnmp多版本安装的时候提示我
Please reinstall the libcurl distribution 缺少这个库
执行命令
```bash
sudo apt-get install libcurl4-gnutls-dev
```
然后进行lnmp包的多版本安装,此时php5版本的curl https访问竟然成功了,我去查看phpinfo,发现SSL Version版本为:GnuTLS/3.5.18
回想一下,应该是libcurl这库的版本决定了SSL Versionl的版本

#### file_get_contents无法获取到https开头的链接内容问题
解决:
```php
$options = array(
            'http' => array(
                'method' => 'POST',
                'header' => 'Content-type:application/x-www-form-urlencoded;charset=UTF-8',
                'content' => $stringData
            ),
            // 解决SSL证书验证失败的问题
            "ssl"=>array(
                "verify_peer"=>false,
                "verify_peer_name"=>false,
            )
        );
        $context = stream_context_create($options);
        $data = file_get_contents($url, false, $context);

```



