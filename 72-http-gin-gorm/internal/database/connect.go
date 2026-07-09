package database

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	RETRY_TIMES    = 5
	RETRY_DURATION = time.Second * 5
)

// /go:generate docker ps
func Connect(dsn string) (db *gorm.DB, err error) {
	for i := range RETRY_TIMES {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			slog.Error(fmt.Sprint(err.Error(), "Trying to connect to the database--> ", i+1, " times"))
			//return nil, err
		} else {
			return db, nil
		}
		time.Sleep(RETRY_DURATION) // Wait

	}
	return nil, err
}
