/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-01 19:52:49
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-04 12:05:18
 */
package validators

import (
	"weego/pkg/captcha"
	"weego/pkg/verifycode"
)

// ValidatorCaptcha 检查验证码是否正确
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {

	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// ValidatePassConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password string, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

// ValidateVerifyCode 自定义规则，验证【手机/邮箱验证码】
func ValidateVerifyCode(key string, answer string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}
