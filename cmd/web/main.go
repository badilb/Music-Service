package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"music-service-v2.0.0/internal/config"
	"music-service-v2.0.0/internal/db"
	"music-service-v2.0.0/internal/repository"
	"music-service-v2.0.0/internal/rest/handlers"
	"music-service-v2.0.0/internal/rest/routers"
	"music-service-v2.0.0/pkg/compose"
	"music-service-v2.0.0/pkg/logger"
	"music-service-v2.0.0/pkg/session"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func initializeDB() config.Database {
	dbConfig := config.Database{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
		Name:     os.Getenv("POSTGRES_NAME"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	return dbConfig
}

func initializeSession() config.Session {
	sessionConfig := config.Session{
		Secret: os.Getenv("SESSION_SECRET_KEY"),
		Name:   os.Getenv("SESSION_NAME"),
		Key:    os.Getenv("SESSION_KEY"),
	}

	return sessionConfig
}

func initializeRedis() config.RedisSessionConfig {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error converting REDIS_DB to int: %s", err)
	}

	sessionConfig := initializeSession()

	redisConfig := config.RedisSessionConfig{
		ConnectionTimeoutSeconds: os.Getenv("REDIS_CONNECTION_TIMEOUT_SECONDS"),
		NetworkType:              os.Getenv("REDIS_NETWORK_TYPE"),
		Host:                     os.Getenv("REDIS_HOST"),
		Port:                     os.Getenv("REDIS_PORT"),
		Password:                 os.Getenv("REDIS_PASSWORD"),
		DB:                       redisDB,
		Session:                  sessionConfig,
	}

	return redisConfig
}

var appConfig config.App

func main() {
	logger.InitLogger()
	err := compose.StartDockerComposeService()

	if err != nil {
		logger.GetLogger().Fatal("Error starting Docker Compose service:", err)
	}
	logger.GetLogger().Info("Docker Compose service started successfully!")

	appConfig = config.App{
		PORT:  os.Getenv("APP_PORT"),
		DB:    initializeDB(),
		Redis: initializeRedis(),
	}
	session.NewUserSessionStore(&appConfig.Redis)

	dbInstance, err := db.GetDBInstance(appConfig.DB)
	if err != nil {
		logger.GetLogger().Fatal("Error initializing DB:", err)
	}

	userRepo := repository.NewUserRepository(dbInstance)
	roleRepo := repository.NewRoleRepository(dbInstance)
	playlistRepo := repository.NewPlaylistRepository(dbInstance)
	trackRepo := repository.NewTrackRepository(dbInstance)

	authHandlers := handlers.NewAuthHandlers(userRepo, roleRepo)
	playlistHandlers := handlers.NewPlaylistHandler(playlistRepo)
	trackHandlers := handlers.NewTrackHandler(trackRepo, playlistRepo)

	r := gin.Default()

	router := routers.NewRouters(*authHandlers, *playlistHandlers, *trackHandlers)
	router.SetupRoutes(r)
	r.Use(rateLimitMiddleware())

	server := &http.Server{
		Addr:    ":" + appConfig.PORT,
		Handler: r,
	}

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		logger.GetLogger().Info("Server is shutting down...")

		if err := compose.StopDockerComposeService(); err != nil {
			logger.GetLogger().Error("Error stopping Docker Compose service:", err)
		}
		logger.GetLogger().Info("Docker Compose service stopped successfully!")

		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.GetLogger().Fatal("Server shutdown error:", err)
		}

		logger.GetLogger().Info("Server has gracefully stopped")
		os.Exit(0)
	}()

	logger.GetLogger().Info("Server is running on :" + appConfig.PORT)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.GetLogger().Fatal("Error starting server:", err)
	}
}

func rateLimitMiddleware() gin.HandlerFunc {
	limiter := time.Tick(time.Second)

	return func(c *gin.Context) {
		select {
		case <-limiter:
			c.Next()
		default:
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
		}
	}
}
