package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"music-service-v2.0.0/internal/rest/handlers"
	"music-service-v2.0.0/pkg/middleware"
	"strings"
	"time"
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
	authRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3006"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3006"
		},
		MaxAge: 12 * time.Hour,
	}))
	{
		authRouter.GET("/profile", middleware.RequireAuthMiddleware, r.authHandlers.Profile)
		authRouter.POST("/register", r.authHandlers.Register)
		authRouter.POST("/login", r.authHandlers.Login)
		authRouter.POST("/logout", middleware.RequireAuthMiddleware, r.authHandlers.Logout)
		authRouter.PUT("/update", middleware.RequireAuthMiddleware, r.authHandlers.Update)
		authRouter.DELETE("/delete", middleware.RequireAuthMiddleware, r.authHandlers.Delete)
	}

	playlistRouter := app.Group("/api/playlists")
	{
		playlistRouter.GET("/", middleware.RequireAuthMiddleware, r.playlistHandler.GetAllPlaylists)
		playlistRouter.GET("/:id", middleware.RequireAuthMiddleware, r.playlistHandler.GetPlaylistByID)
		playlistRouter.POST("/add", middleware.RequireAuthMiddleware, r.playlistHandler.CreatePlaylist)
		playlistRouter.PUT("/:id/update", middleware.RequireAuthMiddleware, r.playlistHandler.UpdatePlaylistByID)
		playlistRouter.DELETE("/:id/delete", middleware.RequireAuthMiddleware, r.playlistHandler.DeletePlaylistByID)
	}

	trackRouter := app.Group("/api/track")
	{
		trackRouter.GET("/:trackID", middleware.RequireAuthMiddleware, r.trackHandler.GetTrackByID)
		trackRouter.POST("/:playlistID/add", middleware.RequireAuthMiddleware, r.trackHandler.CreateTrack)
		trackRouter.PUT("/:trackID/update", middleware.RequireAuthMiddleware, r.trackHandler.UpdateTrackByID)
		trackRouter.DELETE("/:trackID/delete", middleware.RequireAuthMiddleware, r.trackHandler.DeleteTrackByID)
	}
	// SPA fallback
	app.NoRoute(func(c *gin.Context) {
		// Checking that the requested path does not start with '/api' and does not apply to the API
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.File("./web/dist/index.html")
		}
	})

}
