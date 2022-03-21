package staff

import (
    "weego/pkg/app"
    "weego/pkg/database"
    "weego/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (staff Staff) {
    database.DB.Where("id",idstr).First(&staff)
    return
}

func GetBy(field, value string) (staff Staff) {
    database.DB.Where("? = ?", field, value).First(&staff)
    return
}

func All() (staff [] Staff) {
    database.DB.Find(&staff)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Staff{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (staff []Staff, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Staff{}),
        &staff,
        app.V1URL(database.TableName(&Staff{})),
        perPage,
    )
    return
}