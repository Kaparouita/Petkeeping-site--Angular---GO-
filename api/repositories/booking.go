package repositories

import (
	"359/domain"
	"fmt"
)

func (db *Db) GetBooking(booking *domain.Booking) error {
	return db.DB.First(&booking, booking.Id).Error
}

func (db *Db) SaveBooking(booking *domain.Booking) error {
	fmt.Println(booking)
	return db.DB.Create(&booking).Error
}

func (db *Db) UpdateBooking(booking *domain.Booking) error {
	return db.DB.Model(domain.Booking{}).Where("id = ?", booking.Id).Update("status", booking.Status).Error
}

func (db *Db) GetBookings() ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := db.DB.Find(&bookings).Error
	return bookings, err
}

func (db *Db) GetBookingsByKeeperId(id uint) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := db.DB.Where("keeper_id = ?", id).Find(&bookings).Error
	return bookings, err
}

func (db *Db) DeleteBooking(booking *domain.Booking) error {
	return db.DB.Delete(&booking).Error
}

func (db *Db) BookingStatus(keeper *domain.Keeper) *domain.Booking {
	var booking domain.Booking
	err := db.DB.Where("keeper_id = ? AND (status = ? OR status = ?)", keeper.Id, "accepted", "requested").First(&booking).Error
	if err != nil {
		return nil
	}
	return &booking
}

func (db *Db) GetBookingsByOwner(owner *domain.Owner) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := db.DB.Model(&domain.Booking{}).Where("owner_id = ?", owner.Id).Find(&bookings).Error
	return bookings, err
}
