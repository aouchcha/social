package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"socialNetwork/internal/entity"
)

func CheckEmptyInfo(user entity.UserProfile) error {
	fields := []string{user.FirstName, user.LastName, user.Email, user.Password, user.BirthDate}
	for _, field := range fields {
		if strings.TrimSpace(field) == "" {
			return errors.New("all fields are required")
		}
	}
	return nil
}

func CheckEmptyInfoLogin(user entity.UserProfile) error {
	fields := []string{user.Email, user.Password}
	for _, field := range fields {
		if strings.TrimSpace(field) == "" {
			return errors.New("all fields are required")
		}
	}
	return nil
}

func CheckUsernameFormat(username string) (bool, error) {
	if strings.TrimSpace(username) != "" {
		valid, err := regexp.MatchString(`(?i)^[a-z0-9]{3,21}$`, username)
		if err != nil || !valid {
			return false, err
		}
	}
	return true, nil
}

func CheckEmailFormat(email string) (bool, error) {
	if len(email) > 60 {
		return false, nil
	}

	isValid, err := regexp.MatchString(`(?i)^[a-z0-9]+\.?[a-z0-9]+@[a-z0-9]+\.[a-z]+$`, email)
	if err != nil {
		return false, err
	} else if !isValid {
		return false, nil
	}
	return true, nil
}

func CheckAge(BirthDate string) (int, error) {
	regex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !regex.MatchString(BirthDate) {
		return -1, errors.New("invalid date format: expected YYYY-MM-DD")
	}
	year, err := strconv.Atoi(BirthDate[0:4])
	if err != nil {
		return -1, errors.New("invalid year")
	}
	month, err := strconv.Atoi(BirthDate[5:7])
	if err != nil {
		return -1, errors.New("invalid month")
	}
	day, err := strconv.Atoi(BirthDate[8:10])
	if err != nil {
		return -1, errors.New("invalid day")
	}

	currentTime := time.Now()
	currentYear := currentTime.Year()
	currentMonth := int(currentTime.Month())
	currentDay := currentTime.Day()

	userAge := currentYear - year
	if currentMonth < month || (currentMonth == month && currentDay < day) {
		userAge--
	}

	if userAge < 16 {
		return -1, nil
	}
	return userAge, nil
}

func CheckPasswordFormat(password string) bool {
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	isSpecial := regexp.MustCompile(`[^\w\s]`)
	isLower := regexp.MustCompile(`[a-z]`)
	isUpper := regexp.MustCompile(`[A-Z]`)
	isDigit := regexp.MustCompile(`[0-9]`)
	return isLower.MatchString(password) && isUpper.MatchString(password) && isDigit.MatchString(password) && isSpecial.MatchString(password)
}
