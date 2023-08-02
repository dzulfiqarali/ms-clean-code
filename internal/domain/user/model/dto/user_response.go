package dto

import "bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"

type RegistUserResponse struct {
	Name string `json:"name"`
}

func DtoResponseUser(data *model.User) (rur RegistUserResponse) {

	rur.Name = data.Nama
	return
}
