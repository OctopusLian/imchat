/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 22:47:41
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 22:47:42
 */
package ctrl

import "net/http"

func UserLogin(writer http.ResponseWriter,
	request *http.Request) {
	//数据库操作
	//逻辑处理
	//restapi json/xml返回
	//1.获取前端传递的参数
	//mobile,passwd
	//解析参数
	//如何获得参数
	//解析参数
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	//模拟
	user, err := userService.Login(mobile, passwd)

	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}

}
