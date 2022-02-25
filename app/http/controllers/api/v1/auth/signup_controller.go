/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-22 18:18:22
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-25 14:45:11
 */
package auth

import (
	"net/http"
	v1 "weego/app/http/controllers/api/v1"
	"weego/app/models/user"

	"weego/app/http/requests"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// 检验数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exit": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignipEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
