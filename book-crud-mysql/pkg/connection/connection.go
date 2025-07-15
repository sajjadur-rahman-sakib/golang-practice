package connection

import (
	"book-crud/pkg/config"
	"book-crud/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func connect() {
	dbConfig := config.LocalConfig

	dsn := fmt.
		Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
			dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIP, dbConfig.DbName)

	fmt.Printf("Attempting to connect to database with DSN: %s\n",
		fmt.Sprintf("%s:***@%s/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
			dbConfig.DBUser, dbConfig.DBIP, dbConfig.DbName))

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		log.Printf("Please ensure MySQL is running on %s and database '%s' exists", dbConfig.DBIP, dbConfig.DbName)
		panic(err)
	}
	fmt.Println("Database Connected Successfully")
	db = d
}

func migrate() {
	db.Migrator().AutoMigrate(&models.Book{})
}

func GetDB() *gorm.DB {
	if db == nil {
		connect()
	}
	migrate()
	return db
}
