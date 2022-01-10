/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 22:21:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 22:25:51
 */
package main

import (
	"io"
	"net/http"
)

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login",
		func(rw http.ResponseWriter, r *http.Request) {
			//数据库操作
			//逻辑处理
			//restapi json/xml返回
			io.WriteString(rw, "hello world")
		})
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
