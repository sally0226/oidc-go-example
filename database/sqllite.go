package database

import (
	"github.com/sally0226/oidc-go-example/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func MustInitSQLlite() *gorm.DB {
	//db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}) // in-memory DB
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{}) // file DB
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}

	return db
}

func MustMigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	log.Println("Database migrated successfully")
}
