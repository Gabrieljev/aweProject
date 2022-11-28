package inventory

type (
	Publisher struct {
		ID        int `gorm:"primaryKey"`
		Name      string
		CountryId int `gorm:"column:country_id"`
	}
)

func (Publisher) TableName() string {
	return "publisher"
}
