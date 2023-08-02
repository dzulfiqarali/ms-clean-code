package service

import "bitbucket.org/bridce/ms-clean-code/internal/domain/user/model/dto"

type UserServiceInterface interface {
	RegistrationUser(request dto.RegistUserRequest) (resp dto.RegistUserResponse, err error)
}
