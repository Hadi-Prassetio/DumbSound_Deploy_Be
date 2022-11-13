package models

type Song struct {
	ID       int           `json:"id"  gorm:"primary_key:auto_increment"`
	Title    string        `json:"title"`
	Image    string        `json:"image"`
	Year     int           `json:"year"`
	Song     string        `json:"song"`
	ArtistID int           `json:"-"`
	Artist   ArtistProfile `json:"artist"`
}
