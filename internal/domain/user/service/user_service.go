package service

import (
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model/dto"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/repository"
)

type UserService struct {
	ur repository.UserRepoInterface
}

func UserInterface(ur repository.UserRepoInterface) UserServiceInterface {
	return &UserService{ur}
}

func (us UserService) RegistrationUser(request dto.RegistUserRequest) (resp dto.RegistUserResponse, err error) {

	dataDb, err := request.DtoRequest()
	if err != nil {
		return
	}

	data, err := us.ur.InsertDataUser(dataDb)
	if err != nil {
		return
	}

	resp = dto.DtoResponseUser(data)
	return
}
