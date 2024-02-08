package core

import (
	"359/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
)

func (srv *Service) RegisterOwner(user *domain.Owner) *domain.Owner {
	err := srv.db.SaveOwner(user)
	if err != nil {
		user.StatusCode = 403
		user.Message = fmt.Sprintf("Couldnt register USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) UpdateOwner(user *domain.Owner) *domain.Owner {
	err := srv.db.UpdateOwner(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt update USER : %v", err)
		return user
	}
	user.StatusCode = 201
	return user
}

func (srv *Service) GetOwner(user *domain.Owner) *domain.Owner {
	err := srv.db.GetOwner(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) DeleteOwner(user *domain.Owner) *domain.Response {
	resp := &domain.Response{}
	bookings := []domain.Booking{}
	bookings, err := srv.db.GetBookingsByOwner(user)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get bookings : %v", err)
		return resp
	}
	for _, booking := range bookings {
		err := srv.db.DeleteBooking(&booking)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt delete booking : %v", err)
			return resp
		}
	}

	pets := []domain.Pet{}
	pets, _ = srv.db.GetPetsByOwner(user)
	for _, pet := range pets {
		err := srv.db.DeletePet(&pet)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt delete pet : %v", err)
			return resp
		}
	}

	err = srv.db.DeleteOwner(user)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete USER : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted user successfully"
	return resp
}

