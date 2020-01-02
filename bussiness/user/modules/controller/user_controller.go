package controller

import(
	"github.com/gin-gonic/gin"
)

//注册
type RegisterParam struct{
	Phone string `json:"phone" binding:"required,phone"`
	Role int `json:"role" binding:"required"`
	Token string `json:"token" binding:"required"`
}
func(p RegisterParam) Message() map[string]map[string]string {
	return map[string]map[string]string{
		"phone":map[string]string{
			"required":"手机号不能为空",
			"phone":"手机号格式不正确",
		},
		"token":map[string]string{
			"required":"参数错误",
		},
	}
}

func Register(ctx *gin.Context) {
	
}