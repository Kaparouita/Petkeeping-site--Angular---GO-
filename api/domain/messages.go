package domain

type Message struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	FromName string `json:"from_name"`
	FromId   int    `json:"from_id"`
	ToName   string `json:"to_name"`
	ToId     int    `json:"to_id"`
	Content  string `json:"content"`
	Date     string `json:"date"`
}
