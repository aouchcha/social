package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"socialNetwork/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) RegisterUser(user entity.UserProfile) (int, error) {
	// this should be in service layer
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	Query := `INSERT INTO users(username,first_name,last_name,email,password,date_of_birth) VALUES(?,?,?,?,?,?);`
	if _, err := r.db.Exec(Query, user.Username, user.FirstName, user.LastName, user.Email, string(hashedPass), user.BirthDate); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (r *UserRepository) LoginUser(user entity.UserProfile) (entity.UserProfile, int, error) {
	userValid := entity.UserProfile{}
	Query := `SELECT user_id, password FROM users WHERE email = ?`
	err := r.db.QueryRow(Query, user.Email).Scan(&userValid.Id, &userValid.Password)
	if err == sql.ErrNoRows {
		return entity.UserProfile{}, http.StatusUnauthorized, errors.New("User Not Found")
	} else if err != nil {
		return entity.UserProfile{}, http.StatusInternalServerError, err
	}
    //this also should be in service layer  
	if err := bcrypt.CompareHashAndPassword([]byte(userValid.Password), []byte(user.Password)); err != nil {
		return entity.UserProfile{}, http.StatusUnauthorized, errors.New("incorrect Password")
	}

	return userValid, http.StatusOK, nil
}

func (u *UserRepository) CheckUserExist(user entity.UserProfile) error {
	var IsExist bool

	Query := "SELECT EXISTS(SELECT * FROM users u WHERE u.username = ? OR u.email = ?);"

	if err := u.db.QueryRow(Query, user.Username, user.Email).Scan(&IsExist); err != nil {
		return err
	}
	if IsExist {
		return errors.New("username or email already taken")
	}
	return nil
}
