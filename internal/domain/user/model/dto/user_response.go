package dto

import "bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"

type RegistUserResponse struct {
	Name string `json:"name"`
}

func (rur *RegistUserResponse) DtoResponseUser(data model.User) {

	rur.Name = data.Nama
}
