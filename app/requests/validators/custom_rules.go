/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-04 11:04:42
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-04 11:58:25
 */
package validators

import (
	"errors"
	"fmt"
	"strings"
	"weego/pkg/database"

	"github.com/thedevsaddam/govalidator"
)

// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {

	// 自定义规则 not_exists，验证请求数据必须不存在与数据库中
	// 常用于保证数据库某个字段的值唯一，如用户名，邮箱，手机号，或者分类的名称
	// not_exists参数可以有两种，一种是 2 个参数，一种是 3个参数
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉id 为 32的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := rng[0]

		dbFiled := rng[1]

		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		requestValue := value.(string)

		query := database.DB.Table(tableName).Where(dbFiled+"= ?", requestValue)

		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		var count int64
		query.Count(&count)

		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})
}
