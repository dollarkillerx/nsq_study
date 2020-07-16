# Nsq Study  nsq.io
NSQ是实时的分布式消息处理平台，其设计的目的是用来大规模地处理每天数以十亿计级别的消息。
### 优点 
- 原生分布式
- 支持无缝的水平扩展
- 支持数据本地化 (先向缓存里面插入,到了阀值往硬盘里面插入)
- TLS

### install
docker-compose
```
version: '3'
services:
  nsqlookupd:      # 有点像zookeeper 
    image: nsqio/nsq
    command: /nsqlookupd
    ports:
      - "4160:4160"   # 用来nsqd广播
      - "4161:4161"   # 用于客户端发现和管理
  nsqd:
    image: nsqio/nsq   # nsq core
    command: /nsqd --lookupd-tcp-address=nsqlookupd:4160
    depends_on:
      - nsqlookupd
    ports:
      - "4150:4150"         # TCP： 4150 给客户端使用
      - "4151:4151"         # HTTP 4151 HTTP API  
  nsqadmin:         # 可视化管理界面
    image: nsqio/nsq
    command: /nsqadmin --lookupd-http-address=nsqlookupd:4161
    depends_on:
      - nsqlookupd
    ports:
      - "4171:4171"      #  web
```

### 服务解析
- `nsqlookupd`
    - 基础概念 (服务发现与注册)
        - 管理拓扑信息,提供目录服务
        - nsqd节点通过nsqlookupd广播话题和通道信息
        - 客户端查询nsqlookupd来发现指定话题的生产者nsqd节点
        - 由此nsqlookupd将消费者与生产者解耦开。
          nsqlookupd当前的实现是简单地返回所有地址。
    - 端口
        - TCP: 4160 用来nsqd广播
        - HTTP 4161 用于客户端发现和管理
- `nsqd`
    - 基础概念
        - 负责接收、排队、投递消息给客户端
        - 每个nsqd有一个与nsqlookupd的长期TCP连接，定期推动其状态（sending heartbeat）
        - 对于消费者来说，一个暴露的HTTP /lookup接口用于轮询。
        - nsqd会发送心跳包给客户端，连续两个没有应答则超时关闭连接。
    - 端口
        - TCP： 4150 给客户端使用
        - HTTP 4151 HTTP API  

### 基础概念
- 话题
- channel
    - 对订阅了同一个topic，同一个channel的消费者使用负载均衡策略（不是轮询）
    - 只要channel存在，即使没有该channel的消费者，也会将生产者的message缓存到队列中（注意消息的过期处理）
- 消息

![](https://f.cloud.github.com/assets/187441/1700696/f1434dc8-6029-11e3-8a66-18ca4ea10aca.gif)

这图很明确的说明了
- 当同一个话题使用同一个channels   消息竞争(消息只会分配给一个消费者)
- 其他channels当没有消费者是  消息会堆积到channels里面  (FBI WARNING警告：topic，channel一旦建立，将会一直存在，要及时在管理台或者用代码清除无效的topic和channel，避免资源的浪费)
- 当同一个话题被不同的channels所消费 不存在消息竞争

### 连接方式
- 生产者
    - 生产者必须直连nsqd去投递message   (傻逼nsq)
- 消费者
    1. 直接连接nsqd
    2. 通过 nsqlookupd 查询 当前消息所在的 nsqd
#### 注意事项:
- Producer短线后不会重连,Consumer断线后会自动重连
  

      
      
