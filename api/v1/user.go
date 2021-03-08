package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study07/dto"
	"study07/middleware"
	"study07/model"
	"study07/utils/errmsg"
	"study07/utils/validator"
)

var code int

// 注册
func Register(c *gin.Context) {
	var msg string
	var user model.User
	_ = c.ShouldBindJSON(&user)

	msg, code = validator.Validate(&user)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg,
		})
		c.Abort()
		return
	}

	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&user)
	}
	if code == errmsg.USER_EXIST {
		code = errmsg.USER_EXIST
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 登录
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	var token string
	var code int

	user, code = model.CheckLogin(user.Username, user.Password)

	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(user)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

// 获取用户信息

func GetInfo(c *gin.Context) {
	user, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.SUCCSE,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}
