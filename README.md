<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2021-11-05 13:28:23
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-29 22:29:50
-->
# 支持10万人同时在线Go语言打造高并发web即时聊天(IM)应用  

## 需求分析及拆分  

### 基本需求  

- 发送/接收  
- 实现群聊  
- 高并发=单机最好+分布式+弹性扩容  

### 需求拆分  

- 1，实现功能界面  
- 2，实现资源标准化编码  
- 3，确保消息体的可扩展性  
- 4，接收消息并解析显示  
- 5，群聊的特殊需求  
- 6，高并发  

## IM系统一般架构  

![](./res/IM系统一般架构.png)  

## 如何运行  

```
$ go run main.go
```

浏览器访问`http://localhost:8080/user/login.shtml`  