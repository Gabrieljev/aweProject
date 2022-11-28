package shopping

type (
	Publisher struct {
		ID        uint `gorm:"primaryKey"`
		Name      string
		CountryId uint `gorm:"column:country_id"`
	}
)

func (Publisher) TableName() string {
	return "publisher"
}
