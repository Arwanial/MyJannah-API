package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//Response
type Response struct {
	code int            `json:"code"`
	massage  string      `json:"massage"`
	Data    interface{}  `json:"data"`
}

//Response
type ErrResponse struct {
	Errcode int         `json:"code"`
	Errmsg  interface{} `json:"message"`
}

type ResponseWithoutData struct {
	code int            `json:"code"`
	massage  string     `json:"massage"`
}
