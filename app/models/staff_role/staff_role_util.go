package staff_role

import (
    "weego/pkg/app"
    "weego/pkg/database"
    "weego/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (staffRole StaffRole) {
    database.DB.Where("id",idstr).First(&staffRole)
    return
}

func GetBy(field, value string) (staffRole StaffRole) {
    database.DB.Where("? = ?", field, value).First(&staffRole)
    return
}

func All() (staffRoles [] StaffRole) {
    database.DB.Find(&staffRoles)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(StaffRole{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (staffRoles []StaffRole, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(StaffRole{}),
        &staffRoles,
        app.V1URL(database.TableName(&StaffRole{})),
        perPage,
    )
    return
}