// @title           SpyCat Agency API
// @version         1.0
// @description     A spy cat management system API.

package main

import (
	_ "SpyCatAgency/cmd/api/docs"
	"SpyCatAgency/internal/client"
	"SpyCatAgency/internal/config"
	"SpyCatAgency/internal/handler"
	"SpyCatAgency/internal/infrastructure/database"
	"SpyCatAgency/internal/infrastructure/repository"
	"SpyCatAgency/internal/logger"
	"SpyCatAgency/internal/server"
	"SpyCatAgency/internal/service"
	"context"
	"os"
	"os/signal"
	"syscall"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	// Initialize database
	db, err := database.NewPostgresDB(cfg.GetDSN())
	if err != nil {
		logger.Fatal(ctx, err)
	}

	// Initialize CatAPI client
	catAPI := client.NewCatAPI(cfg.CatAPIURL, cfg.CatAPIKey)

	// Initialize repositories
	catRepo := repository.NewCatRepository(db.DB)
	missionRepo := repository.NewMissionRepository(db.DB)
	targetRepo := repository.NewTargetRepository(db.DB)

	// Initialize services
	catService := service.NewCatService(catRepo, catAPI)
	missionService := service.NewMissionService(missionRepo, targetRepo, catRepo)

	// Initialize handlers
	catHandler := handler.NewCatHandler(catService)
	missionHandler := handler.NewMissionHandler(missionService)

	// Initialize server
	srv := server.NewServer(cfg)

	// Add Swagger UI endpoint
	srv.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	catHandler.RegisterRoutes(srv.Router)
	missionHandler.RegisterRoutes(srv.Router)

	// Start server
	go srv.Run(ctx)

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(ctx, "Shutting down server...")
}
