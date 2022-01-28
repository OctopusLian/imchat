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
	"fmt"
	"html/template"
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

func RegisterTemplate() {
	//全局扫描模板
	GlobTemplete := template.New("root")
	GlobTemplete, err := GlobTemplete.ParseGlob("view/**/*")
	if err != nil {
		//打印错误信息
		//退出系统
		log.Fatal(err)
	}
	//分别对每一个模板进行注册

	for _, tpl := range GlobTemplete.Templates() {
		patern := tpl.Name()
		http.HandleFunc(patern,
			func(w http.ResponseWriter,
				r *http.Request) {
				GlobTemplete.ExecuteTemplate(w, patern, nil)
			})
		fmt.Println("register=>" + patern)
	}
}

func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		//
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)
	//提供静态资源支持
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//提供指定目录的静态文件支持
	http.Handle("/asset", http.FileServer(http.Dir(".")))
	//user/login.html
	http.HandleFunc("/user/login.shtml", func(rw http.ResponseWriter, r *http.Request) {
		//解析
		tpl, err := template.ParseFiles("view/user/login.html")
		if err != nil {
			//打印并直接突出
			log.Fatal(err.Error())
		}
		tpl.ExecuteTemplate(rw, "/user/login.shtml", nil)
	})

	RegisterView()
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
