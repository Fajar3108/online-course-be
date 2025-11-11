package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/Fajar3108/online-course-be/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func getDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config().Database.User,
		config.Config().Database.Pass,
		config.Config().Database.Host,
		config.Config().Database.Port,
		config.Config().Database.Name,
	)
}

func DB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(getDsn()), &gorm.Config{TranslateError: true})
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})

	return db
}
