/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-01 13:15:24
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-01 13:15:25
 */
package mail

type Driver interface {

	// 检查验证码
	Send(email Email, config map[string]string) bool
}
