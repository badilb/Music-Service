package repository

import (
	"gorm.io/gorm"
	"music-service-v2.0.0/internal/models"
)

type TrackRepo interface {
	CreateTrack(track *models.Tracks) error
	GetAllTracks() ([]models.Tracks, error)
	GetTracksByID(id uint) (*models.Tracks, error)
	UpdateTracksByID(id uint, trackRequest models.Tracks) error
	DeleteTracksByID(id uint) error
}

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

func (tr *TrackRepository) CreateTrack(track *models.Tracks) error {
	if err := tr.db.Create(&track).Error; err != nil {
		return err
	}

	return nil
}

func (tr *TrackRepository) GetAllTracks() ([]models.Tracks, error) {
	var tracks []models.Tracks
	if err := tr.db.Find(&tracks).Error; err != nil {
		return nil, err
	}

	return tracks, nil
}

func (tr *TrackRepository) GetTracksByID(id uint) (*models.Tracks, error) {
	var track models.Tracks
	if err := tr.db.First(&track, id).Error; err != nil {
		return nil, err
	}

	return &track, nil
}

func (tr *TrackRepository) UpdateTracksByID(id uint, trackRequest models.Tracks) error {
	var track models.Tracks
	if err := tr.db.First(&track, id).Error; err != nil {
		return err
	}

	track.Title = trackRequest.Title
	track.Artist = trackRequest.Artist
	track.AudioData = trackRequest.AudioData

	if err := tr.db.Save(&track).Error; err != nil {
		return err
	}

	return nil
}

func (tr *TrackRepository) DeleteTracksByID(id uint) error {
	var track models.Tracks
	if err := tr.db.First(&track, id).Error; err != nil {
		return err
	}

	if err := tr.db.Delete(&track).Error; err != nil {
		return err
	}

	return nil
}
