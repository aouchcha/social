package service

import (
	"errors"
	"net/http"

	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
)

type SessionService struct {
	sessionRepo repository.Session
}

func newSessionService(sessionRepo repository.Session) *SessionService {
	return &SessionService{sessionRepo: sessionRepo}
}

func (s *SessionService) ServiceMiddleware(session entity.Session) (int, error) {
	if valid, err := s.sessionRepo.IsTokenExist(session); err != nil || !valid {
		if err != nil {
			return http.StatusInternalServerError, errors.New("internal Server Error")
		} else {
			return http.StatusBadRequest, errors.New("invalid Token")
		}
	}
	return http.StatusOK, nil
}

func (s *SessionService) ServerLogout(session entity.Session) (int, error) {
	if valid, err := s.sessionRepo.IsTokenExist(session); err != nil || !valid {
		if err != nil {
			return http.StatusInternalServerError, err
		} else {
			return -1, nil
		}
	}

	if err := s.sessionRepo.DeleteSessionByToken(session); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}
