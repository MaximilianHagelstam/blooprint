package handlers

import (
	"gostarter/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type PostHandlerConfig struct {
	PostRepository repository.PostRepository
	Logger         log.Logger
}

func GetPostsHandler(config *PostHandlerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := config.PostRepository.GetPosts()
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"posts": posts})
	}
}

func GetPostByIDHandler(config *PostHandlerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		post, err := config.PostRepository.GetPostByID(id)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"post": post})
	}
}

func DeletePostHandler(config *PostHandlerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := config.PostRepository.DeletePost(id)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
