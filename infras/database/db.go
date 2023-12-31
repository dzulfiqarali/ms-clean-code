package database

import (
	"fmt"
	"github.com/ms-clean-code/configs"
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

	return db
}
