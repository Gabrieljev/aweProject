package inventory

import "gorm.io/gorm"

type (
	Book struct {
		gorm.Model
		Title     string
		Author    string
		PubId     int       `gorm:"column:pub_id"`
		Publisher Publisher `gorm:"foreignKey:pub_id"`
		IsDeleted int
	}
)

func (Book) TableName() string {
	return "book"
}
