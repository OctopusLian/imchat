package main

import (
	"net/http"

		"./ctrl"
	"log"
	"html/template"
	"fmt"
		    _ "github.com/go-sql-driver/mysql"
	)


func RegisterTemplate(){
	//全局扫描模板
	GlobTemplete := template.New("root")
	GlobTemplete ,err:=GlobTemplete.ParseGlob("view/**/*")
	if err!=nil {
		//打印错误信息
		//退出系统
		log.Fatal(err)
	}
	//分别对每一个模板进行注册

	for _,tpl := range  GlobTemplete.Templates(){
		patern := tpl.Name()
		http.HandleFunc(patern,
			func(w http.ResponseWriter,
			r *http.Request) {
			GlobTemplete.ExecuteTemplate(w,patern,nil)
		})
		fmt.Println("register=>"+patern)
	}
}
func RegisterView(){
	//一次解析出全部模板
	tpl,err := template.ParseGlob("view/**/*")
	if nil!=err{
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _,v := range tpl.Templates(){
		//
		tplname := v.Name();
		fmt.Println("HandleFunc     "+v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     "+v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w,tplname,nil)
			if err!=nil{
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	//http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/chat", ctrl.Chat)
	http.HandleFunc("/attach/upload", ctrl.Upload)
	//1 提供静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))

	//2 指定目录的静态文件
	http.Handle("/asset/",http.FileServer(http.Dir(".")))
	http.Handle("/mnt/",http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080",nil)
}