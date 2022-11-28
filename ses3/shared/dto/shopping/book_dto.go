package shopping

type (
	BookDto struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
		PubId  int    `json:"-,omitempty"`
		Publisher PublisherDto `json:"publisher"`
	}
)
