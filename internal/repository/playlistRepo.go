package repository

import (
	"gorm.io/gorm"
	"music-service-v2.0.0/internal/models"
	"music-service-v2.0.0/internal/rest/dto"
)

type PlaylistRepo interface {
	CreatePlaylist(playlist *models.Playlist) error
	GetAllPlaylist() ([]models.Playlist, error)
	GetPlaylistByID(id uint) (*models.Playlist, error)
	UpdatePlaylistByID(id uint, playlistForm dto.PlaylistForm) error
	DeletePlaylistByID(id uint) error
}

type PlaylistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: db,
	}
}

func (pr *PlaylistRepository) CreatePlaylist(playlist *models.Playlist) error {
	if err := pr.db.Create(&playlist).Error; err != nil {
		return err
	}

	return nil
}

func (pr *PlaylistRepository) GetAllPlaylist() ([]models.Playlist, error) {
	var playlists []models.Playlist
	if err := pr.db.Preload("Tracks").Find(&playlists).Error; err != nil {
		return nil, err
	}

	return playlists, nil
}

func (pr *PlaylistRepository) GetPlaylistByID(id uint) (*models.Playlist, error) {
	var playlist models.Playlist
	if err := pr.db.First(&playlist, id).Error; err != nil {
		return nil, err
	}

	return &playlist, nil
}

func (pr *PlaylistRepository) DeletePlaylistByID(id uint) error {
	var playlist models.Playlist
	if err := pr.db.First(&playlist, id).Error; err != nil {
		return err
	}

	if err := pr.db.Delete(&playlist).Error; err != nil {
		return err
	}

	return nil
}

func (pr *PlaylistRepository) UpdatePlaylistByID(id uint, playlistForm dto.PlaylistForm) error {
	var playlist models.Playlist
	if err := pr.db.First(&playlist, id).Error; err != nil {
		return err
	}

	playlist.Title = playlistForm.Title

	if err := pr.db.Save(&playlist).Error; err != nil {
		return err
	}

	return nil
}
