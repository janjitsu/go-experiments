package albums

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DbRepo struct {
	DB     *sql.DB
	Albums []Album
}

func NewDbRepo(ConnectionString string) *DbRepo {

	db, err := sql.Open("pgx", "postgresql://gouser:123456@localhost:5432/gin-rest-api")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	return &DbRepo{
		DB:     db,
		Albums: []Album{},
	}
}

func (r *DbRepo) GetAll() []Album {
	//select all from albums bind to list

	rows, err := r.DB.Query("SELECT * FROM albums")
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
		return nil
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var albums []Album

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return albums
		}
		albums = append(albums, alb)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("could not list albums: %v", err)
		return albums
	}
	return albums
}

func (r *DbRepo) Add(newAlbum Album) {
	sqlStatement := `
		INSERT INTO albums (id, title, artist, price)
		VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(sqlStatement, newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	if err != nil {
		log.Fatalf("could not insert album: %v", err)
	}
}

func (r *DbRepo) GetById(id string) (*Album, error) {

	row := r.DB.QueryRow("SELECT id, title, artist, price FROM albums where id = $1  LIMIT 1", id)

	var album Album

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		log.Fatalf("could not scan row: %v", err)
		return nil, errors.New("Album not found")
	}

	return &album, nil
}

func (r *DbRepo) RemoveById(id string) error {

	sqlStatement := "DELETE FROM albums where id = $1"
	_, err := r.DB.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("could not remove album: %v", err)
		return errors.New("Album could not be removed")
	}

	return nil
}
