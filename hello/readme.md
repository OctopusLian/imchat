###6.1 UDP协议实现分布式

####6.1.1 支持分布式
回顾单体应用
开启ws接收协程recvproc/ws发送协程sendproc
websocket收到消息->dispatch发送给dstid

基于UDP的分布式应用
开启ws接收协程recvproc/ws发送协程sendproc
开启udp接收协程udprecvproc/udp发送协程udpsendproc

websocket收到消息->broadMsg广播到局域网
udp接收到收到消息->dispatch发送给dstid
自己是局域网一份子,所以也能接收到消息

####6.1.2 实现

####6.1.3 nginx反向代理
```html
	upstream wsbackend {
			server 192.168.0.102:8080;
			server 192.168.0.100:8080;
			hash $request_uri;
	}
	map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
	}
    server {
	  listen  80;
	  server_name localhost;
	  location / {
	   proxy_pass http://wsbackend;
	  }
	  location ^~ /chat {
	   proxy_pass http://wsbackend;
	   proxy_connect_timeout 500s;
       proxy_read_timeout 500s;
	   proxy_send_timeout 500s;
	   proxy_set_header Upgrade $http_upgrade;
       proxy_set_header Connection "Upgrade";
	  }
	 }

}
```
####6.3 打包发布
#windows平台
```bash
::remove dir
rd /s/q release
::make dir 
md release
::go build -ldflags "-H windowsgui" -o chat.exe
go build -o chat.exe
::
COPY chat.exe release\
COPY favicon.ico release\favicon.ico
::
XCOPY asset\*.* release\asset\  /s /e
XCOPY view\*.* release\view\  /s /e 
```
#linux平台
```bash
#!/bin/sh
rm -rf ./release
mkdir  release
go build -o chat
chmod +x ./chat
cp chat ./release/
cp favicon.ico ./release/
cp -arf ./asset ./release/
cp -arf ./view ./release/
```
#运行注意事项
linux 下
```bash
nohup ./chat >>./log.log 2>&1 &
```
