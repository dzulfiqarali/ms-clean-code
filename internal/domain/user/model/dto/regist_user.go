package dto

import (
	"bitbucket.org/bridce/ms-clean-code/external/fakeapi"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"
	"github.com/go-playground/validator/v10"
)

type RegistUserRequest struct {
	Nama       string `json:"nama" validate:"required,alphaunicode"`
	Alamat     string `json:"alamat" validate:"required,alphanumunicode"`
	Umur       string `json:"umur" validate:"required,numeric"`
	Pendidikan string `json:"pendidikan" validate:"required,alphanumunicode"`
}

func (u *RegistUserRequest) DtoRequest() (model.User, error) {

	err := u.ValidateRequest()
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Nama:       u.Nama,
		Alamat:     u.Alamat,
		Umur:       u.Umur,
		Pendidikan: u.Pendidikan,
	}

	return user, nil
}

func (u *RegistUserRequest) ValidateRequest() error {

	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
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
