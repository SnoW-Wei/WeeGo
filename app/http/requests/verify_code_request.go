/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-28 19:28:50
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-28 19:59:12
 */
package requests

import (
	"weego/pkg/captcha"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required: 手机号为必填项，参数名称 phone",
			"digits: 手机号长度必须为 11位数字",
		},
		"captcha_id": []string{
			"required: 图片验证码的ID 为必填",
		},
		"captcha_answer": []string{
			"required: 图片验证码答案必填",
			"digits: 图片验证码长度必须为 6位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*VerifyCodePhoneRequest)

	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
