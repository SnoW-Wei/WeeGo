package menu

import (
    "weego/pkg/app"
    "weego/pkg/database"
    "weego/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (menu Menu) {
    database.DB.Where("id",idstr).First(&menu)
    return
}

func GetBy(field, value string) (menu Menu) {
    database.DB.Where("? = ?", field, value).First(&menu)
    return
}

func All() (menus [] Menu) {
    database.DB.Find(&menus)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Menu{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (menus []Menu, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Menu{}),
        &menus,
        app.V1URL(database.TableName(&Menu{})),
        perPage,
    )
    return
}