package shopping

type (
	Book struct {
		ID        int `gorm:"primaryKey"`
		Title     string
		Author    string
		PubId     int       `gorm:"column:pub_id"`
		Publisher Publisher `gorm:"foreignKey:pub_id"`
	}
)

func (Book) TableName() string {
	return "book"
}
