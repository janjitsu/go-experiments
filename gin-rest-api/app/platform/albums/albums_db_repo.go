package albums

import "errors"

type DbRepo struct {
	Conn string
}

func NewDbRepo(ConnectionString string) *DbRepo {
	return &DbRepo{
		Conn: ConnectionString
	}
}

func (r *DbRepo) GetAll() []Album {
	return r.Albums
}

func (r *DbRepo) Add(newAlbum Album) {
	r.Albums = append(r.Albums, newAlbum)
}

func (r *DbRepo) GetById(id string) (*Album, error) {
	for _, a := range r.Albums {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Album not found")
}

func (r *DbRepo) RemoveById(id string) error {
	for i, a := range r.Albums {
		if a.ID == id {
			r.Albums = append(r.Albums[:i], r.Albums[i+1:len(r.Albums)]...)
			return nil
		}
	}

	return errors.New("Album not found")
}