func (srv *Service) GetOwners() ([]domain.Owner, error) {
	users, err := srv.db.GetOwners()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (srv *Service) AvailableKeepers(id int) ([]domain.Keeper, error) {
	owner := &domain.Owner{}
	owner.Id = uint(id)
	err := srv.db.GetOwner(owner)
	if err != nil {
		owner.StatusCode = 400
		owner.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		return nil, err
	}
	if len(owner.Pets) == 0 {
		owner.StatusCode = 400
		owner.Message = fmt.Sprintf("Owner has no pets")
		return nil, err
	}
	users, err := srv.db.AvailableKeepers(petTypesOwned(owner.Pets))
	if err != nil {
		return nil, err
	}

	// Trim users from duplicate IDs
	uniqueUsers := make([]domain.Keeper, 0, len(users))
	seen := make(map[uint]bool)
	for _, user := range users {
		if !seen[user.Id] {
			uniqueUsers = append(uniqueUsers, user)
			seen[user.Id] = true
		}
	}

	return uniqueUsers, nil
}

func petTypesOwned(pets []domain.Pet) []string {
	hasCat := false
	hasDog := false

	for _, pet := range pets {
		if pet.Type == "cat" {
			hasCat = true
		} else if pet.Type == "dog" {
			hasDog = true
		}
	}

	if hasCat && hasDog {
		return []string{"cat", "dog"}
	} else if hasCat {
		return []string{"cat"}
	} else if hasDog {
		return []string{"dog"}
	} else {
		return []string{}
	}
}

func (srv *Service) RegisterKeeper(user *domain.Keeper) *domain.Keeper {
	err := srv.db.SaveKeeper(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt register USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) UpdateKeeper(user *domain.Keeper) *domain.Keeper {
	err := srv.db.UpdateKeeper(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt update USER : %v", err)
		return user
	}
	user.StatusCode = 201
	return user
}

func (srv *Service) GetKeeper(user *domain.Keeper) *domain.Keeper {
	err := srv.db.GetKeeper(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) DeleteKeeper(user *domain.Keeper) *domain.Response {
	resp := &domain.Response{}
	bookings := []domain.Booking{}
	bookings, err := srv.db.GetBookingsByKeeperId(user.Id)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get bookings : %v", err)
		return resp
	}
	fmt.Println("BOOKINGS", bookings)
	for _, booking := range bookings {
		err := srv.db.DeleteBooking(&booking)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt delete booking : %v", err)
			return resp
		}
	}

	err = srv.db.DeleteKeeper(user)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete USER : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted user successfully"
	return resp
}

func (srv *Service) GetKeepers() ([]domain.Keeper, error) {
	users, err := srv.db.GetKeepers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (srv *Service) Login(cred *domain.LoginResp) *domain.LoginResp {
	err := srv.db.Login(cred)
	if err != nil {
		cred.StatusCode = 400
		cred.Message = fmt.Sprintf("Couldnt login USER : %v", err)
		return cred
	}
	cred.StatusCode = 200
	fmt.Println("CORE", cred)
	return cred
}

func (srv *Service) GetAdmin(user *domain.Admin) *domain.Admin {
	err := srv.db.GetAdmin(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt get ADMIN: %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) LoginAdmin(cred *domain.LoginResp) *domain.LoginResp {
	err := srv.db.LoginAdmin(cred)
	if err != nil {
		cred.StatusCode = 400
		cred.Message = fmt.Sprintf("Couldnt login ADMIN : %v", err)
		return cred
	}
	cred.StatusCode = 200
	return cred
}

func (srv *Service) GetBookingsByOwner(owner *domain.Owner) ([]domain.Booking, error) {
	bookings, err := srv.db.GetBookingsByOwner(owner)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

type KeeperDistance struct {
	Keeper domain.Keeper
	Value  float64
}

func (srv *Service) OrderClosestKeepers(owner *domain.Owner, orderBy string) ([]domain.Keeper, error) {
	err := srv.db.GetOwner(owner)
	if err != nil {
		owner.StatusCode = 400
		owner.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		return nil, err
	}
	keepers, err := srv.AvailableKeepers(int(owner.Id))
	if err != nil {
		return nil, err
	}
	keepersDistance := []KeeperDistance{}

	for _, keeper := range keepers {
		resp, err := srv.GetDistance(owner.Lat, owner.Lon, keeper.Lat, keeper.Lon)
		if err != nil {
			return nil, err
		}

		if orderBy == "distance" {
			if len(resp.Distances) == 0 {
				keepersDistance = append(keepersDistance, KeeperDistance{keeper, 0})
				continue
			}

			keepersDistance = append(keepersDistance, KeeperDistance{keeper, resp.Distances[0][0]})
		}
		if orderBy == "duration" {
			if len(resp.Durations) == 0 {
				keepersDistance = append(keepersDistance, KeeperDistance{keeper, 0})
				continue
			}
			keepersDistance = append(keepersDistance, KeeperDistance{keeper, resp.Durations[0][0]})
		}

	}
	return orderKeepersDistance(keepersDistance, orderBy), nil
}

func orderKeepersDistance(keepersDist []KeeperDistance, orderBy string) []domain.Keeper {
	keepers := []domain.Keeper{}
	if orderBy == "distance" {
		sort.Slice(keepersDist, func(i, j int) bool {
			return keepersDist[i].Value < keepersDist[j].Value
		})
	}
	if orderBy == "duration" {
		sort.Slice(keepersDist, func(i, j int) bool {
			return keepersDist[i].Value < keepersDist[j].Value
		})
	}
	for _, keeperDist := range keepersDist {
		keepers = append(keepers, keeperDist.Keeper)
	}
	return keepers
}

type MatrixResponse struct {
	Distances [][]float64 `json:"distances"`
	Durations [][]float64 `json:"durations"`
}

func (srv *Service) GetDistance(lat1, lon1, lat2, lon2 float64) (*MatrixResponse, error) {
	url := constructURL(lat1, lon1, lat2, lon2)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "9b7cf22216msh6f6222ee8772ef3p12fceejsn1e75992b55cb")
	req.Header.Add("X-RapidAPI-Host", "trueway-matrix.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var matrixResponse MatrixResponse
	err := json.Unmarshal(body, &matrixResponse)
	if err != nil {
		return nil, err
	}

	distances := matrixResponse.Distances
	durations := matrixResponse.Durations

	fmt.Println(distances)
	fmt.Println(durations)

	return &matrixResponse, nil
}

func constructURL(lat1, lon1, lat2, lon2 float64) string {
	baseUrl := "https://trueway-matrix.p.rapidapi.com/CalculateDrivingMatrix"
	origins := fmt.Sprintf("%f,%f", lat1, lon1)
	destinations := fmt.Sprintf("%f,%f", lat2, lon2)

	return fmt.Sprintf("%s?origins=%s&destinations=%s", baseUrl, url.QueryEscape(origins), url.QueryEscape(destinations))
}

func (srv *Service) CreateReview(review *domain.Review) *domain.Review {
	err := srv.db.CreateReview(review)
	if err != nil {
		review.Comment = fmt.Sprintf("Couldnt create review : %v", err)
	}
	return review
}

func (srv *Service) GetReviewsByKeeper(keeper *domain.Keeper) ([]domain.Review, error) {
	reviews, err := srv.db.GetReviewsByKeeper(keeper)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
