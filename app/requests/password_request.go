/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-12 01:46:35
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-12 02:18:54
 */
package requests

import (
	"weego/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

// ResetByPhone 验证表单，返回长度等于零即可通过
func ResetByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone":       []string{"required:手机号不能为空", "digits:手机号必须是11位数字"},
		"verify_code": []string{"required:验证码不能为空", "digits:验证码必须是6位数字"},
		"password":    []string{"required:密码不能为空", "digits:密码大于 6位数字"},
	}

	errs := validate(data, rules, messages)

	// 检查密码
	_data := data.(*ResetByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
