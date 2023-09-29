package dto

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/ms-clean-code/external/fakeapi"
	"github.com/ms-clean-code/internal/domain/user/model"
	"strings"
)

type RegistUserRequest struct {
	Nama       string `json:"nama" validate:"required,alphaunicode"`
	Alamat     string `json:"alamat" validate:"required,alphanumunicode"`
	Umur       string `json:"umur" validate:"required,numeric"`
	Pendidikan string `json:"pendidikan" validate:"required,alphanumunicode"`
}

type ErrorCustom struct {
	Field   string
	TypeErr error
}

func (u *RegistUserRequest) DtoRequest() (model.User, error) {

	user := model.User{
		Nama:       u.Nama,
		Alamat:     u.Alamat,
		Umur:       u.Umur,
		Pendidikan: u.Pendidikan,
	}

	return user, nil
}

func (u *RegistUserRequest) ValidateRequest() (errC ErrorCustom) {

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		fieldError := err.(validator.ValidationErrors)
		errC = parsingError(strings.Split(fieldError.Error(), " ")[5], strings.Split(fieldError.Error(), " ")[9])
		return
	}

	return
}

func parsingError(str ...string) (err ErrorCustom) {
	for i, s := range str {
		str[i] = strings.ReplaceAll(s, "'", "")
	}
	err.Field = str[0]
	err.TypeErr = errors.New(str[1])

	return
}

func (u *RegistUserRequest) DtoFakeApi() fakeapi.RequestFakeAPI {

	return fakeapi.RequestFakeAPI{
		Tittle:      "test product",
		Price:       13.5,
		Description: "test",
		Image:       "'https://i.pravatar.cc",
		Category:    "electronic",
	}
}
