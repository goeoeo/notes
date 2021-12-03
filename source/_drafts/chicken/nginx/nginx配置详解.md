
# nginx 中文文档网站
[nginx.cn](https://nginx.cn)
解锁文章: localstorage设置 TOKEN_19093-1583317806202-845=>19093-1583317806202-845

# nginx默认配置文件

## nginx.conf
```

#运行用户
user       www www;  ## Default: nobody 

#工作进程，通常设置成和cpu的数量相等
worker_processes  5;  ## Default: 1 

#全局错误日志及PID文件
error_log  logs/error.log;
pid        logs/nginx.pid;

# nginx 进程打开的最多文件描述符数目
worker_rlimit_nofile 8192;

events {
  #epoll是多路复用IO(I/O Multiplexing)中的一种方式,但是仅用于linux2.6以上内核,可以大大提高nginx的性能
  use  epoll;       

  #单个后台worker process进程的最大并发链接数
  #通过worker_connections和worker_proceses可以计算出maxclients  max_clients = worker_processes * worker_connections
  # 作为反向代理，max_clients=worker_processes * worker_connections/4
  worker_connections  4096;  ## Default: 1024
}

#设定http服务器，利用它的反向代理功能提供负载均衡支持
http {
  include    conf/mime.types;
  include    /etc/nginx/proxy.conf;
  include    /etc/nginx/fastcgi.conf;
  index    index.html index.htm index.php;

  #nginx默认文件类型
  #Nginx 会根据mime type定义的对应关系来告诉浏览器如何处理服务器传给浏览器的这个文件，是打开还是下载
  #如果Web程序没设置，Nginx也没对应文件的扩展名，就用Nginx 里默认的 default_type定义的处理方式。
  default_type application/octet-stream;

  #设定日志格式
  log_format   main '$remote_addr - $remote_user [$time_local]  $status '
    '"$request" $body_bytes_sent "$http_referer" '
    '"$http_user_agent" "$http_x_forwarded_for"';
  access_log   logs/access.log  main;

  #sendfile 指令指定 nginx 是否调用 sendfile 函数（zero copy 方式）来输出文件，对于普通应用，必须设为 on
  #如果用来进行下载等应用磁盘IO重负载应用，可设置为 off，以平衡磁盘与网络I/O处理速度，降低系统的uptime.
  sendfile     on; #服务器若开启sendfile可以降低系统开销，例如在内核空间中完成数据由磁盘读入缓存转移到sokcet关联的缓存，避免多余的上下文切换。
  tcp_nopush   on; # 开启后在连接中有未确认包时依然可以发送新包，在sendfile=on的情况下生效
  tcp_nodely   on; # 开启后当tcp填满后，不需要等待上一个ACK就立即发送
  server_names_hash_bucket_size 128; # 定义的server_name长度
  server_names_hash_max_size 512; #定义server_name数量 

  

  server { # php/fastcgi
    listen       80;
    server_name  domain1.com www.domain1.com;
    access_log   logs/domain1.access.log  main;
    root         html;  #定义服务器的默认网站根目录位置

    location ~ \.php$ {
      fastcgi_pass   127.0.0.1:1025;
    }
  }

  server { # simple reverse-proxy
    listen       80;
    server_name  domain2.com www.domain2.com;
    access_log   logs/domain2.access.log  main;

    # 静态文件，nginx自己处理
    location ~ ^/(images|javascript|js|css|flash|media|static)/  {
      root    /var/www/virtual/big.server.com/htdocs;
      expires 30d;
    }

    # pass requests for dynamic content to rails/turbogears/zope, et al
    location / {
      proxy_pass      http://127.0.0.1:8080;
    }
  }

  upstream big_server_com {
    server 127.0.0.3:8000 weight=5;
    server 127.0.0.3:8001 weight=5;
    server 192.168.0.1:8000;
    server 192.168.0.1:8001;
  }

  server { # simple load balancing
    listen          80;
    server_name     big.server.com;
    access_log      logs/big.server.access.log main;

    location / {
      proxy_pass      http://big_server_com;
    }
  }
}
```

# proxy.conf
```
proxy_redirect          off;
proxy_set_header        Host            $host;
proxy_set_header        X-Real-IP       $remote_addr;
proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
client_max_body_size    10m;
client_body_buffer_size 128k;
proxy_connect_timeout   90;
proxy_send_timeout      90;
proxy_read_timeout      90;
proxy_buffers           32 4k;
```

# fastcgi.conf
```
fastcgi_param  SCRIPT_FILENAME    $document_root$fastcgi_script_name;
fastcgi_param  QUERY_STRING       $query_string;
fastcgi_param  REQUEST_METHOD     $request_method;
fastcgi_param  CONTENT_TYPE       $content_type;
fastcgi_param  CONTENT_LENGTH     $content_length;
fastcgi_param  SCRIPT_NAME        $fastcgi_script_name;
fastcgi_param  REQUEST_URI        $request_uri;
fastcgi_param  DOCUMENT_URI       $document_uri;
fastcgi_param  DOCUMENT_ROOT      $document_root;
fastcgi_param  SERVER_PROTOCOL    $server_protocol;
fastcgi_param  GATEWAY_INTERFACE  CGI/1.1;
fastcgi_param  SERVER_SOFTWARE    nginx/$nginx_version;
fastcgi_param  REMOTE_ADDR        $remote_addr;
fastcgi_param  REMOTE_PORT        $remote_port;
fastcgi_param  SERVER_ADDR        $server_addr;
fastcgi_param  SERVER_PORT        $server_port;
fastcgi_param  SERVER_NAME        $server_name;

fastcgi_index  index.php;

fastcgi_param  REDIRECT_STATUS    200;
```

# mime.types
```
types {
  text/html                             html htm shtml;
  text/css                              css;
  text/xml                              xml rss;
  image/gif                             gif;
  image/jpeg                            jpeg jpg;
  application/x-javascript              js;
  text/plain                            txt;
  text/x-component                      htc;
  text/mathml                           mml;
  image/png                             png;
  image/x-icon                          ico;
  image/x-jng                           jng;
  image/vnd.wap.wbmp                    wbmp;
  application/java-archive              jar war ear;
  application/mac-binhex40              hqx;
  application/pdf                       pdf;
  application/x-cocoa                   cco;
  application/x-java-archive-diff       jardiff;
  application/x-java-jnlp-file          jnlp;
  application/x-makeself                run;
  application/x-perl                    pl pm;
  application/x-pilot                   prc pdb;
  application/x-rar-compressed          rar;
  application/x-redhat-package-manager  rpm;
  application/x-sea                     sea;
  application/x-shockwave-flash         swf;
  application/x-stuffit                 sit;
  application/x-tcl                     tcl tk;
  application/x-x5d09-ca-cert            der pem crt;
  application/x-xpinstall               xpi;
  application/zip                       zip;
  application/octet-stream              deb;
  application/octet-stream              bin exe dll;
  application/octet-stream              dmg;
  application/octet-stream              eot;
  application/octet-stream              iso img;
  application/octet-stream              msi msp msm;
  audio/mpeg                            mp3;
  audio/x-realaudio                     ra;
  video/mpeg                            mpeg mpg;
  video/quicktime                       mov;
  video/x-flv                           flv;
  video/x-msvideo                       avi;
  video/x-ms-wmv                        wmv;
  video/x-ms-asf                        asx asf;
  video/x-mng                           mng;
}
```

# sendfile tcp_nodelay, tcp_nopush

[tcp_nodelay,tcp_nopush](https://blog.zp25.ninja/miscellaneous/2019/01/18/nodelay-nopush.html)

## sendfile
使用read, write读写数据需要在两个系统调用间切换回用户空间，切换过程中数据需要从内核空间转移到用户空间。使用sendfile数据直接在内核空间
中完成两个文件描述符(file descriptors)间的传递。  

服务器若开启sendfile可以降低系统开销，例如在内核空间中完成数据由磁盘读入缓存转移到sokcet关联的缓存，避免多余的上下文切换。
## Nagle’s algorithm (纳格算法)
纳格算法(Nagle’s algorithm)通过减少数据包发送量来增进TCP/IP网络的性能。  

纳格算法主要用于解决小包问题(small-packet problem)，例如连续发送1byte数据块的应用(Telnet)，因为TCP和IP包头部合计40bytes，
整个包有效数据占比极低，并且发送大量小包可能造成网络拥塞。纳格算法可以累积多个小块数据后以一个TCP报文发送，实现方式是要求同一时刻连接中
仅存在一个未确认包，在等待ACK期间持续缓冲小块数据准备下一次发送。  

不适宜使用纳格算法的情形，例如发送数据量大于MSS，发送最后一段数据(小于MSS)前需等待接收之前所有包的ACK。或者遇上写-写-读情形，向发送缓存
写入两个小块数据(小于MSS)后等待，因为连接中没有未确认包，第一个小块数据会立即发送，但第二小块数据需要等待ACK后发送。  
纳格算法和延时确认一起使用会造成额外延时。  
## 延时确认
延时确认(TCP delayed acknowledgment)是另一种减少数据包发送量的技术，接收端可以选择延迟发送ACK，等待将多个ACK合并或与响应数据一起发送。
延时最长不超过500ms。  

例如配置延时40ms，接收端接收到第一个TCP包后等待最长40ms，若期间接收到另一个TCP包，两个ACK将通过一个包发送。若等待超时，接收端会发送单个ACK，
若有响应数据需要发送，等待时间也利于响应数据累积，ACK和响应数据将一并发送。  

同时应用纳格算法和延时确认会造成额外的延时。因为纳格算法同一时刻仅允许一个未确认包，延时确认却将ACK延时发送，在发送端缓存数据不足MSS时，
总需要等待延时确认超时。  

## tcp_nodelay, tcp_nopush
Linux默认开启纳格算法解决小包问题，可通过TCP_NODELAY选项(per-socket)关闭。nginx通过tcp_nodelay管理TCP_NODELAY选项，默认tcp_nodelay on;
关闭纳格算法，因为web服务器更多发送大块数据，关闭纳格算法可以减小时延。  

纳格算法并不关心发送包的大小，第一个包或接收到ACK时，无论缓存数据是否到达MSS都立即发送。Linux另一个socket选项TCP_CORK可用于提高每个TCP包的利用率。
若设置TCP_CORK选项，socket类似于堵上了软木塞子(cork)，只有累积了足够数据才允许发送，例如达到MSS。或者需要发送小块数据时移除塞子，
即清理TCP_CORK选项。  

nginx通过tcp_nopush管理TCP_CORK选项，默认tcp_nopush: off;，需要在sendfile on;时开启才有效。开启后，nginx会等待缓存中数据累积
到MSS才发送，若到达最后一段数据，nginx会移除TCP_CORK选项使小块数据立即发送。

```
sendfile     on;
tcp_nodelay  on;
tcp_nopush   on;
```

我最初看到这样的配置时非常疑惑，一个设置立即发送，另一个设置不立即发送，而且还同时开启了。事实是两个配置并不冲突，
立即发送是在连接中有未确认包时依然可以发送新包，而不立即发送是需要充分利用一个TCP包的荷载。即先要填满一个包，
然后不需要等待之前包的ACK就可以立即发送。  


