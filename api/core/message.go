package core

import "359/domain"

func (srv *Service) CreateMessage(message *domain.Message) *domain.Message {
	err := srv.db.CreateMessage(message)
	if err != nil {
		return nil
	}
	return message
}

func (srv *Service) GetMessagesByUsername(username string) ([]domain.Message, error) {
	messages, err := srv.db.GetMessagesByUsername(username)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
