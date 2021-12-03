1.将angular程序打包好，放到网站根目下的static目录中

2.在angular打包的时候注意要将base_url设置成/static/

3.nginx的配置，因为是静态文件，所以我这里开启了压缩
```
 location /static {
        # First attempt to serve request as file, then
        sendfile on;
        gzip_static on;
        gzip_http_version 1.1;
        gzip_proxied expired no-cache no-store private auth;
        gzip_disable "MSIE [1-6] .";
        gzip_vary on;

        gzip  on;
        gzip_min_length 1k;
        gzip_buffers 4 16k;
        gzip_comp_level 2;
        gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript application/x-httpd-php image/jpeg image/gif image/png;
        gzip_disable "MSIE [1-6]\.";

        try_files $uri $uri/ /www/index.html;
   }
```
4.将网站根目录重定向到static目录下 
```
 location / {
        index index.php  index.html index.htm;
        if ( $request_uri = "/") {
                rewrite / /static/index.html break;
        }
        try_files $uri $uri/ /index.php?$query_string;#这里是laravel的请求配置
    }
```
 

