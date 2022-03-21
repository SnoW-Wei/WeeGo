package role_menu

import (
    "weego/pkg/app"
    "weego/pkg/database"
    "weego/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (roleMenu RoleMenu) {
    database.DB.Where("id",idstr).First(&roleMenu)
    return
}

func GetBy(field, value string) (roleMenu RoleMenu) {
    database.DB.Where("? = ?", field, value).First(&roleMenu)
    return
}

func All() (roleMenus [] RoleMenu) {
    database.DB.Find(&roleMenus)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(RoleMenu{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (roleMenus []RoleMenu, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(RoleMenu{}),
        &roleMenus,
        app.V1URL(database.TableName(&RoleMenu{})),
        perPage,
    )
    return
}