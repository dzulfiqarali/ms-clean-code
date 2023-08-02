package repository

import (
	"bitbucket.org/bridce/ms-clean-code/infras/database"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"
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
