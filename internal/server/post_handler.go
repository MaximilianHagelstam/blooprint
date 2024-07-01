package server

import "github.com/gofiber/fiber/v2"

func (s *Server) GetPostsHandler(c *fiber.Ctx) error {
	posts, err := s.Repository.GetPosts()
	if err != nil {
		s.logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"posts": posts})
}
