package inventory

type (
	PublisherDto struct {
		ID        int    `json:"id,omitempty"`
		Name      string `json:"name"`
		CountryId int    `json:"countryId"`
	}
)
