/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 21:24:00
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 21:27:48
 */
package category

import (
	"weego/pkg/app"
	"weego/pkg/database"
	"weego/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (category Category) {
	database.DB.Where("id", idstr).First(&category)
	return
}

func GetBy(field, value string) (category Category) {
	database.DB.Where("? = ?", field, value).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
