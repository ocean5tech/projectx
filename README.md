# projectx 

# EP1 - EP2

# PACKAGE MAIN

main.go
1. 初期化Localtransport
2. 连接， 为啥非得22连接？
3. trRemort持续sendmessage到trlocal
4. 初期化Server，trlocal作为输入参数
5. 实例化server
6. server.start

# PACKAGE NETWORK - 网络层，与业务无关

transport.go

1. module部分，负载结构体RPC， Transport公用接口

local_transport.go

1. 声明了个本地Node
  - 本机地址 addr
  - 接收消息通道 consumeCh
  - 锁 lock
  - 交互Node对，每台其他客户端都是22交互，而不是通过server交互？ peers
2. 实现了Transport公用接口
 - Consume 返回通道，但是没有处理消息
 - Connect 设置的交互地址对
 - SendMessage 取得交互Node对对方的consumeCh，把消息体RPC送入
 - Addr 返回本Node地址

local_transport_test.go

1. TestConnect
 - 实例化两个Node
 - 互相连接两个Node
 - 确认连接成功

2. TestSendMessage - 两个Node直接发送Message，不需要Server！！
 - 实例化两个Node
 - 互相连接两个Node
 - NodeA发送消息给Nodeb
 - 取得Nodeb的通道值，跟NodeA的发送值比较，证明成功

server.go

1. Server结构
 - 一个注册到server里的Node列表
 - 一个消息通道
 - 一个退出通道

2. Start
 - 初始化通道，遍历每个Node，取出所有在该Node的接收消息通道里的消息负载，并传送给Server自身的通道
 - 启动一个计时器，每5秒一个提醒
 - 开始监控无限循环
   - 取出Server的消息通道里的值，打印出来
   - 如果有退出通道的消息就退出无限循环
   - 监视计时器提醒通道，不让程序deadlock

# Issues
1. initTransports取了一次所有Node的消息， start以后怎么办？
  -> 现在来看，initTransports时，用了go func(tr Transport)，所以每个Node在Server上都有独立的goroutine，一直在循环，等待Consume里面的消息，一来就发给
  server的消息通道

2. Server跟Node都是各自独立？
  -> 只要在开始的Node列表中写入的Node，都会在Server里面有个单独的goroutine监控給该Node的所有Message，收到了就由Server处理
2-1. Node对中一个在server注册，一个没有，通信会如何？
  -> Node对可以继续通信，只有在Server注册的Node收到的消息，会被Server处理
2-2. Server中注册的Node，消息会被Server截获吗？ 别的Node来的消息收得到？
  -> 会被截获，但是消息也正常收到
3-3. 单向连接又如何？
  -> 单项连接的正常处理，反向不处理，也不出错？？？

3. Node在Sendmessage时是直接send给了对方Node的接受通道，Server感知得到吗？
 -> 能，Server和对方Node一起收到消息

4. LocalTransport不会自己接受消息，只会把消息通道共有给别人，让调用者Main，Test，server来取得值？
 -> 对

