package service

import (
	"bitbucket.org/bridce/ms-clean-code/external/fakeapi"
	"bitbucket.org/bridce/ms-clean-code/infras/log"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model/dto"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/repository"
)

type UserService struct {
	ur repository.UserRepoInterface
	fa fakeapi.FakeApiImpl
	l  *log.LogCustom
}

func UserInterface(ur repository.UserRepoInterface, fa fakeapi.FakeApiImpl, l *log.LogCustom) UserServiceInterface {
	return &UserService{
		ur: ur,
		fa: fa,
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

	reqFA := request.DtoFakeApi()

	_, err = us.fa.AddNewProductFakeApi(reqFA)
	if err != nil {
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

func (us UserService) ResovleListUserByFilter(req dto.UserListRequest) (response dto.ResponseListUser) {
	filter := req.ToFilter()

	users, err := us.ur.List(filter)
	if err != nil {
		return
	}

	response = dto.NewResponseListUser(users, filter)

	return
}
