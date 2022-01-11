package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Fortunate", Artist: "Gojira", Price: 99.99},
	{ID: "2", Title: "Amazonia", Artist: "Gojira", Price: 99.99},
	{ID: "3", Title: "Hold On", Artist: "Gojira", Price: 99.99},
	{ID: "4", Title: "New Found", Artist: "Gojira", Price: 99.99},
}
