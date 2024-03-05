package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"music-service-v2.0.0/internal/models"
	"music-service-v2.0.0/internal/repository"
	"music-service-v2.0.0/pkg/logger"
	"net/http"
	"strconv"
)

type TrackHandler struct {
	trackRepo    repository.TrackRepo
	playlistRepo repository.PlaylistRepo
}

func NewTrackHandler(trackRepo repository.TrackRepo, playlistRepo repository.PlaylistRepo) *TrackHandler {
	return &TrackHandler{
		trackRepo:    trackRepo,
		playlistRepo: playlistRepo,
	}
}

func (h *TrackHandler) CreateTrack(c *gin.Context) {
	playlistIdStr := c.Param("playlistID")
	playlistId, err := strconv.Atoi(playlistIdStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	var track models.Tracks

	title := c.PostForm("title")
	artist := c.PostForm("artist")
	file, err := c.FormFile("audioData")
	if err != nil {
		logger.GetLogger().Error("Failed to get audio data:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get audio data: " + err.Error()})
		return
	}

	audioFile, err := file.Open()
	if err != nil {
		logger.GetLogger().Error("Failed to open audio file:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open audio file: " + err.Error()})
		return
	}
	defer audioFile.Close()

	audioData, err := ioutil.ReadAll(audioFile)
	if err != nil {
		logger.GetLogger().Error("Failed to read audio file:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read audio file: " + err.Error()})
		return
	}

	track.Title = title
	track.Artist = artist
	track.AudioData = audioData
	//if err := c.BindJSON(&track); err != nil {
	//	logger.GetLogger().Error("Invalid registration request:", err)
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	return
	//}

	playlistData, err := h.playlistRepo.GetPlaylistByID(uint(playlistId))
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	track.Playlist = append(track.Playlist, *playlistData)

	if err := h.trackRepo.CreateTrack(&track); err != nil {
		logger.GetLogger().Error("Failed to save track:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save track: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Track created successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Track created successfully"})
}

func (h *TrackHandler) GetTrackByID(c *gin.Context) {
	trackIDStr := c.Param("trackID")
	trackID, err := strconv.Atoi(trackIDStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	trackData, err := h.trackRepo.GetTracksByID(uint(trackID))
	if err != nil {
		logger.GetLogger().Error("Failed to fetch track from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch track from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Track fetched successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Track fetched successfully", "data": trackData})
}

func (h *TrackHandler) UpdateTrackByID(c *gin.Context) {
	trackIDStr := c.Param("trackID")
	trackID, err := strconv.Atoi(trackIDStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	var track models.Tracks

	title := c.PostForm("title")
	artist := c.PostForm("artist")
	file, err := c.FormFile("audioData")
	if err != nil {
		logger.GetLogger().Error("Failed to get audio data:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get audio data: " + err.Error()})
		return
	}

	audioFile, err := file.Open()
	if err != nil {
		logger.GetLogger().Error("Failed to open audio file:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open audio file: " + err.Error()})
		return
	}
	defer audioFile.Close()

	audioData, err := ioutil.ReadAll(audioFile)
	if err != nil {
		logger.GetLogger().Error("Failed to read audio file:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read audio file: " + err.Error()})
		return
	}

	track.Title = title
	track.Artist = artist
	track.AudioData = audioData

	if err := h.trackRepo.UpdateTracksByID(uint(trackID), track); err != nil {
		logger.GetLogger().Error("Failed to update track from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update track from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Track updated successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Track updated successfully", "status": true})
}

func (h *TrackHandler) DeleteTrackByID(c *gin.Context) {
	trackIDStr := c.Param("trackID")
	trackID, err := strconv.Atoi(trackIDStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	if err := h.trackRepo.DeleteTracksByID(uint(trackID)); err != nil {
		logger.GetLogger().Error("Failed to delete track from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete track from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Track deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Track deleted successfully"})
}
