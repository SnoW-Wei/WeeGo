/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-05 14:04:18
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:14:33
 */
package auth

import (
	"errors"
	"weego/app/models/user"
	"weego/pkg/logger"

	"github.com/gin-gonic/gin"
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

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户id
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}

// CurrentUID 从 gin.context 中获取当前登录用户名
func CurrentUserName(c *gin.Context) string {
	return c.GetString("current_user_name")
}
