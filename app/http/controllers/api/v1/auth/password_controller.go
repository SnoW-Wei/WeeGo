/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-12 02:00:59
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-12 02:15:04
 */
package auth

import (
	v1 "weego/app/http/controllers/api/v1"
	"weego/app/models/user"
	"weego/app/requests"
	"weego/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

//
func (pc *PasswordController) ResetByphone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	userModel := user.GetByPhone(request.Phone)

	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}

}
