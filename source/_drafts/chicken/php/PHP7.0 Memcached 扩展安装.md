#### 编译安装
1.安装依赖库 libmemcached #必须安装，不然后面编译memcached扩展会报错
```
$ wget https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.tar.gz
$ tar -zxf libmemcached-1.0.18.tar.gz
$ cd libmemcached-1.0.18/
$ ./configure
$ make && make install
注意：在ubuntu18下，make会报错
```
2.安装php7的memcached扩展
```
$ git clone https://github.com/php-memcached-dev/php-memcached.git
$ cd php-memcached/
$ git checkout php7
$ phpize  # 如果未安装php-dev需先安装
$ ./configure --disable-memcached-sasl
$ make && make install

```
如果编译成功会显示出扩展安装路径,例如：/usr/lib/php/20151012/


3.启动扩展
	1.在php.ini末尾加入 


	extension_dir="/usr/lib/php/20151012/"
	extension="memcached.so"

	2.fpm路径：/etc/php/7.0/fpm/php.ini   #修改了这个配置文件，web访问才能生效，需要使用service php7.0-fpm restart 重新加载配置文件。

	3.cli路径:/etc/php/7.0/cli/php.ini #在使用laravel的是否，缓存系统使用的是memcached的时候，使用composer update的时候，是用的cli来进行扩展的检查的，composer update 会报错。

	4.使用 php -m 查看扩展是否加载成功 

#### pecl安装
```bash
pecl install memcached
```
注意ubuntu18下需要先安装libmemcached
```bash
sudo apt install libmemcached-dev
```









