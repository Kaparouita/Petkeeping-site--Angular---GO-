package repositories

import (
	"359/domain"
	"errors"
	"strings"
)

func (db *Db) SaveOwner(user *domain.Owner) error {
	var keeperCount int64
	db.Model(&domain.Keeper{}).Where("username = ? OR email = ?", user.Username, user.Email).Count(&keeperCount)
	if keeperCount > 0 {
		return errors.New("owner has the same username or email as a keeper")
	}
	return db.Model(user).Create(user).Error
}

func (db *Db) UpdateOwner(user *domain.Owner) error {
	var keeperCount int64
	db.Model(&domain.Keeper{}).Where("username = ? OR email = ?", user.Username, user.Email).Count(&keeperCount)
	if keeperCount > 0 {
		return errors.New("owner has the same username or email as a keeper")
	}
	return db.Model(user).Updates(user).Error
}

func (db *Db) GetOwners() ([]domain.Owner, error) {
	var users []domain.Owner
	err := db.Find(&users).Error
	for i := range users {
		pets, _ := db.GetPetsByOwner(&users[i])
		users[i].Pets = pets
	}
	return users, err

}

func (db *Db) GetOwner(user *domain.Owner) error {
	err := db.Model(user).Find(user).Error
	if err != nil {
		return err
	}
	pets, _ := db.GetPetsByOwner(user)
	user.Pets = pets
	return nil
}

func (db *Db) DeleteOwner(user *domain.Owner) error {
	return db.Delete(&domain.Owner{Id: user.Id}).Error
}

func (db *Db) AvailableKeepers(petsType []string) ([]domain.Keeper, error) {
	var keepers []domain.Keeper
	if len(petsType) == 0 {
		return keepers, nil
	}
	for i := range petsType {
		if strings.ToLower(petsType[i]) == "dog" {
			var dogs []domain.Keeper
			err := db.Model(&domain.Keeper{}).Where("dog_keep = ?", true).Find(&dogs).Error
			if err != nil {
				return nil, err
			}
			for i := range dogs {
				if db.BookingStatus(&dogs[i]) == nil {
					keepers = append(keepers, dogs[i])
				}
			}
		}
		if strings.ToLower(petsType[i]) == "cat" {
			var cats []domain.Keeper
			err := db.Model(&domain.Keeper{}).Where("cat_keep = ?", true).Find(&cats).Error
			if err != nil {
				return nil, err
			}
			for i := range cats {
				if db.BookingStatus(&cats[i]) == nil {
					keepers = append(keepers, cats[i])
				}
			}
		}
	}
	return keepers, nil
}

func (db *Db) SaveKeeper(user *domain.Keeper) error {
	var keeperCount int64
	db.Model(&domain.Owner{}).Where("username = ? OR email = ?", user.Username, user.Email).Count(&keeperCount)
	if keeperCount > 0 {
		return errors.New("keeper has the same username or email as an owner")
	}
	return db.Model(user).Create(user).Error
}

func (db *Db) UpdateKeeper(user *domain.Keeper) error {
	var keeperCount int64
	db.Model(&domain.Owner{}).Where("username = ? OR email = ?", user.Username, user.Email).Count(&keeperCount)
	if keeperCount > 0 {
		return errors.New("keeper has the same username or email as an owner")
	}
	return db.Model(user).Updates(user).Error
}

func (db *Db) GetKeepers() ([]domain.Keeper, error) {
	var users []domain.Keeper
	err := db.Find(&users).Error
	return users, err
}

func (db *Db) GetKeeper(user *domain.Keeper) error {
	return db.Model(user).Find(user).Error
}

func (db *Db) DeleteKeeper(user *domain.Keeper) error {
	return db.Delete(&domain.Keeper{Id: user.Id}).Error
}

func (db *Db) Login(cred *domain.LoginResp) error {
	var owner domain.Owner
	var keeper domain.Keeper
	err := db.Model(&owner).Where("username = ?", cred.Username).Where("password = ?", cred.Password).First(&owner).Error
	if err != nil {
		err = db.Model(&keeper).Where("username = ?", cred.Username).Where("password = ?", cred.Password).First(&keeper).Error
		if err != nil {
			return errors.New("invalid username or password")
		}
		cred.UserId = keeper.Id
		cred.UserType = keeper.UserType
		return nil
	}
	cred.UserId = owner.Id
	cred.UserType = owner.UserType
	return nil
}

func (db *Db) LoginAdmin(cred *domain.LoginResp) error {
	var admin domain.Admin
	err := db.Model(&admin).Where("username = ?", cred.Username).Where("password = ?", cred.Password).First(&admin).Error
	if err != nil {
		return errors.New("invalid username or password")
	}
	cred.UserId = admin.Id
	cred.UserType = admin.UserType
	return nil
}

func (db *Db) GetAdmin(user *domain.Admin) error {
	return db.Model(user).Find(user).Error
}

func (db *Db) SaveAdmin(user *domain.Admin) error {
	err := db.Create(user).Error
	return err
}

func (db *Db) CreateReview(review *domain.Review) error {
	err := db.Create(review).Error
	return err
}

func (db *Db) GetReviewsByKeeper(keeper *domain.Keeper) ([]domain.Review, error) {
	var reviews []domain.Review
	err := db.Model(&domain.Review{}).Where("keeper_id = ?", keeper.Id).Find(&reviews).Error
	return reviews, err
}
