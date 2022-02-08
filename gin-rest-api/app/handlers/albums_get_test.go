package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/janjitsu/go-experiments/gin-rest-api/app/platform/albums"
)

func TestGetEmptyAlbums(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	albumRepo := albums.NewInMemoryRepo()

	r.GET("/albums", GetAlbums(albumRepo))

	req, _ := http.NewRequest("GET", "/albums", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Should return status OK")
	}

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	albums := []albums.Album{}
	json.Unmarshal([]byte(body), &albums)

	if len(albums) > 0 {
		t.Errorf("Should return empty albums list, returned %+v", albums)
	}
}

func TestGetAlbums(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	albumRepo := albums.NewInMemoryRepo()

	albumRepo.Add(albums.Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  99.99,
	})

	r.GET("/albums", GetAlbums(albumRepo))

	req, _ := http.NewRequest("GET", "/albums", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Should return status OK")
	}

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	albums := []albums.Album{}
	json.Unmarshal([]byte(body), &albums)

	if len(albums) != 1 {
		t.Errorf("Should return albums list, returned %+v", albums)
	}
}
