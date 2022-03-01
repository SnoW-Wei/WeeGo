/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-01 19:52:49
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-01 19:54:52
 */
package validators

import "weego/pkg/captcha"

func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {

	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}
