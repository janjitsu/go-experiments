package albums

import "errors"

type Repo struct {
	Albums []Album
}

func NewInMemoryRepo() *Repo {
	return &Repo{
		Albums: []Album{},
	}
}

func (r *Repo) GetAll() []Album {
	return r.Albums
}

func (r *Repo) Add(newAlbum Album) {
	r.Albums = append(r.Albums, newAlbum)
}

func (r *Repo) GetById(id string) (*Album, error) {
	for _, a := range r.Albums {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Album not found")
}

func (r *Repo) RemoveById(id string) error {
	for i, a := range r.Albums {
		if a.ID == id {
			r.Albums = append(r.Albums[:i], r.Albums[i+1:len(r.Albums)]...)
			return nil
		}
	}

	return errors.New("Album not found")
}
