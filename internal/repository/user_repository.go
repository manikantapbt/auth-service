package repository

import (
	"auth-service/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	INSERT_QUERY = `
		INSERT INTO users (name,user_name, email, is_verified, country_code, phone_number)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`
	GET_QUERY       = "SELECT * FROM users WHERE id = $1"
	GET_USER_BY_PH  = "SELECT id, name, email, is_verified, country_code, phone_number FROM users WHERE country_code = $1 AND phone_number = $2"
	UPDATE_VERIFIED = "UPDATE users SET is_verified = true WHERE id = $1"
)

type IUserRepository interface {
	SaveUser(user *models.User) (*models.User, error)
	GetUser(userId int32) (*models.User, error)
	GetUserByPhoneNumberAndCountry(countryCode int32, phoneNumber string) (*models.User, error)
	UpdateVerifiedTrueById(id int32) error
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &psqlUserRepository{db: db}
}

type psqlUserRepository struct {
	db *sql.DB
}

func (p *psqlUserRepository) SaveUser(user *models.User) (*models.User, error) {
	var id int32
	err := p.db.QueryRow(INSERT_QUERY, user.Name, user.UserName, user.Email, user.Verified, user.CountryCode, user.PhoneNumber).Scan(&id)
	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}

func (p *psqlUserRepository) GetUser(userId int32) (*models.User, error) {
	var user models.User
	err := p.db.QueryRow(GET_QUERY, userId).Scan(&user.Id, &user.Name, &user.UserName, &user.Email, &user.Verified, &user.CountryCode, &user.PhoneNumber, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *psqlUserRepository) GetUserByPhoneNumberAndCountry(countryCode int32, phoneNumber string) (*models.User, error) {
	var user models.User
	err := p.db.QueryRow(GET_USER_BY_PH, countryCode, phoneNumber).Scan(&user.Id, &user.Name, &user.Email, &user.Verified, &user.CountryCode, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with country code %d and phone number %s not found", countryCode, phoneNumber)
		}
		return nil, err
	}

	return &user, nil
}

func (p *psqlUserRepository) UpdateVerifiedTrueById(id int32) error {
	_, err := p.db.Exec(UPDATE_VERIFIED, id)
	if err != nil {
		return err
	}
	return nil
}
