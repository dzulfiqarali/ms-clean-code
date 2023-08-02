package repository

import (
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"
	"gorm.io/gorm"
)

type UserRepoStrct struct {
	db *gorm.DB
}

func UserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserRepoStrct{db}
}

func (ur UserRepoStrct) InsertDataUser(u model.User) (model.User, error) {
	tx := ur.db.Begin()

	err := tx.Debug().Create(&u).Error
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	tx.Commit()
	return u, nil
}
