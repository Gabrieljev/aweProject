package member

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Username  string
		Password  string
		IsDeleted int
	}
)

func (User) TableName() string {
	return "user"
}
