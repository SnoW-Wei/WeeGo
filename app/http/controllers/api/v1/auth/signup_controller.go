package auth

import (
	"fmt"
	"net/http"
	v1 "weego/app/http/controllers/api/v1"
	"weego/app/models/user"

	"weego/app/http/requests"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		fmt.Println(err.Error())
		return
	}

	errs := requests.ValidateSignupPhoneExist(&request, c)

	if len(errs) >= 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exit": user.IsPhoneExist(request.Phone),
	})

}
