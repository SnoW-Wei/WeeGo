/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-05 14:04:18
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-05 14:10:14
 */
package auth

import (
	"errors"
	"weego/app/models/user"
)

//Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {

	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {

	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	return userModel, nil
}
