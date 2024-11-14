package connection

import (
	"fmt"
	osutil "gin-seed/app/common/os"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var db *gorm.DB

func GetConnection() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		var err error
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s",
			os.Getenv("DB_HOST"),
			osutil.DefaultGetIEnv("DB_PORT", 5432),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("Cannot connect to database")
		}

		if sqlDb, err := db.DB(); err != nil {
			osutil.DoGetIEnv("DB_MAX_IDLE_CONNS", sqlDb.SetMaxIdleConns)
			osutil.DoGetIEnv("DB_MAX_OPEN_CONNS", sqlDb.SetMaxOpenConns)
		}
	}

	return db
}
