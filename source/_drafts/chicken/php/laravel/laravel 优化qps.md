* 关闭应用debug app.debug=false
* 缓存配置信息 php artisan config:cache
* 缓存路由信息 php artisan router:cache
* 类映射加载优化 php artisan optimize
* 自动加载优化 composer dumpautoload
* 根据需要只加载必要的中间件
* 使用即时编译器（JIT），如：HHVM、OPcache
* 使用 PHP 7.x

* 开启opcache
opcache 一定要开,开启与关闭 OPcache，数据上竟有几倍的差别。我测试的时候有大概8,9倍的餐具
```bash
zend_extension=opcache.so
opcache.enable=1
opcache.enable_cli=1
```