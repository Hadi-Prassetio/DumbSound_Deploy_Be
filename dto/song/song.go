package songdto

type RequestSong struct {
	Title    string `jsom:"title" validate:"required"`
	Image    string `json:"image"`
	Year     int    `json:"year" validate:"required"`
	Song     string `json:"song"`
	ArtistID int    `json:"artist_id"`
}
