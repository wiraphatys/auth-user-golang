package server

import (
	"banky/config"
	"banky/user/handlers"
	"banky/user/repositories"
	"banky/user/usecases"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type fiberServer struct {
	app *fiber.App
	db  *gorm.DB
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) Server {
	return &fiberServer{
		app: fiber.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *fiberServer) Start() {
	url := fmt.Sprintf("%v:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	s.initializeUserHttpHandler()

	// listen to port
	log.Printf("Server is starting on %v", url)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}

func (s *fiberServer) initializeUserHttpHandler() {
	// init all layers
	userPostgresRepository := repositories.NewUserPostgresRepository(s.db)
	userUsecase := usecases.NewUserUsecaseImpl(userPostgresRepository)
	userHttpHandler := handlers.NewUserHttpHandler(userUsecase)

	// router
	userRouters := s.app.Group("v1/user")
	userRouters.Post("/register", userHttpHandler.RegisterUser)
	userRouters.Post("/signin", userHttpHandler.SignInUser)
}
