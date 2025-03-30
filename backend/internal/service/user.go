package service

import (
	"errors"
	"net/http"

	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
	"socialNetwork/pkg/utils"
)

type UserService struct {
	userRepo    repository.User
	sessionRepo repository.Session
}

func newUserService(userRepo repository.User, sessionRepo repository.Session) *UserService {
	return &UserService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *UserService) ServiceSingUp(UserProfile entity.UserProfile) (int, error) {
	if err := utils.CheckEmptyInfo(UserProfile); err != nil {
		return http.StatusBadRequest, err
	}

	if err := s.userRepo.CheckUserExist(UserProfile); err != nil {
		return http.StatusBadRequest, err
	}

	if isValid, err := utils.CheckUsernameFormat(UserProfile.Username); err != nil || !isValid {
		if !isValid {
			return http.StatusBadRequest, errors.New("invalid Username Format")
		} else {
			return http.StatusInternalServerError, errors.New("internal Server Error")
		}
	}

	if isValid, err := utils.CheckEmailFormat(UserProfile.Email); err != nil || !isValid {
		if !isValid {
			return http.StatusBadRequest, errors.New("invalid Email Format")
		} else {
			return http.StatusBadRequest, errors.New("internal Server Error")
		}
	}

	if !utils.CheckPasswordFormat(UserProfile.Password) {
		return http.StatusBadRequest, errors.New("invalid Password Format")
	}

	valid_age, err := utils.CheckAge(UserProfile.BirthDate)
	if err != nil {
		return http.StatusBadRequest, err
	} else if valid_age == -1 {
		return http.StatusBadRequest, errors.New("user is under 16 years old")
	}

	status, err := s.userRepo.RegisterUser(UserProfile)
	if err != nil {
		return status, err
	}

	return status, nil
}

func (s *UserService) ServiceLogIn(UserProfile entity.UserProfile) (entity.Session, entity.UserProfile, int, error) {
	if err := utils.CheckEmptyInfoLogin(UserProfile); err != nil {
		return entity.Session{}, entity.UserProfile{}, http.StatusBadRequest, err
	}

	ProfileInfo, status, err := s.userRepo.LoginUser(UserProfile)
	if err != nil {
		return entity.Session{}, entity.UserProfile{}, status, err
	}

	token, status, err := utils.GenerateToken()
	if err != nil {
		return entity.Session{}, entity.UserProfile{}, status, err
	}

	session := entity.Session{
		UserID: uint(ProfileInfo.Id),
		Token:  token,
	}

	status, err = s.sessionRepo.CreatSession(session)
	if err != nil {
		return entity.Session{}, entity.UserProfile{}, status, err
	}
	return session, ProfileInfo, http.StatusOK, nil
}
