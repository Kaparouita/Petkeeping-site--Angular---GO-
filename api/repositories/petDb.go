package repositories

import (
	"359/domain"
	"errors"
	"regexp"
	"strings"
)

func (db *Db) SavePet(pet *domain.Pet) error {
	if !IsPetIDValid(pet) {
		return errors.New("pet format is invalid")
	}

	// Check if pet owner exists
	var owner domain.Owner
	err := db.Model(&owner).Where("id = ?", pet.OwnerId).First(&owner).Error
	if err != nil {
		return errors.New("pet owner does not exist")
	}

	// Check if pet already exists
	var existingPet domain.Pet
	err = db.Model(&existingPet).Where("pet_id = ?", pet.PetID).First(&existingPet).Error
	if err == nil {
		return errors.New("pet_id already exists")
	}

	err = db.Model(pet).Create(pet).Error
	if err != nil {
		return err
	}
	// Add pet to owner's pets
	owner.Pets = append(owner.Pets, *pet)
	return db.Model(&owner).Updates(&owner).Error
}

func (db *Db) UpdatePet(pet *domain.Pet) error {
	return db.Model(pet).Updates(pet).Error
}

func (db *Db) GetPets() ([]domain.Pet, error) {
	var pets []domain.Pet
	err := db.Find(&pets).Error
	return pets, err
}

func (db *Db) GetPetsByOwner(owner *domain.Owner) ([]domain.Pet, error) {
	var pets []domain.Pet
	err := db.Model(&domain.Pet{}).Where("owner_id = ?", owner.Id).Find(&pets).Error
	return pets, err
}

func (db *Db) GetPet(pet *domain.Pet) error {
	return db.Model(pet).Find(pet).Error
}

func (db *Db) DeletePet(pet *domain.Pet) error {
	err := db.Model(pet).Where("pet_id = ?", pet.PetID).Find(pet).Error
	if err != nil {
		return err
	}
	return db.Delete(pet).Error
}

func IsPhotoURLValid(photoURL string) bool {
	return strings.HasPrefix(photoURL, "http")
}

func (db *Db) GetPetsByTypeAndBreed(pagin *domain.PetTypeBreedPagination) ([]domain.Pet, error) {
	var pets []domain.Pet

	if pagin.Breed == "all" {
		err := db.Where("type = ?", pagin.Type).Where("weight <= ?", pagin.MaxWeight).Where("weight >= ?", pagin.MinWeight).Find(&pets).Error
		return pets, err
	}

	err := db.Where("type = ?", pagin.Type).Where("breed = ?", pagin.Breed).Where("weight <= ?", pagin.MaxWeight).Where("weight >= ?", pagin.MinWeight).
		Find(&pets).Error
	return pets, err
}

func (db *Db) UpdatePetWeight(pet *domain.Pet) error {
	err := db.Model(pet).Where("pet_id = ?", pet.PetID).Find(pet).Error
	if err != nil {
		return err
	}
	return db.Model(pet).Updates(pet).Error
}

func IsPetIDValid(pet *domain.Pet) bool {
	match, _ := regexp.MatchString(`^\d{10}$`, pet.PetID)
	if !match {
		pet.StatusCode = 403
		pet.Message = "PetID must be 10 digits"
		return match
	}
	if pet.BirthYear < 2000 || pet.Weight < 0 {
		pet.StatusCode = 406
		pet.Message = "BirthDate must be after 2000 and Weight must be positive"
		return false
	}
	if strings.HasPrefix(pet.Photo, "http") {
		pet.StatusCode = 406
		pet.Message = "Photo must be a URL"
		return false
	}
	return true
}

func (db *Db) GetNumberOfCats() (int, error) {
	var count int64
	err := db.Model(&domain.Pet{}).Where("type = ?", "cat").Count(&count).Error
	return int(count), err
}

func (db *Db) GetNumberOfDogs() (int, error) {
	var count int64
	err := db.Model(&domain.Pet{}).Where("type = ?", "dog").Count(&count).Error
	return int(count), err
}
