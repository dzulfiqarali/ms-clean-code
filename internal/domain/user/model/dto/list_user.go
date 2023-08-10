package dto

import (
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"
	"github.com/guregu/null"
	"strconv"
)

type UserListRequest struct {
	Nama       null.String `json:"nama" example:"dzul"`
	Alamat     null.String `json:"alamat" example:"bogor"`
	Pendidikan null.String `json:"pendidikan" example:"STM"`
	Page       null.String `json:"page" example:"1"`
	Size       null.String `json:"size" example:"50"`
}

func (u *UserListRequest) ToFilter() model.Filter {
	var (
		page int = 1
		size int = 10
	)

	if u.Page.String != "" {
		page, _ = strconv.Atoi(u.Page.ValueOrZero())
	}
	if u.Size.String != "" {
		size, _ = strconv.Atoi(u.Size.ValueOrZero())
	}

	filter := model.Filter{
		Pagination: model.Pagination{
			Page:     page,
			PageSize: size,
		},
		Sorts: []model.Sort{
			{
				Field: "created_at",
				Order: model.SortDesc,
			},
		},
	}

	if u.Nama.String != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "nama",
			Operator: model.OperatorEqual,
			Value:    u.Nama.ValueOrZero(),
		})
	}
	if u.Alamat.String != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "alamat",
			Operator: model.OperatorEqual,
			Value:    u.Alamat.ValueOrZero(),
		})
	}

	if u.Pendidikan.String != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "pendidikan",
			Operator: model.OperatorEqual,
			Value:    u.Pendidikan.ValueOrZero(),
		})
	}

	return filter
}
