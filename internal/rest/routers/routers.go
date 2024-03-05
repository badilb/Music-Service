package routers

import (
	"github.com/gin-gonic/gin"
	"music-service-v2.0.0/internal/rest/handlers"
	"music-service-v2.0.0/pkg/middleware"
)

type Routers struct {
	authHandlers    handlers.AuthHandlers
	playlistHandler handlers.PlaylistHandler
	trackHandler    handlers.TrackHandler
}

func NewRouters(authHandlers handlers.AuthHandlers, playlistHandler handlers.PlaylistHandler, trackHandler handlers.TrackHandler) *Routers {
	return &Routers{
		authHandlers:    authHandlers,
		playlistHandler: playlistHandler,
		trackHandler:    trackHandler,
	}
}

func (r *Routers) SetupRoutes(app *gin.Engine) {
	authRouter := app.Group("/auth")
	{
		authRouter.GET("/profile", middleware.RequireAuthMiddleware, r.authHandlers.Profile)
		authRouter.POST("/register", r.authHandlers.Register)
		authRouter.POST("/login", r.authHandlers.Login)
		authRouter.POST("/logout", middleware.RequireAuthMiddleware, r.authHandlers.Logout)
		authRouter.PUT("/update", middleware.RequireAuthMiddleware, r.authHandlers.Update)
		authRouter.DELETE("/delete", middleware.RequireAuthMiddleware, r.authHandlers.Delete)
	}

	playlistRouter := app.Group("/playlists")
	{
		playlistRouter.GET("/", middleware.RequireAuthMiddleware, r.playlistHandler.GetAllPlaylists)
		playlistRouter.GET("/:id", middleware.RequireAuthMiddleware, r.playlistHandler.GetPlaylistByID)
		playlistRouter.POST("/add", middleware.RequireAuthMiddleware, r.playlistHandler.CreatePlaylist)
		playlistRouter.PUT("/:id/update", middleware.RequireAuthMiddleware, r.playlistHandler.UpdatePlaylistByID)
		playlistRouter.DELETE("/:id/delete", middleware.RequireAuthMiddleware, r.playlistHandler.DeletePlaylistByID)
	}

	trackRouter := app.Group("/track")
	{
		trackRouter.GET("/:trackID", middleware.RequireAuthMiddleware, r.trackHandler.GetTrackByID)
		trackRouter.POST("/:playlistID/add", middleware.RequireAuthMiddleware, r.trackHandler.CreateTrack)
		trackRouter.PUT("/:trackID/update", middleware.RequireAuthMiddleware, r.trackHandler.UpdateTrackByID)
		trackRouter.DELETE("/:trackID/delete", middleware.RequireAuthMiddleware, r.trackHandler.DeleteTrackByID)
	}
}
