package inventory

type (
	BookBulkReq struct {
		BookBulk []BookReq `json:"books,omitempty" swaggertype:"array,object,string"`
	}
	BookReq struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		PubId  int    `json:"pubId"`
	}
	BookDto struct {
		ID        int          `json:"id,omitempty"`
		Title     string       `json:"title"`
		Author    string       `json:"author"`
		PubId     int          `json:"-,omitempty"`
		Publisher PublisherDto `json:"publisher"`
	}
)
