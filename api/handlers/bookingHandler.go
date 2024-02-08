package handlers

import (
	"359/domain"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) CreateBooking(c *fiber.Ctx) error {
	booking := &domain.Booking{}
	err := json.Unmarshal(c.Body(), &booking)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to create booking")
	}

	resp := handler.Srv.CreateBooking(booking)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) UpdateBooking(c *fiber.Ctx) error {
	booking := &domain.Booking{}
	err := json.Unmarshal(c.Body(), &booking)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := handler.Srv.UpdateBooking(booking)
	if resp.StatusCode != 201 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetBooking(c *fiber.Ctx) error {
	booking := &domain.Booking{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	booking.Id = uint(id)

	resp := handler.Srv.GetBooking(booking)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetBookings(c *fiber.Ctx) error {
	bookings, err := handler.Srv.GetBookings()
	if err != nil {
		return c.Status(500).JSON("Unable to get bookings")
	}

	return c.Status(200).JSON(bookings)
}

func (handler *Handler) GetBookingsByKeeperId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	bookings, err := handler.Srv.GetBookingsByKeeperId(int(id))
	if err != nil {
		return c.Status(500).JSON("Unable to get bookings")
	}
	return c.Status(200).JSON(bookings)
}

func (handler *Handler) DeleteBooking(c *fiber.Ctx) error {
	booking := &domain.Booking{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	booking.Id = uint(id)

	resp := handler.Srv.DeleteBooking(booking)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetBookingsByOwner(c *fiber.Ctx) error {
	owner := &domain.Owner{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	owner.Id = uint(id)

	bookings, err := handler.Srv.GetBookingsByOwner(owner)
	if err != nil {
		return c.Status(500).JSON("Unable to get bookings")
	}

	return c.Status(200).JSON(bookings)
}
