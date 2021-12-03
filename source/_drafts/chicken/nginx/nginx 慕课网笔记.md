1.nginx 是一个高效，可靠的http中间件，代理服务。
2.进程是应用程序的实例，每个进程至少包含一个主线程，线程是操作系统的执行单元。线程用于执行进程中的代码。
3.socket是对TCP/IP协议的封装，socket本身不是协议，而是一个调用接口。网络上的两个程序通过一个双向的通信连接实现数据的交换，这个连接的一端称为一个socket。
HTTP是轿车，提供了封装或者显示数据的具体形式；Socket是发动机，提供了网络通信的能力。传输层的TCP基于网络层的IP,应用层的HTTP基于TCP，三次握手。
问题1：为什么选择nginx
   1.IO多路复用epoll
   2.轻量级：功能模块少，模块化代码
   3.cpu亲和：将cpu核心和nginx工作进程进行绑定，把每个worker进程固定到cpu的一个核心上面减少cpu切换带来的cache miss ，获得更好的性能
   4.sendfile:一般的服务器，静态文件传递到socket，需要经过，内核空间，用户空间。而sendfile可以忽略用户空间，直接通过内核空间，就可传递到socket;
----------
模块介绍：
1.http_stub_status_module：用于监控nginx 运行状态
2.http_random_index_module:从主目录中随机选则一个主页进行展示
3.http_sub_module:用于替换网页中的内容
----------
nginx 访问控制
1.基于IP的访问控制：http_access_module (allow deny)
 默认是通过remote_addr实现的。有一定的局限性：解决方案：
    1:采用别的http头信息，如http_x_forwarded_for
    2:结合geo模块
    3:通过Http自定义变量传递
2.基于用户的信任登录：http_auth_basic_module
    auth_basic auth_basic_file
    局限性：依赖文件
    解决方案：
    1：nginx结合lua实现高效验证
    2.nginx和LDAP打通，利用nginx-auth-ldap模块
--------------------
常见的nginx的中间件架构
一。静态资源web服务
1.静态资源web服务
2.代理服务
3.负载均衡调度器SLB
4.动态缓存
压缩：gzip,
压缩比：gzip_comp_level
http协议版本：gzip_http_version
预读nginx功能：http_gzip_static_module
控制客户端缓存时间：expires
防盗链（基于http_refer）：
 valid_referers none blocked ~/google\./;#允许referer从google过来的。
 if($invalid_referer){
    return 403
 }
二。代理服务
1.正向代理：代理的对象是客户端。
2.反向代理：代理的对象是服务端。
常用代理配置
  location / {
    proxy_pass http://127.0.0.1:8000
    proxy_redirect default;

    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;

    proxy_connect_timeout 30;
    proxy_send_timeout 60;
    proxy_read_timeout 60;

    proxy_buffer_size 32k;
    proxy_buffering on;
    proxy_buffers 4 128k;
    proxy_busy_buffers_size 256k;
    proxy_max_temp_file_size 256k;
  }
---------------------------------------------------------------
负载均衡
    根据网络模型分为4层负载均衡和7层负载均衡
    4层负载均衡：基于TCP/IP协议进行包的转发。
    7层负载均衡：在应用层进行处理，修改Http头信息，转发，重定向。nginx是典型的7层负载均衡的slb
nginx负载均衡
    proxy_pass,upstream,
    upstream需要配置在server层以外
    配置语法：
    upstream backend{
        server backend1.example.com weight=5;
        server backend2.example.com:8000 ;
        server unix:/tmp/backend3
    }
    后端服务器在负载均衡调度中的状态：
        down:当前的server暂时不参与负载均衡
        backup:预留的备份服务器
        max_fails:允许请求失败的次数
        fail_timeout:经过max_fails次的失败后，暂停服务的时间
        max_conns:限制最大的接受的连接数

