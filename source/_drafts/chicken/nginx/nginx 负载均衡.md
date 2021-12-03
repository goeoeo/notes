linux负载均衡总结性说明（四层负载/七层负载）

一，什么是负载均衡
1）负载均衡（Load Balance）建立在现有网络结构之上，它提供了一种廉价有效透明的方法扩展网络设备和服务器的带宽、增加吞吐量、加强网络数据处理能力、提高网络的灵活性和可用性。负载均衡有两方面的含义：首先，大量的并发访问或数据流量分担到多台节点设备上分别处理，减少用户等待响应的时间；其次，单个重负载的运算分担到多台节点设备上做并行处理，每个节点设备处理结束后，将结果汇总，返回给用户，系统处理能力得到大幅度提高。
2）简单来说就是：其一是将大量的并发处理转发给后端多个节点处理，减少工作响应时间；其二是将单个繁重的工作转发给后端多个节点处理，处理完再返回给负载均衡中心，再返回给用户。目前负载均衡技术大多数是用于提高诸如在Web服务器、FTP服务器和其它关键任务服务器上的Internet服务器程序的可用性和可伸缩性。
二，负载均衡分类
1）二层负载均衡（mac）
根据OSI模型分的二层负载，一般是用虚拟mac地址方式，外部对虚拟MAC地址请求，负载均衡接收后分配后端实际的MAC地址响应）
2）三层负载均衡（ip）
一般采用虚拟IP地址方式，外部对虚拟的ip地址请求，负载均衡接收后分配后端实际的IP地址响应）
3）四层负载均衡（tcp）
在三次负载均衡的基础上，用ip+port接收请求，再转发到对应的机器。
4）七层负载均衡（http）
根据虚拟的url或IP，主机名接收请求，再转向相应的处理服务器）。

我们运维中最常见的四层和七层负载均衡，这里重点说下这两种负载均衡。
1）四层的负载均衡就是基于IP+端口的负载均衡：在三层负载均衡的基础上，通过发布三层的IP地址（VIP），然后加四层的端口号，来决定哪些流量需要做负载均衡，对需要处理的流量进行NAT处理，转发至后台服务器，并记录下这个TCP或者UDP的流量是由哪台服务器处理的，后续这个连接的所有流量都同样转发到同一台服务器处理。
对应的负载均衡器称为四层交换机（L4 switch），主要分析IP层及TCP/UDP层，实现四层负载均衡。此种负载均衡器不理解应用协议（如HTTP/FTP/MySQL等等）。
实现四层负载均衡的软件有：
F5：硬件负载均衡器，功能很好，但是成本很高。
lvs：重量级的四层负载软件
nginx：轻量级的四层负载软件，带缓存功能，正则表达式较灵活
haproxy：模拟四层转发，较灵活
2）七层的负载均衡就是基于虚拟的URL或主机IP的负载均衡：在四层负载均衡的基础上（没有四层是绝对不可能有七层的），再考虑应用层的特征，比如同一个Web服务器的负载均衡，除了根据VIP加80端口辨别是否需要处理的流量，还可根据七层的URL、浏览器类别、语言来决定是否要进行负载均衡。举个例子，如果你的Web服务器分成两组，一组是中文语言的，一组是英文语言的，那么七层负载均衡就可以当用户来访问你的域名时，自动辨别用户语言，然后选择对应的语言服务器组进行负载均衡处理。
对应的负载均衡器称为七层交换机（L7 switch），除了支持四层负载均衡以外，还有分析应用层的信息，如HTTP协议URI或Cookie信息，实现七层负载均衡。此种负载均衡器能理解应用协议。
实现七层负载均衡的软件有：
haproxy：天生负载均衡技能，全面支持七层代理，会话保持，标记，路径转移；
nginx：只在http协议和mail协议上功能比较好，性能与haproxy差不多；
apache：功能较差
Mysql proxy：功能尚可。

总的来说，一般是lvs做4层负载；nginx做7层负载；haproxy比较灵活，4层和7层负载均衡都能做


nginx 负载均衡 

测试域名 www.nginx.cy
a服务器地址：172.16.0.10（主服务器 反向代理服务器）
b服务器地址：172.16.0.100
c服务器地址：172.16.0.200

a服务器的配置(nginx.conf)，在http段加入以下代码，配置服务器集群server cluster，weight代表权重
upstream www.nginx.cy{
	server 172.16.0.100:80 weight=2;#有2/5的机率访问打这台服务器
	server 172.16.0.200:80 weight=3;
}
server{
	listen 80;
	server_name www.nginx.cy;
	location /{
		#反向代理地址
		proxy_pass http://www.nginx.cy;
		#设置主机头和客户端的真实地址，以便服务器获取客户端真实IP
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	}
}
b,c 服务器正常配置即可

