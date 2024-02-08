package domain

import "time"

type Booking struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	OwnerId   uint      `json:"owner_id"`
	KeeperId  uint      `json:"keeper_id"`
	PetId     uint      `json:"pet_id"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Price     int       `json:"price"`
	Status    string    `json:"status"`
	Response
}
