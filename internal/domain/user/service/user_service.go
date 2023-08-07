package service

import (
	"bitbucket.org/bridce/ms-clean-code/infras/log"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model/dto"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/repository"
)

type UserService struct {
	ur repository.UserRepoInterface
	l  *log.LogCustom
}

func UserInterface(ur repository.UserRepoInterface, l *log.LogCustom) UserServiceInterface {
	return &UserService{
		ur: ur,
		l:  l,
	}
}

func (us UserService) RegistrationUser(request dto.RegistUserRequest) (resp dto.RegistUserResponse, err error) {

	defer func() {
		if err != nil {
			us.l.Error(err, "service : RegistrationUser", "", nil, request, err, nil, nil)
		}
		us.l.Success(request, resp, nil, nil, "service : RegistrationUser", "nil", nil)
	}()

	dataDb, err := request.DtoRequest()
	if err != nil {
		//
		return
	}

	data, err := us.ur.InsertDataUser(dataDb)
	if err != nil {
		//
		return
	}

	resp = dto.DtoResponseUser(*data)
	return
}
