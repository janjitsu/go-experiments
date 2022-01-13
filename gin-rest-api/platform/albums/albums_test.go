package albums

import "testing"

func TestAdd(t *testing.T) {
	repo := New()

	newAlbum := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	}

	repo.Add(newAlbum)

	if len(repo.Albums) != 1 {
		t.Errorf("Should Add Album to Repo")
	}
}

func TestRemove(t *testing.T) {
	repo := New()

	newAlbum := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	}

	repo.Add(newAlbum)
	repo.RemoveById("1")

	if len(repo.Albums) == 1 {
		t.Errorf("Should Add Album to Repo")
	}
}

func TestRemoveError(t *testing.T) {
	repo := New()

	newAlbum := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	}

	repo.Add(newAlbum)

	err := repo.RemoveById("2")

	if err == nil {
		t.Errorf("Should Throw Error when trying to remove album with invalid ID")
	}
}

func TestGetAll(t *testing.T) {
	repo := New()

	newAlbum := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	}

	repo.Add(newAlbum)
	repo.Add(newAlbum)

	if len(repo.GetAll()) != 2 {
		t.Errorf("Should Get All Albums on Repo")
	}
}

func TestGetById(t *testing.T) {
	repo := New()

	newAlbum := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	}

	repo.Add(newAlbum)

	albumRequest, _ := repo.GetById("1")

	if albumRequest == nil {
		t.Errorf("Should Get Album By Id on Repo")
	}
}

func TestGetByIdError(t *testing.T) {
	repo := New()

	newAlbum := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	}

	repo.Add(newAlbum)

	_, err := repo.GetById("2")

	if err == nil {
		t.Errorf("Should Throw Error when query for invalid ID")
	}
}
