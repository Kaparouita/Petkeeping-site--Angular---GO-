package handlers

import (
	"359/domain"
	"359/ports"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Srv ports.Service
}

func NewHandler(srv ports.Service) *Handler {
	return &Handler{
		Srv: srv,
	}
}

func (handler *Handler) RegisterOwner(c *fiber.Ctx) error {
	user := &domain.Owner{}
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to register user")
	}

	resp := handler.Srv.RegisterOwner(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.SendStatus(resp.StatusCode)
}

func (handler *Handler) UpdateOwner(c *fiber.Ctx) error {
	user := &domain.Owner{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.UpdateOwner(user)
	if resp.StatusCode != 201 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetOwner(c *fiber.Ctx) error {
	user := &domain.Owner{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.GetOwner(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetOwners(c *fiber.Ctx) error {
	users, err := handler.Srv.GetOwners()
	if err != nil {
		return c.Status(404).JSON("Unable to retrieve Users")
	}

	return c.Status(200).JSON(users)
}

func (handler *Handler) DeleteOwner(c *fiber.Ctx) error {
	user := &domain.Owner{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.DeleteOwner(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp.Message)
}

func (handler *Handler) RegisterKeeper(c *fiber.Ctx) error {
	user := &domain.Keeper{}
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := handler.Srv.RegisterKeeper(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.SendStatus(resp.StatusCode)
}

func (handler *Handler) UpdateKeeper(c *fiber.Ctx) error {
	user := &domain.Keeper{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.UpdateKeeper(user)
	if resp.StatusCode != 201 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetKeeper(c *fiber.Ctx) error {
	user := &domain.Keeper{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.GetKeeper(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetKeepers(c *fiber.Ctx) error {
	users, err := handler.Srv.GetKeepers()
	if err != nil {
		return c.Status(404).JSON("Unable to retrieve Users")
	}

	return c.Status(200).JSON(users)
}

func (handler *Handler) DeleteKeeper(c *fiber.Ctx) error {
	user := &domain.Keeper{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.DeleteKeeper(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp.Message)
}

func (handler *Handler) GetAdmin(c *fiber.Ctx) error {
	user := &domain.Admin{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Id = uint(id)

	resp := handler.Srv.GetAdmin(user)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) Login(c *fiber.Ctx) error {
	cred := &domain.LoginResp{}
	err := json.Unmarshal(c.Body(), &cred)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := handler.Srv.Login(cred)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) LoginAdmin(c *fiber.Ctx) error {
	cred := &domain.LoginResp{}
	err := json.Unmarshal(c.Body(), &cred)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := handler.Srv.LoginAdmin(cred)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) AvailableKeepers(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // owner id
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	keepers, err := handler.Srv.AvailableKeepers(id)
	if err != nil {
		return c.Status(404).JSON(fmt.Sprintf("Unable to retrieve keepers: %v", err))
	}

	return c.Status(200).JSON(keepers)
}

func (handler *Handler) OrderClosestKeepers(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // owner id
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	orderBy := c.Query("orderBy")
	if orderBy == "" {
		orderBy = "distance"
	}
	owner := &domain.Owner{}
	owner.Id = uint(id)
	keepers, err := handler.Srv.OrderClosestKeepers(owner, orderBy)
	if err != nil {
		return c.Status(404).JSON(fmt.Sprintf("Unable to retrieve keepers: %v", err))
	}

	return c.Status(200).JSON(keepers)
}

func (handler *Handler) CreateReview(c *fiber.Ctx) error {
	review := &domain.Review{}
	err := json.Unmarshal(c.Body(), &review)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := handler.Srv.CreateReview(review)
	if resp == nil {
		return c.Status(400).JSON("Unable to create review")
	}

	return c.Status(200).JSON(resp)
}

func (handler *Handler) GetReviewsByKeeper(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // keeper id
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	keeper := &domain.Keeper{Id: uint(id)}
	reviews, err := handler.Srv.GetReviewsByKeeper(keeper)
	if err != nil {
		return c.Status(404).JSON(fmt.Sprintf("Unable to retrieve reviews: %v", err))
	}

	return c.Status(200).JSON(reviews)
}

func (handler *Handler) GetBookingsNumberByKeeperId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // keeper id
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	bookingsNumber, err := handler.Srv.GetBookingsNumberByKeeperId(id)
	if err != nil {
		return c.Status(404).JSON(fmt.Sprintf("Unable to retrieve bookings number: %v", err))
	}

	return c.Status(200).JSON(bookingsNumber)
}

func (handler *Handler) GetPetKeepersDays(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // keeper id
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	keepingDays, err := handler.Srv.GetPetKeepersDays(id)
	if err != nil {
		return c.Status(404).JSON(fmt.Sprintf("Unable to retrieve pet keeping days: %v", err))
	}

	return c.Status(200).JSON(keepingDays)
}
