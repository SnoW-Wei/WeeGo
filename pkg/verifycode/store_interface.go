/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-28 18:42:46
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-28 18:43:47
 */
package verifycode

type Store interface {
	// 保存验证码
	Set(id string, value string) bool

	// 获取验证码
	Get(id string, clear bool) string

	// 检查验证码
	Verify(id, answer string, clear bool) bool
}
