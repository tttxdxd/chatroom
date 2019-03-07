# chatroom
# 基于Go的多人聊天室

1. 自定义消息协议

![架构图](https://raw.githubusercontent.com/tttxdxd/chatroom/master/README/chatroom.png)

## 服务端

1. 建立与客户端的连接
2. 消息分发到对应处理器
3. 处理器处理信息返回response到客户端

## 客户端

1. 选择客户端功能
2. 建立与服务器的连接
3. 发送消息到服务器，并处理回应
