/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 22:21:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 22:41:46
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
			//1，获取前端传递的参数
			//mobile,passwd
			//解析参数
			//如何获得参数
			//解析参数
			request.ParseForm()
			mobile := request.PostForm().Get
			passwd

			loginok := false
			if mobile == "18600000000" && passwd == “123456” {
				lologinok = true
			}

			//如何返回json
			if loginok {
				//返回json ok
			} else {
				//返回json failed
			}
			//io.WriteString(rw, "hello world")
		})
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
