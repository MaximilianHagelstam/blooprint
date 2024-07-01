package main

import (
	"fmt"
	"gostarter/internal/server"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := server.New()

	s.App.Use(recover.New())
	s.App.Use(logger.New())
	s.RegisterRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	err := s.App.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
