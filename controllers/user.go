package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vuuvv/orca/server"
	"io/ioutil"
	"net/http"
)

type userController struct {
	server.BaseController
}

var UserController = &userController{}

func (this *userController) Name() string {
	return "用户管理"
}

func (this *userController) Path() string {
	return "user"
}

func (this *userController) Middlewares() []gin.HandlerFunc {
	return nil
}

func (this *userController) Mount(router *gin.RouterGroup) {
	this.Get("", this.index).WithName("查询")
	this.Post("", this.create).WithName("新增")
	this.Put("", this.edit).WithName("修改")
	this.Delete("", this.delete).WithName("删除")
}

func (this *userController) index(ctx *gin.Context) {
	this.Send("index")
}

func (this *userController) create(ctx *gin.Context) {
	this.Send("create")
}

func (this *userController) edit(ctx *gin.Context) {
	this.Send("edit")
}

func (this *userController) delete(ctx *gin.Context) {
	resp, _ := http.Get("https://www.taobao.com")
	bs, _ := ioutil.ReadAll(resp.Body)
    this.Send(string(bs))
}
