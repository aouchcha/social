package service

import (
	"socialNetwork/internal/repository"
)

type MessagesService struct {
	messagesRepo repository.Message
	usersRepo    repository.User
}

func newMessagesService(msgRepo repository.Message, userRepo repository.User) *MessagesService {
	return &MessagesService{
		messagesRepo: msgRepo,
		usersRepo:    userRepo,
	}
}
