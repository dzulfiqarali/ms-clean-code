package database

import (
	"bitbucket.org/bridce/ms-clean-code/configs"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Conn struct {
	Read  *gorm.DB
	Write *gorm.DB
}

// ProvideConn is the provider for db connection.
func ProvideConn(config *configs.Config) *Conn {
	return &Conn{
		Read:  ConnectDatabaseRead(config),
		Write: ConnectDatabaseWrite(config),
	}
}

func ConnectDatabaseRead(configs *configs.Config) *gorm.DB {
	dsn := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s",
		configs.Database.Database,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Username,
		configs.Database.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connecting to database with error: " + err.Error())
	}

	err = db.Debug().AutoMigrate(&model.User{})
	if err != nil {
		return db
	}

	return db
}

func ConnectDatabaseWrite(configs *configs.Config) *gorm.DB {
	dsn := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s",
		configs.Database.Database,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Username,
		configs.Database.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connecting to database with error: " + err.Error())
	}

	//err = db.AutoMigrate(&file_model.File{})
	//if err != nil {
	//	return nil, err
	//}

	return db
}
