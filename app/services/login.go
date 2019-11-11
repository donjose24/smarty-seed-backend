package services

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/app/utils"
	"gopkg.in/go-playground/validator.v9"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User        models.User `json:"user"`
	AccessToken string      `json:"access_token"`
}

func Login(r LoginRequest, db *gorm.DB) (LoginResponse, error) {
	v := validator.New()
	err := v.Struct(r)
	var result error

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			result = multierror.Append(result, errors.New(utils.FormatErrors(e.ActualTag(), e.Field(), e.Param())))
		}

		return LoginResponse{}, result
	}

	// TEMPORARILY THIS. HAHA DON"T DO THIS.
	// i believe there are better ways on authentcating the user. I just don't have time to search for it
	user := models.User{}
	db.Where("email = ?", r.Email).Find(&user)
	if user.FirstName == "" {
		return LoginResponse{}, errors.New("Invalid Credentials. Please check your email and password")
	}
	err = CompareToHash(user.Password, r.Password)

	if err != nil {
		fmt.Println(r.Password)
		return LoginResponse{}, errors.New("Invalid Credentials. Please check your email and password")
	}

	userToken := EncodeUserInfo(user)
	return LoginResponse{User: user, AccessToken: userToken}, nil
}
