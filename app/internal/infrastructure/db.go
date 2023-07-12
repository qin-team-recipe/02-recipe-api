package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(c *config.Config) *DB {
	return &DB{
		Connection: connection(
			c.DBHost,     // host
			c.DBUsername, // user name
			c.DBPassword, // password
			c.DBName,     // database name
		),
	}
}

func connection(host, username, password, dbName string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Nanosecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	count := 0

	conn, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		for {
			if err == nil {
				break
			}
			fmt.Print(".\n")
			time.Sleep(time.Second)
			count++
			// connection wait 10 seconds for database starting...
			if count > 10 {
				fmt.Print("database connection failed\n")
				panic(err.Error())
			}
			conn, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{})
		}
	}

	log.Print("database connection success\n")

	return conn
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}
