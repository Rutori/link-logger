package db

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"link-logger/db/models"
)

const dbFileName = "Links.db"

// databaseConn database connection
var databaseConn *gorm.DB

func Init() (err error) {
	defer func() {
		if err != nil {
			return
		}

		err = databaseConn.AutoMigrate(&models.Links{})
	}()
	databaseConn, err = gorm.Open(sqlite.Open(dbFileName), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Silent, // Log level
			}),
	})
	if err == nil {
		return nil
	}

	// DB missing
	err = createDB()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createDB() error {
	file, err := os.Create(dbFileName)
	if err != nil {
		return errors.WithStack(err)
	}

	err = file.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
