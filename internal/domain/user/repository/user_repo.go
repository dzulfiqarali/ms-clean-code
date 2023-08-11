package repository

import (
	"fmt"
	"github.com/ms-clean-code/infras/database"
	"github.com/ms-clean-code/internal/domain/user/model"
	"github.com/ms-clean-code/internal/domain/user/query"
)

type UserRepoStrct struct {
	db *database.Conn
}

func UserRepoImpl(db *database.Conn) UserRepoInterface {
	return &UserRepoStrct{
		db: db,
	}
}

func (ur UserRepoStrct) InsertDataUser(u model.User) (*model.User, error) {
	tx := ur.db.Write.Begin()

	err := tx.Debug().Create(&u).Error
	if err != nil {
		tx.Rollback()
		return &model.User{}, err
	}

	tx.Commit()
	return &u, nil
}

func (ur UserRepoStrct) List(filter model.Filter) (result []model.ListUser, err error) {
	var (
		user        model.ListUser
		whereClause string
	)
	clauses, args, err := composeFilter(filter)
	if err != nil {
		return
	}

	baseQuery := fmt.Sprintf("%s %s",
		query.UserQuery.SelectListUser,
		whereClause,
	)

	if len(args) > 0 && clauses != "" {
		whereClause += clauses
	}
	rows, err := ur.db.Read.Debug().Raw(baseQuery+whereClause, args...).Rows()
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Nama, &user.Alamat, &user.Pendidikan, &user.FilterCount)
		result = append(result, user)
	}

	return
}
