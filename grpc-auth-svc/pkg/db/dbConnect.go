package db

import (
	"fmt"
	"log"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config config.Config) *gorm.DB {
	connstr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.Db.DbHost, config.Db.DbUser, config.Db.DbName, config.Db.DbPort, config.Db.DbPaswword)
	db, dberr := gorm.Open(postgres.Open(connstr), &gorm.Config{})

	if dberr != nil {
		log.Fatal("Failed to connect database")
		return nil
	}

	fmt.Println("Connected succesfully....!")

	return db

}
