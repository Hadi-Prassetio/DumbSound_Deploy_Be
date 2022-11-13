package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type SongRepository interface {
	FindSong() ([]models.Song, error)
	GetSong(ID int) (models.Song, error)
	CreateSong(song models.Song) (models.Song, error)
}

func RepositorySong(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSong() ([]models.Song, error) {
	var songs []models.Song
	err := r.db.Preload("Artist").Find(&songs).Error

	return songs, err
}

func (r *repository) GetSong(ID int) (models.Song, error) {
	var song models.Song
	err := r.db.Preload("Artist").First(&song, ID).Error

	return song, err
}

func (r *repository) CreateSong(song models.Song) (models.Song, error) {
	err := r.db.Create(&song).Error

	return song, err
}
