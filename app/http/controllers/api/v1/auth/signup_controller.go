package auth

import (
	"fmt"
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

	// 解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {

		// 解析失败，返回422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		// 打印错误信息
		fmt.Println(err.Error())
		// 出错，中断请求
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零通过，大于0 即有错误发生
	if len(errs) >= 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检验数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exit": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}

	// 解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {

		// 解析失败，返回422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		// 打印错误信息
		fmt.Println(err.Error())
		// 出错，中断请求
		return
	}

	errs := requests.ValidateSignipEmailExist(&request, c) //

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
