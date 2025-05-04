分布式模型

Hub&Spoke
一对多
所有服务都通过一个中心服务进行通信,
利于负载均衡，
存在单点故障，中心服务处理各种请求

Peer to Peer
多对多
没有单点故障，解耦程度高
服务很难被发现，负载均衡困难

Message Queue
消息队列
利于扩展，利于消息的持久化
存在单点故障，配置复杂

混合模型：
                    <->  server
                        /       \
client <-> netmask  <-> hub -> cloud
                        \       /
                    <->  server

优点：
有利于负载均衡
对服务失败的防范更加健壮

缺点：
架构更加复杂
这个Hub的作用范围难以界定

服务注册：服务注册，健康检查
用户门户：web应用，api网关
日志服务：集中式日志中心
业务服务：业务逻辑，数据持久化

技术选型： go，http，json


创建web服务
创建注册服务
注册web服务
2.1 创建日志服务
2.2 独立的日志服务

log   service  registry
 |      /     
 ([logservice]    cmd )