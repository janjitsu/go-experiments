package albums

type Getter interface {
	GetAll() []Album
}

type GetterById interface {
	GetById(id string) (*Album, error)
}

type Adder interface {
	Add(newAlbum Album)
}

type Remover interface {
	RemoveById(id string) error
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
