package handlers

import (
	"359/domain"
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) RegisterPet(c *fiber.Ctx) error {
	pet := &domain.Pet{}
	err := json.Unmarshal(c.Body(), &pet)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to register pet")
	}

	resp := handler.Srv.SavePet(pet)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}
	resp.Message = "Registered pet successfully"
	return c.Status(resp.StatusCode).JSON(resp.Message)
}

func (handler *Handler) UpdatePet(c *fiber.Ctx) error {
	pet := &domain.Pet{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = json.Unmarshal(c.Body(), &pet)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	pet.ID = id
	resp := handler.Srv.UpdatePet(pet)
	if resp.StatusCode != 201 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}
	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetPet(c *fiber.Ctx) error {
	pet := &domain.Pet{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	pet.ID = id
	resp := handler.Srv.GetPet(pet)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}
	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *Handler) GetPets(c *fiber.Ctx) error {
	pets, err := handler.Srv.GetPets()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to get pets")
	}
	return c.Status(fiber.StatusOK).JSON(pets)
}

func (handler *Handler) DeletePet(c *fiber.Ctx) error {
	pet := &domain.Pet{}
	petID := c.Params("pet_id", "")
	match, _ := regexp.MatchString(`^\d{10}$`, petID)
	if !match {
		return c.Status(406).JSON("Invalid pet id")
	}
	pet.PetID = petID
	resp := handler.Srv.DeletePet(pet)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}
	return c.Status(resp.StatusCode).JSON(resp.Message)
}

func (handler *Handler) GetPetsByTypeAndBreed(c *fiber.Ctx) error {
	petType := c.Params("type", "")
	if petType == "cats" {
		petType = "cat"
	} else if petType == "dogs" {
		petType = "dog"
	} else {
		return c.Status(403).JSON("Invalid pet type")
	}

	breed := c.Params("breed", "all")

	maxWeight, err := strconv.Atoi(c.Query("toWeight", "1000"))
	if err != nil || maxWeight < 0 {
		return c.Status(406).JSON("Invalid max weight")
	}

	minWeight, err := strconv.Atoi(c.Query("fromWeight", "0"))
	if err != nil || minWeight < 0 {
		return c.Status(406).JSON("Invalid min weight")
	}

	pagin := &domain.PetTypeBreedPagination{
		Type:      petType,
		Breed:     breed,
		MinWeight: minWeight,
		MaxWeight: maxWeight,
	}

	pets, err := handler.Srv.GetPetsByTypeAndBreed(pagin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(200).JSON(pets)
}

func (handler *Handler) UpdatePetWeight(c *fiber.Ctx) error {
	petID := c.Params("pet_id", "")
	match, _ := regexp.MatchString(`^\d{10}$`, petID)
	if !match {
		return c.Status(406).JSON("Invalid pet id")
	}
	weight, err := c.ParamsInt("weight")
	if err != nil || weight < 0 {
		return c.Status(406).JSON("Invalid weight")
	}
	pet := &domain.Pet{
		PetID:  petID,
		Weight: weight,
	}
	resp := handler.Srv.UpdatePetWeight(pet)
	if resp.StatusCode != 201 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}
	resp.Message = "Updated pet weight successfully"
	return c.Status(resp.StatusCode).JSON(resp.Message)
}

func (handler *Handler) GetNumberOfCats(c *fiber.Ctx) error {
	count, err := handler.Srv.GetNumberOfCats()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(200).JSON(count)
}

func (handler *Handler) GetNumberOfDogs(c *fiber.Ctx) error {
	count, err := handler.Srv.GetNumberOfDogs()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(200).JSON(count)
}
