package category

import (
	"weego/app/models"
	"weego/pkg/database"
)

type Category struct {
	models.BaseModel

	Name         string `json:"name,omitempty"`
	Descripttion string `json:"descripttion,omitempty"`

	models.CommonTimeStampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
