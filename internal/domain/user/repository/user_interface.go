package repository

import "github.com/ms-clean-code/internal/domain/user/model"

type UserRepoInterface interface {
	InsertDataUser(u model.User) (*model.User, error)
	List(filter model.Filter) (result []model.ListUser, err error)
}
