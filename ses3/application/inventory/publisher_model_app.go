package inventory

import "gorm.io/gorm"

type (
	Publisher struct {
		gorm.Model
		Name      string
		CountryId int `gorm:"column:country_id"`
		IsDeleted int
	}
)

func (Publisher) TableName() string {
	return "publisher"
}