以上测试是通过权重 进行负载均衡,也可以按照轮询,ip哈希,url哈希等多种方式对服务器做负载均衡
nginx 的 upstream目前支持 4 种方式的分配 
1)、轮询（默认） 
      每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器down掉，能自动剔除。 
2)、weight 
      指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。 
2)、ip_hash 
      每个请求按访问ip的hash结果分配，这样每个访客固定访问一个后端服务器，可以解决session的问题。  
3)、fair（第三方） 
      按后端服务器的响应时间来分配请求，响应时间短的优先分配。  
4)、url_hash（第三方）

配置负载均衡比较简单,但是最关键的一个问题是怎么实现多台服务器之间session的共享
下面有几种方法(以下内容来源于网络,第四种方法没有实践.)

1) 不使用session，换作cookie

能把session改成cookie，就能避开session的一些弊端，在从前看的一本J2EE的书上，也指明在集群系统中不能用session，否则惹出祸端来就不好办。如果系统不复杂，就优先考虑能否将session去掉，改动起来非常麻烦的话，再用下面的办法。

2) 应用服务器自行实现共享

asp.net可以用数据库或memcached来保存session，从而在asp.net本身建立了一个session集群，用这样的方式可以令 session保证稳定，即使某个节点有故障，session也不会丢失，适用于较为严格但请求量不高的场合。但是它的效率是不会很高的，不适用于对效率 要求高的场合。

以上两个办法都跟nginx没什么关系，下面来说说用nginx该如何处理：

3) ip_hash

nginx中的ip_hash技术能够将某个ip的请求定向到同一台后端，这样一来这个ip下的某个客户端和某个后端就能建立起稳固的session，ip_hash是在upstream配置中定义的：

upstream backend {
  server 127.0.0.1:8080 ;
  server 127.0.0.1:9090 ;
   ip_hash;
}

ip_hash是容易理解的，但是因为仅仅能用ip这个因子来分配后端，因此ip_hash是有缺陷的，不能在一些情况下使用：

1/ nginx不是最前端的服务器。ip_hash要求nginx一定是最前端的服务器，否则nginx得不到正确ip，就不能根据ip作hash。譬如使用的是squid为最前端，那么nginx取ip时只能得到squid的服务器ip地址，用这个地址来作分流是肯定错乱的。

2/ nginx的后端还有其它方式的负载均衡。假如nginx后端又有其它负载均衡，将请求又通过另外的方式分流了，那么某个客户端的请求肯定不能定位到同一台session应用服务器上。这么算起来，nginx后端只能直接指向应用服务器，或者再搭一个squid，然后指向应用服务器。最好的办法是用location作一次分流，将需要session的部分请求通过ip_hash分流，剩下的走其它后端去。

4) upstream_hash

为了解决ip_hash的一些问题，可以使用upstream_hash这个第三方模块，这个模块多数情况下是用作url_hash的，但是并不妨碍将它用来做session共享：

假如前端是squid，他会将ip加入x_forwarded_for这个http_header里，用upstream_hash可以用这个头做因子，将请求定向到指定的后端：

可见这篇文档：http://www.sudone.com/nginx/nginx_url_hash.html

在文档中是使用

hash   ;

这样就改成了利用x_forwarded_for这个头作因子，在nginx新版本中可支持读取cookie值，所以也可以改成：

hash   ;

nginx的upstream目前支持的5种方式的分配
1、轮询（默认）
每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器down掉，能自动剔除。
upstream backserver {
server 192.168.0.14;
server 192.168.0.15;
}
2、weight
指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。
upstream backserver {
server 192.168.0.14 weight=10;
server 192.168.0.15 weight=10;
}
3、ip_hash
每个请求按访问ip的hash结果分配，这样每个访客固定访问一个后端服务器，可以解决session的问题。
upstream backserver {
ip_hash;
server 192.168.0.14:88;
server 192.168.0.15:80;
}
4、fair（第三方）
按后端服务器的响应时间来分配请求，响应时间短的优先分配。
upstream backserver {
server server1;
server server2;
fair;
}
5、url_hash（第三方）
按访问url的hash结果来分配请求，使每个url定向到同一个后端服务器，后端服务器为缓存时比较有效。
upstream backserver {
server squid1:3128;
server squid2:3128;
hash $request_uri;
hash_method crc32;
}