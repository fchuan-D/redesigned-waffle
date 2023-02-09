package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`          // 错误代码
	StatusMsg  string      `json:"status_msg,omitempty"` // 消息提示
	Data       interface{} `json:"data,omitempty"`       // 数据内容
}

const (
	Ok                  = 201
	NotOk               = 405
	Unauthorized        = 401
	Forbidden           = 403
	InternalServerError = 500
)

const (
	OkMsg                  = "操作成功"
	NotOkMsg               = "操作失败"
	NotFindMsg             = "查询失败"
	UnauthorizedMsg        = "登录过期, 需要重新登录"
	LoginCheckErrorMsg     = "用户名或密码错误"
	ForbiddenMsg           = "无权访问该资源, 请联系网站管理员授权"
	InternalServerErrorMsg = "服务器内部错误"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: code,
		StatusMsg:  msg,
		Data:       data,
	})
}

func OkResult(c *gin.Context) {
	Result(SUCCESS, OkMsg, map[string]interface{}{}, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, OkMsg, data, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}

func FailResult(c *gin.Context) {
	Result(ERROR, NotOkMsg, map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func FailWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(ERROR, message, data, c)
}
