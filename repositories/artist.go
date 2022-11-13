package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type ArtistRepository interface {
	FindArtist() ([]models.Artist, error)
	GetArtist(ID int) (models.Artist, error)
	CreateArtist(artist models.Artist) (models.Artist, error)
}

func RepositoryArtist(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArtist() ([]models.Artist, error) {
	var artists []models.Artist
	err := r.db.Preload("Songs").Find(&artists).Error
	return artists, err
}

func (r *repository) GetArtist(ID int) (models.Artist, error) {
	var artist models.Artist
	err := r.db.Preload("Songs").First(&artist).Error
	return artist, err
}

func (r *repository) CreateArtist(artist models.Artist) (models.Artist, error) {
	err := r.db.Create(&artist).Error

	return artist, err
}
