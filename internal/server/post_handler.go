package server

import (
	"gostarter/internal"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetPostsHandler(c *fiber.Ctx) error {
	posts, err := s.Repository.GetPosts()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"posts": posts})
}

func (s *Server) CreatePostHandler(c *fiber.Ctx) error {
	post := new(internal.Post)
	if err := c.BodyParser(post); err != nil {
		return fiber.ErrBadRequest
	}

	if post.Caption == "" {
		return fiber.ErrBadRequest
	}

	if err := s.Repository.CreatePost(post); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)
}
