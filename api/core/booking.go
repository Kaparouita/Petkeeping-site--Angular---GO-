package core

import (
	"359/domain"
	"fmt"
)

func (srv *Service) GetBooking(booking *domain.Booking) *domain.Booking {
	err := srv.db.GetBooking(booking)
	if err != nil {
		booking.StatusCode = 400
		booking.Message = fmt.Sprintf("Couldnt get Booking : %v", err)
		return booking
	}

	booking.StatusCode = 200
	return booking
}

func (srv *Service) CreateBooking(booking *domain.Booking) *domain.Booking {
	booking.Status = "requested"
	err := srv.db.SaveBooking(booking)
	if err != nil {
		booking.StatusCode = 400
		booking.Message = fmt.Sprintf("Couldnt create Booking : %v", err)
		return booking
	}

	booking.StatusCode = 200
	return booking
}

func (srv *Service) UpdateBooking(booking *domain.Booking) *domain.Booking {
	err := srv.db.UpdateBooking(booking)
	if err != nil {
		booking.StatusCode = 400
		booking.Message = fmt.Sprintf("Couldnt update Booking : %v", err)
		return booking
	}
	booking.StatusCode = 201
	return booking
}

func (srv *Service) GetBookings() ([]domain.Booking, error) {
	bookings, err := srv.db.GetBookings()
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (srv *Service) GetBookingsByKeeperId(keeperId int) ([]domain.Booking, error) {
	bookings, err := srv.db.GetBookingsByKeeperId(uint(keeperId))
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (srv *Service) DeleteBooking(booking *domain.Booking) *domain.Response {
	resp := &domain.Response{}
	err := srv.db.DeleteBooking(booking)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete Booking : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted Booking successfully"
	return resp
}

func (srv *Service) GetMoney() ([]int, error) {
	var money []int
	keeperMoney := 0
	appMoney := 0
	bookings, err := srv.db.GetBookings()
	if err != nil {
		return nil, err
	}
	for _, booking := range bookings {
		keeperMoney += booking.Price * 85 / 100
		appMoney += booking.Price * 15 / 100
	}
	money = append(money, keeperMoney)
	money = append(money, appMoney)

	return money, nil
}

func (srv *Service) GetBookingsNumberByKeeperId(userId int) (int, error) {
	num := 0
	bookings, err := srv.db.GetBookingsByKeeperId(uint(userId))
	if err != nil {
		return 0, err
	}
	for _, booking := range bookings {
		if booking.Status == "finished" {
			num++
		}
	}
	return num, nil
}

func (srv *Service) GetPetKeepersDays(userId int) (int, error) {
	days := 0
	bookings, err := srv.db.GetBookingsByKeeperId(uint(userId))
	if err != nil {
		return 0, err
	}
	for _, booking := range bookings {
		if booking.Status == "finished" {
			days += booking.EndDate.Day() - booking.StartDate.Day()
		}
	}
	return days, nil
}
