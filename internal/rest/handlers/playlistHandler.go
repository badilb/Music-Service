package handlers

import (
	"github.com/gin-gonic/gin"
	"music-service-v2.0.0/internal/models"
	"music-service-v2.0.0/internal/repository"
	"music-service-v2.0.0/internal/rest/dto"
	"music-service-v2.0.0/pkg/logger"
	"net/http"
	"strconv"
)

type PlaylistHandler struct {
	playlistRepo repository.PlaylistRepo
}

func NewPlaylistHandler(PlaylistRepo repository.PlaylistRepo) *PlaylistHandler {
	return &PlaylistHandler{
		playlistRepo: PlaylistRepo,
	}
}

func (h *PlaylistHandler) CreatePlaylist(c *gin.Context) {
	var playlistForm models.Playlist
	if err := c.BindJSON(&playlistForm); err != nil {
		logger.GetLogger().Error("Invalid registration request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.playlistRepo.CreatePlaylist(&playlistForm); err != nil {
		logger.GetLogger().Error("Failed to save playlist:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save playlist: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Playlist created successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Playlist created successfully"})
}

func (h *PlaylistHandler) GetAllPlaylists(c *gin.Context) {
	playlists, err := h.playlistRepo.GetAllPlaylist()
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Playlist fetched successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Playlist fetched successfully", "data": playlists})
}

func (h *PlaylistHandler) GetPlaylistByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	playlist, err := h.playlistRepo.GetPlaylistByID(uint(id))
	if err != nil {
		logger.GetLogger().Error("Failed to get playlist by id:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get playlist by id: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Playlist fetched by id successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Playlist fetched by id successfully", "data": playlist})
}

func (h *PlaylistHandler) UpdatePlaylistByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	var playlistForm dto.PlaylistForm
	if err := c.BindJSON(&playlistForm); err != nil {
		logger.GetLogger().Error("Invalid registration request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.playlistRepo.UpdatePlaylistByID(uint(id), playlistForm); err != nil {
		logger.GetLogger().Error("Failed to update from db:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update from db"})
		return
	}

	logger.GetLogger().Info("Playlist updated successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Playlist updated successfully", "success": true})
}

func (h *PlaylistHandler) DeletePlaylistByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to fetch playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists from db: " + err.Error()})
		return
	}

	if err := h.playlistRepo.DeletePlaylistByID(uint(id)); err != nil {
		logger.GetLogger().Error("Failed to delete playlists from db:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete playlists from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Playlist deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Playlist deleted successfully", "success": true})
}
