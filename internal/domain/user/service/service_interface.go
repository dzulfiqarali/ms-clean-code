package service

import "github.com/ms-clean-code/internal/domain/user/model/dto"

type UserServiceInterface interface {
	RegistrationUser(request dto.RegistUserRequest) (resp dto.RegistUserResponse, err error)
	ResovleListUserByFilter(req dto.UserListRequest) (response dto.ResponseListUser)
	ResovleUserByName(req dto.UserListRequest) (response dto.ListUser, err error)
	Error(err error) *UserError
}
