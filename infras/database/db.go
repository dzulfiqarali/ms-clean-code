package database

import (
	"bitbucket.org/bridce/ms-clean-code/configs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s",
		configs.AppConfig.Database.Database,
		configs.AppConfig.Database.Host,
		configs.AppConfig.Database.Port,
		configs.AppConfig.Database.Username,
		configs.AppConfig.Database.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//err = db.AutoMigrate(&file_model.File{})
	//if err != nil {
	//	return nil, err
	//}

	return db, nil
}
