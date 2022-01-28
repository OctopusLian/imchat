/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 22:21:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-28 20:40:02
 */
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func userLogin(writer http.ResponseWriter, request *http.Request) {
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
	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "")
	} else {
		Resp(writer, -1, nil, "密码不正确")
	}
}

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` //omitempty作用：如果为空就不显示给前端
}

func Resp(w http.ResponseWriter, code int, date interface{}, msg string) {
	//设置header为JSON默认的text/html
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: date,
	}
	//将结构体转化为json字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)
	//提供静态资源支持
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//提供指定目录的静态文件支持
	http.Handle("/asset", http.FileServer(http.Dir(".")))
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
