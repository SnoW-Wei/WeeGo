/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-28 17:54:59
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-28 18:04:06
 */
package sms

type Driver interface {
	// 发送短信
	Send(phone string, message Message, config map[string]string) bool
}
