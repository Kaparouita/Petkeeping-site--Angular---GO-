package repositories

import "359/domain"

func (db *Db) CreateMessage(message *domain.Message) error {
	return db.Model(message).Create(message).Error
}

func (db *Db) GetMessagesByUsername(username string) ([]domain.Message, error) {
	var messages []domain.Message
	err := db.Model(domain.Message{}).Where("from_name = ? OR to_name = ?", username, username).Find(&messages).Error
	return messages, err
}
