package core

import (
	"359/domain"
	"359/ports"
	"fmt"
)

type Service struct {
	db ports.Db
}

func NewService(db ports.Db) *Service {
	return &Service{db: db}
}

func (srv *Service) InitFunction() error {

	if err := srv.InitUsers(); err != nil {
		return err
	}
	return nil
}

func (srv *Service) InitUsers() error {

	Lat := []float64{35, 31621, 35.30681, 35.29681, 35.28681}
	Lon := []float64{25.14361, 25.15375, 25.16375, 25.17375, 25.18375}
	catKeep := []bool{true, false, true, false, true}
	catPrice := []int{10, 0, 9, 0, 8}
	dogKeep := []bool{true, true, false, true, false}
	dogPrice := []int{10, 9, 0, 8, 0}

	// Initialize 5 keepers
	for i := 1; i <= 5; i++ {
		keeper := &domain.Keeper{
			Username:  fmt.Sprintf("keeper%d", i),
			Email:     fmt.Sprintf("KeeperEmail%d", i),
			FirstName: fmt.Sprintf("KeeperFirstName%d", i),
			LastName:  fmt.Sprintf("KeeperLastName%d", i),
			Gender:    "male",
			UserType:  "keeper",
			Country:   "Greece",
			City:      "Athens",
			Address:   "Kifisias 44",
			Phone:     "123456789",
			SpaceType: "house",
			CatKeep:   catKeep[i-1],
			DogKeep:   dogKeep[i-1],
			CatPrice:  catPrice[i-1],
			DogPrice:  dogPrice[i-1],
			SpaceDesc: "I have a big house with a big garden",
			Lat:       Lat[i-1],
			Lon:       Lon[i-1],
			Password:  "123456",
		}
		owner := &domain.Owner{
			Username:  fmt.Sprintf("owner%d", i),
			Email:     fmt.Sprintf("OwnerEmail%d", i),
			FirstName: fmt.Sprintf("OwnerFirstName%d", i),
			LastName:  fmt.Sprintf("OwnerLastName%d", i),
			Gender:    "female",
			UserType:  "owner",
			Country:   "Greece",
			City:      "Athens",
			Address:   "Kifisias 44",
			Phone:     "123456789",
			Lat:       Lat[i-1],
			Lon:       Lon[i-1],
			Password:  "123456",
		}
		srv.db.SaveOwner(owner)
		srv.db.SaveKeeper(keeper)
	}

	user3 := &domain.Admin{
		Id:       1,
		Username: "admin",
		Password: "admin",
		UserType: "admin",
	}
	srv.db.SaveAdmin(user3)
	return nil
}
