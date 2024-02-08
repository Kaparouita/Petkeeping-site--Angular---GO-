package handlers

import (
	"359/domain"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) GetMessagesByUsername(c *fiber.Ctx) error {
	resp := &domain.Response{}

	username := c.Query("name")
	if username == "" {
		resp.StatusCode = fiber.StatusBadRequest
		resp.Message = "Username is required"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	messages, err := handler.Srv.GetMessagesByUsername(username)
	if err != nil {
		resp.StatusCode = fiber.StatusInternalServerError
		resp.Message = "Unable to get messages"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(messages)
}

func (handler *Handler) CreateMessage(c *fiber.Ctx) error {
	message := &domain.Message{}
	resp := &domain.Response{}
	err := c.BodyParser(message)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to create message")
	}
	message = handler.Srv.CreateMessage(message)
	if message == nil {
		resp.StatusCode = fiber.StatusInternalServerError
		resp.Message = "Unable to create message"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(message)
}
