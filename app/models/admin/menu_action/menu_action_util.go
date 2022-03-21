package menu_action

import (
    "weego/pkg/app"
    "weego/pkg/database"
    "weego/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (menuAction MenuAction) {
    database.DB.Where("id",idstr).First(&menuAction)
    return
}

func GetBy(field, value string) (menuAction MenuAction) {
    database.DB.Where("? = ?", field, value).First(&menuAction)
    return
}

func All() (menuActions [] MenuAction) {
    database.DB.Find(&menuActions)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(MenuAction{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (menuActions []MenuAction, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(MenuAction{}),
        &menuActions,
        app.V1URL(database.TableName(&MenuAction{})),
        perPage,
    )
    return
}