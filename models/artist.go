package models

type Artist struct {
	ID          int    `json:"id"  gorm:"primary_key:auto_increment"`
	Name        string `json:"name"`
	Old         int    `json:"old"`
	Role        string `json:"role"`
	StartCareer int    `json:"start_career"`
	Songs       []Song `json:"song"`
}

type ArtistProfile struct {
	ID          int    `json:"id"  gorm:"primary_key:auto_increment"`
	Name        string `json:"name"`
	Old         int    `json:"old"`
	Role        string `json:"role"`
	StartCareer int    `json:"start_career"`
}

func (ArtistProfile) TableName() string {
	return "artists"
}