nginx代理调度算法
    轮询（默认） ：按时间顺序逐一分配到不同的后端服务器。
    加权轮询：weight值越大，分配到的访问几率越高。
    ip_hash:每个请求按访问IP的hash结果分配，这样来自同一个IP的客户端则可以固定访问一个后端服务器。
    （可以解决session问题,缺点，当前端还有一台代理服务器的时候，remote_addr始终是一个值，） 开启语法：ip_hash;
    url_hash:按照访问的URL的hash结果分配，每个url定向到一个后端服务器 开启语法：hash $request_uri;
    least_conn:最少链接数，那个机器连接数最少就分发。
    hash关键数值：hash自定义的key

    upstream_hash（第三方模块）
    为了解决ip_hash的一些问题，可以使用upstream_hash这个第三方模块，这个模块多数情况下是用作url_hash的，但是并不妨碍将它用来做session共享：
    假如前端是squid，他会将ip加入x_forwarded_for这个http_header里，用upstream_hash可以用这个头做因子，将请求定向到指定的后端：

    开启语法：hash $http_x_forwarded_for;

    这样就改成了利用x_forwarded_for这个头作因子，在nginx新版本中可支持读取cookie值，所以也可以改成：

    hash $cookie_jsessionid;
------------------------------------------------------------------------------------------------------------------
nginx 作为缓存服务
 客户端缓存，代理缓存，服务端缓存
 客户端nginx服务端
    proxy_cache配置语法

    proxy_cache_path /opt/app/cache levels=1:2 keys_zone=imooc_cache:10m max_size=10g inactive=60m use_temp_path=off;
    （proxy_cache_path:缓存存放的目录，levels:缓存目录的级别，keys_zone:为缓存key值开辟的空间，1m可大约存放8000个key值。max_size:
    缓存目录最大的存储空间，当缓存空间达到这个大小的时候，nginx会开启自己的垃圾回收机制，inactive:当回收机制开始的时候，60m内不活跃的将被剔除，
    use_temp_path:是否使用系统自身的缓存目录，由于前面已经设置，所以这里需要关掉）
    location /{
        proxy_cache imooc_cache;#开启缓存
        proxy_pass http://imooc;#代理
        proxy_cache_valid 200 304 12h;#对于200,304返回状态码的过期时间为12小时
        proxy_cache_valid any 10m;#对于任意的状态码的过期时间是10分钟
        proxy_cache_key $host$uri$is_args$args;#缓存的key值的命名规则
        add_header Nginx-Cache "$upstream_cache_status";#在响应头信息中显示，告诉客户端是否命中。

        proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;#当一台代理服务器出现问题，直接跳过读取下一台。
    }

    设置不缓存路径
    server{
       if($request_uri ~ ^/(url3|login)){
            set $cookie_nocache 1;
       }
       locatin /{
            proxy_no_cache $cookie_nocache $arg_nocache $arg_comment;
            proxy_no_cache $http_pragma $http_authorization;
       }
    }
nginx 作为缓存服务-分片请求

-------------------------------------------------------------------------------------------------------------------------
nginx 动静分离---分离资源，减少不必要的请求消耗，减少请求延时。
-----------------------------------------------------------------------------------------------------------------------
nginx rewrite，（可对url进行重写，或对匹配的url进行重定向。）
场景：
    1.url跳转访问，支持开发设计（页面跳转，兼容性支持，展示效果）
    2.SEO优化
    3.维护（当服务器需要维护的时候，直接可利用rewrite规则，重定向到维护页面）
    4.安全
配置语法
    rewrite regex replacement [flag]

    flag:
        last:停止rewrite检测，使用此描述符时，rewrite时，会重新执行一次http请求
        break:停止rewrite检测，使用此描述符时，rewrite时，会直接去查找文件，未找到则会报404；
        redirect:返回302临时重定向，地址栏会显示跳转后的地址。
        permanent:返回301永久重定向，地址栏会显示跳转后的地址。永久重定向，客户端会永久的保留重定向的结果。

