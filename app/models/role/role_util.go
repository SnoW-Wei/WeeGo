/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:25:09
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:32:05
 */
package role

import (
	"weego/pkg/app"
	"weego/pkg/database"
	"weego/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (roleResource RoleResource) {
	database.DB.Where("id", idstr).First(&roleResource)
	return
}

func GetBy(field, value string) (roleResource RoleResource) {
	database.DB.Where("? = ?", field, value).First(&roleResource)
	return
}

func All() (roleResources []RoleResource) {
	database.DB.Find(&roleResources)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(RoleResource{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (roleResources []RoleResource, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(RoleResource{}),
		&roleResources,
		app.V1URL(database.TableName(&RoleResource{})),
		perPage,
	)
	return
}
