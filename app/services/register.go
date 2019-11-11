package services

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/app/utils"
	"gopkg.in/go-playground/validator.v9"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
}

type RegisterResponse struct {
	User        models.User `json:"user"`
	AccessToken string      `json:"access_token"`
}

func Register(r RegisterRequest, db *gorm.DB) (RegisterResponse, error) {
	v := validator.New()
	err := v.Struct(r)
	var result error

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			result = multierror.Append(result, errors.New(utils.FormatErrors(e.ActualTag(), e.Field(), e.Param())))
		}

		return RegisterResponse{}, result
	}

	var user models.User
	db.Where("email = ?", r.Email).Find(&user)
	if user.ID != 0 {
		return RegisterResponse{}, errors.New("Email is already taken.")
	}

	user = models.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  HashString(r.Password),
	}

	db.Create(&user)

	userToken := EncodeUserInfo(user)
	return RegisterResponse{user, userToken}, nil
}