实例
    简写url :/course/11/22/course_33.html => /course-11-22-33.html
    location /{
        rewrite ^/course-(\d+)-(\d+)-(\d+)\.html$ /course/$1/$2/course_$3.html break;
    }
    对于没有找到文件的请求的重定向
    location /{
        if(! -f $request_filename){
            rewrite ^/(.*)$ http://www.baidu.com/$1 redirect;
        }
    }

------------------------------------------------------------------------------------------------------------------------
nginx 高级模块
1.secure_link_module模块
  1）。制定并允许检查请求链接的真实性以及保护资源免遭未经授权的访问
  2）。限制链接的生效周期。
  配置语法
  secure_link expression;
  secure_link_md5 expression;
2.geoip_module模块
  1）。区别国内国http访问规则
  2）。区别国内城市做http访问规则
------------------------------------------------------------------------------------------------------------------------
https服务

------------------------------------------------------------------------------------------------------------------------
nginx 架构篇

常见问题
1.多个server_name中虚拟主机读取的优先级，nginx 按照优先读取文件的方式进行读取
2.location匹配优先级
    1).精确匹配，=
    2）。前缀匹配^~
    3).正则匹配 ~/~*
3.nginx 中的try_files:按顺序检查文件是否存在
try_files $uri $uri/ /index.php?$query_string;
4.alias和root的区别

5.如何获取用户的真实IP;
在第一级代理的时候设置头信息，set x_real_ip=$remote_addr;这样在后端服务器中直接读取这个变量即可。

nginx常见错误码
1.413 Request Entity Too Large (请求内容内容过大)
 处理方案：修改用户上传文件限制 client_max_body_size
2.502 bad gateway (后台网关错误，后端服务没有响应，可能服务进程已经down掉了)
3.504 gateway time-out(网关超时，后端服务执行超时，后台服务是存活的，但由于某些原因（错误的递归，从数据库取数据过慢），后端服务器处理的时间过长，超过了nginx的keepalive_timeout（60秒）)


性能优化
ab压测工具
 安装 httpd-tools
 参数 ： -n 总的请求次数，-c 并发量
 ab -n 2000 -c 2 http://www.rentmaterial.com/www/

文件句柄设置
全局：/etc/security/limits.conf
针对nginx进程：worker_rlimit_nofile 65535;
CPU亲和配置：把进程通常不会在处理器之间频繁迁移的频率减小，减少性能损耗。

nginx通用配置优化
user nginx;//普通用户
worker_processes 8;//配置成cpu的核数
worker_cpu_affinity auto;//cpu亲和
worker_rlimit_nofile 10000;//进程级别的文件句柄数量，小站点10000,大点的站点20000到30000；

use epoll;//使用epoll模型
worker_connections 10240;//每个worker进程的最大链接数，nginx默认1024,需要调大一下

charset utf-8;//字符集
access_log off;//关闭需要打印的访问日志
sendfile on;//优化静态资源响应
gzip  on;//开启压缩
gzip_disable "MSIE [1-6]\.";//对于ie1-ie6关闭压缩
gzip_http_version 1.1;

------------------------------------------------------------------------------------------------------------------------
nginx 安全
nginx 文件上传漏洞
对于这样的链接：http://www.imooc.com/upload/1.jpg/1.php，nginx会将1.jpg不会当做静态资源处理，而是会将1.jpg当为php文件进行处理，。
解决方案：
location ^~ /upload{
    root /opt/app/images;
    if($request_filename ~ *(.*)\.php){
        return 403;
    }
}
-----------------------------------------------------------------------------------------------------------------
静态资源服务的功能设计
1.定义nginx在服务体系中的角色
  1).静态资源服务。2).代理服务，3).动静分离

  作为静态资源服务需要考虑的因素：浏览器缓存，资源类型分类，防盗链，流量限制，压缩，防资源盗用。
  作为代理服务需要考虑的因素：正向代理，反向代理，负载均衡，LNMP，，Poxypass，头信息处理，代理缓存，分片请求，协议类型
2.设计评估
  硬件：CPU，内存，硬盘
  系统：用户权限，日志目录存放
  关联服务：LVS，keepalive,syslog,Fastcgi










