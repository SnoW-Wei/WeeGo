package menu_action_resource

import (
    "weego/pkg/app"
    "weego/pkg/database"
    "weego/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (menuActionResource MenuActionResource) {
    database.DB.Where("id",idstr).First(&menuActionResource)
    return
}

func GetBy(field, value string) (menuActionResource MenuActionResource) {
    database.DB.Where("? = ?", field, value).First(&menuActionResource)
    return
}

func All() (menuActionResources [] MenuActionResource) {
    database.DB.Find(&menuActionResources)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(MenuActionResource{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (menuActionResources []MenuActionResource, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(MenuActionResource{}),
        &menuActionResources,
        app.V1URL(database.TableName(&MenuActionResource{})),
        perPage,
    )
    return
}