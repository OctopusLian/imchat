package ctrl


import (
	"net/http"
	"../util"
	"../service"
	"../args"
	"../model"
	)

var contactService service.ContactService
func LoadFriend(w http.ResponseWriter, req *http.Request){
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req,&arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w,users,len(users))
}

func LoadCommunity(w http.ResponseWriter, req *http.Request){
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req,&arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w,comunitys,len(comunitys))
}
func JoinCommunity(w http.ResponseWriter, req *http.Request){
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req,&arg)
	err := contactService.JoinCommunity(arg.Userid,arg.Dstid);
	//todo 刷新用户的群组信息
	AddGroupId(arg.Userid,arg.Dstid)
	if err!=nil{
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,nil,"")
	}
}
func CreateCommunity(w http.ResponseWriter, req *http.Request){
	var arg model.Community
	//如果这个用的上,那么可以直接
	util.Bind(req,&arg)
	com,err := contactService.CreateCommunity(arg);
	if err!=nil{
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,com,"")
	}
}
//
func Addfriend(w http.ResponseWriter, req *http.Request) {
	//定义一个参数结构体
	/*request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	*/
	var arg args.ContactArg
	util.Bind(req,&arg)
	//调用service
	err := contactService.AddFriend(arg.Userid,arg.Dstid)
	//
	if err!=nil{
		util.RespFail(w,err.Error())
	}else{
		util.RespOk(w,nil,"好友添加成功")
	}
}