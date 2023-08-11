package dto

import (
	"github.com/ms-clean-code/internal/domain/user/model"
	"math"
)

type RegistUserResponse struct {
	Name string `json:"name"`
}

func DtoResponseUser(data model.User) (rur RegistUserResponse) {

	rur.Name = data.Nama
	return
}

type ListUser struct {
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	Pendidikan string `json:"pendidikan"`
}

type ResponseListUser struct {
	Data     []ListUser `json:"data"`
	Metadata Metadata   `json:"metadata"`
}

func NewResponseListUser(result []model.ListUser, filter model.Filter) (resp ResponseListUser) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}

	resp.Data = NewResponseListUserFromFilter(result, filter)

	return
}

func NewResponseListUserFromFilter(result []model.ListUser, filter model.Filter) (res []ListUser) {
	var user ListUser

	for _, each := range result {
		user.Nama = each.Nama
		user.Alamat = each.Alamat
		user.Pendidikan = each.Pendidikan
		res = append(res, user)
	}

	return
}
