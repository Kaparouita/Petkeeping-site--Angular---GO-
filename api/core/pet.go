package core

import (
	"359/domain"
	"fmt"
)

func (srv *Service) SavePet(pet *domain.Pet) *domain.Pet {

	err := srv.db.SavePet(pet)
	if err != nil {
		pet.StatusCode = 400
		pet.Message = fmt.Sprintf("Couldnt create PET : %v", err)
		return pet
	}
	pet.StatusCode = 200
	return pet
}

func (srv *Service) UpdatePet(pet *domain.Pet) *domain.Pet {
	err := srv.db.UpdatePet(pet)
	if err != nil {
		pet.StatusCode = 400
		pet.Message = fmt.Sprintf("Couldnt update PET : %v", err)
		return pet
	}
	pet.StatusCode = 201
	return pet
}

func (srv *Service) GetPet(pet *domain.Pet) *domain.Pet {
	err := srv.db.GetPet(pet)
	if err != nil {
		pet.StatusCode = 400
		pet.Message = fmt.Sprintf("Couldnt get PET : %v", err)
		return pet
	}
	pet.StatusCode = 200
	return pet
}

func (srv *Service) DeletePet(pet *domain.Pet) *domain.Response {
	resp := &domain.Response{}
	err := srv.db.DeletePet(pet)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete PET : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted pet successfully"
	return resp
}

func (srv *Service) GetPets() ([]domain.Pet, error) {
	pets, err := srv.db.GetPets()
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (srv *Service) GetPetsByTypeAndBreed(pagin *domain.PetTypeBreedPagination) ([]domain.Pet, error) {
	pets, err := srv.db.GetPetsByTypeAndBreed(pagin)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (srv *Service) UpdatePetWeight(pet *domain.Pet) *domain.Pet {
	err := srv.db.UpdatePetWeight(pet)
	if err != nil {
		pet.StatusCode = 400
		pet.Message = fmt.Sprintf("Couldnt update PET : %v", err)
		return pet
	}
	pet.StatusCode = 201
	return pet
}

func (srv *Service) GetNumberOfCats() (int, error) {
	count, err := srv.db.GetNumberOfCats()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (srv *Service) GetNumberOfDogs() (int, error) {
	count, err := srv.db.GetNumberOfDogs()
	if err != nil {
		return 0, err
	}
	return count, nil
}
