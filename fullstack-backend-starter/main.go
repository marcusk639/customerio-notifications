package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/customerio/fullstack-backend-homework/controllers"
	"github.com/customerio/fullstack-backend-homework/database"
	"github.com/customerio/fullstack-backend-homework/middleware"
	"github.com/customerio/fullstack-backend-homework/repositories"
	"github.com/customerio/fullstack-backend-homework/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router                           *gin.Engine
	db                               *sqlx.DB
	authController                   *controllers.AuthController
	notificationController           *controllers.NotificationController
	notificationPreferenceController *controllers.NotificationPreferenceController
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	server := newServer()
	server.run()
}

func newServer() *Server {
	db := database.InitDB()
	db.Ping()

	// Initialize repositories
	notificationRepo := repositories.NewNotificationRepository(db)
	preferenceRepo := repositories.NewNotificationPreferenceRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	notificationService := services.NewNotificationService(notificationRepo)
	preferenceService := services.NewNotificationPreferenceService(preferenceRepo)
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	authController := controllers.NewAuthController(userService)
	notificationController := controllers.NewNotificationController(notificationService)
	notificationPreferenceController := controllers.NewNotificationPreferenceController(preferenceService)

	return &Server{
		db:                               db,
		authController:                   authController,
		notificationController:           notificationController,
		notificationPreferenceController: notificationPreferenceController,
		router:                           setupRouter(),
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Auth())
	return r
}

func (s *Server) registerRoutes() {
	// Auth routes
	s.router.POST("/login", s.authController.Login)

	// Notification routes
	notifications := s.router.Group("/notifications")
	{
		notifications.GET("", s.notificationController.GetNotifications)
	}

	// Preference routes
	preferences := s.router.Group("/preferences")
	{
		preferences.POST("", s.notificationPreferenceController.UpdatePreferences)
	}
}

func (s *Server) run() {
	s.registerRoutes()
	s.router.Run(":8081")
}
