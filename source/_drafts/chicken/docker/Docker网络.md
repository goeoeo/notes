### Docker 网络
* docker network inspect bridge 查看桥接网络模式的连接情况
* sudo docker run -d --name test2 busybox /bin/sh -c "while true; do sleep 3600; done" 实验命令
* 默认桥接网络中的容器只能通过IP地址访问其他容器（除非使用遗留的-link指令连接两个容器），而自定义桥接网络提供DNS解析，可以通过容器的名字或是别名访问其他容器
* overlay网络和etcd实现多机容器通信

### 网络命名空间(linux虚拟网络)
* ip netns list 查看虚拟网络列表
* ip netns delete test1 删除虚拟网络test1
* ip netns add test1 添加虚拟网络test1
 * ip netns exec test1 ip link set dev lo up 启动test的网卡lo
* ip link add veth-test1 type veth peer name veth-test2 添加虚拟网络设备对
* ip link set veth-test1 netns test1 添加veth-test1到test1