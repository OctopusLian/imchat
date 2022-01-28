/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 22:21:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-28 20:06:11
 */
package main

import (
	"net/http"
)

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter, request *http.Request) {
			//数据库操作
			//逻辑处理
			//restapi json/xml返回
			//1，获取前端传递的参数
			//mobile,passwd
			//解析参数
			//如何获得参数
			//解析参数
			request.ParseForm()
			mobile := request.PostForm.Get("mobile")
			passwd := request.PostForm.Get("passwd")

			loginok := false
			if mobile == "18600000000" && passwd == "123456" {
				loginok = true
			}

			//如何返回json
			str := `{"code":0,"msg":"{"id":1,"token":"test"}}"`
			if !loginok {
				str = `{"code":0,"data":"密码不正确"`
			}
			//设置header为JSON默认的text/html
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(str))

		})
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
