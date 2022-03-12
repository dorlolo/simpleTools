# 目录说明

基于董哥的方法重新封装了一下，修改了一些有问题的地方。

```text
.
├── registerClient.go  将gin注册到consul服务
├── filterServer.go   客户端从consul查找服务
├── consulResolver.go 网上的人写的基于grpc协议的consul解析器，用来做服务发现的。已经没什么用了，留个纪念
└── registerServer.go  将Gprc Server注册到consul服务

```