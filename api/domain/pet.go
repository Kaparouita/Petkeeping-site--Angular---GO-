package domain

type Pet struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	PetID       string `json:"pet_id" gorm:"uniqueIndex"` // 10 numbers
	OwnerId     int    `json:"owner_id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Type        string `json:"type"`
	Breed       string `json:"breed"`
	Gender      string `json:"gender"`
	BirthYear   int    `json:"birth_year"` // < 2000
	Weight      int    `json:"weight"`     // < 0
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Response
}
