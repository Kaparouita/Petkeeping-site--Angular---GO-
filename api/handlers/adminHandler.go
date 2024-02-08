package handlers

import (
	"359/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetNumberOfCatsAndDogs(c *fiber.Ctx) error {
	resp := domain.Response{}
	cats, err := h.Srv.GetNumberOfCats()
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Couldnt get number of cats"
	}
	dogs, err := h.Srv.GetNumberOfDogs()
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Couldnt get number of dogs"
	}

	return c.JSON(fiber.Map{
		"cats": cats,
		"dogs": dogs,
	})
}

func (h *Handler) GetMoney(c *fiber.Ctx) error {
	resp := domain.Response{}
	money, err := h.Srv.GetMoney()
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Couldnt get money"
	}

	return c.JSON(fiber.Map{
		"keeper": money[0],
		"app":    money[1],
	})
}

func (h *Handler) GetNumberOfUsers(c *fiber.Ctx) error {
	resp := domain.Response{}
	owners, err := h.Srv.GetOwners()
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Couldnt get number of owners"
	}
	keepers, err := h.Srv.GetKeepers()
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Couldnt get number of keepers"
	}

	return c.JSON(fiber.Map{
		"owners":  len(owners),
		"keepers": len(keepers),
	})
}
