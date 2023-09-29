package service

import (
	"errors"
	"github.com/ms-clean-code/external/fakeapi"
	"github.com/ms-clean-code/infras/log"
	errSvc "github.com/ms-clean-code/internal/domain/error"
	"github.com/ms-clean-code/internal/domain/user/model/dto"
	"github.com/ms-clean-code/internal/domain/user/repository"
	"github.com/ms-clean-code/shared"
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

	errC := request.ValidateRequest()
	switch errC.TypeErr.Error() {
	case "required":
		err = shared.MakeError(errSvc.InvalidMandatory, errC.Field)
	default:
		err = shared.MakeError(errSvc.InvalidFormat, errC.Field)
	}

	dataDb, err := request.DtoRequest()
	if err != nil {
		err = shared.MakeError(errSvc.BadRequest)
		return
	}

	reqFA := request.DtoFakeApi()

	_, err = us.fa.AddNewProductFakeApi(reqFA)
	if err != nil {
		return
	}

	data, err := us.ur.InsertDataUser(dataDb)
	if err != nil {
		err = shared.MakeError(errSvc.InternalServerError)
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

func (us UserService) ResovleUserByName(req dto.UserListRequest) (response dto.ListUser, err error) {
	filter := req.ToFilter()

	users, err := us.ur.List(filter)
	if err != nil {
		return
	}

	if len(users) == 0 {
		err = errors.New("not found")
		return
	}

	response = dto.NewResponseUser(users)

	return
}
