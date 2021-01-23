package repository

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Image struct {
	Id		uuid.UUID
	Name 	string
}

func Connect()  {
	//db, _ := gorm.Open(postgres.New(postgres.Config{
	//	DSN: "host=localhost port=5432 user=cbE8Eg2ukBD7u6yxuhXkqZb8vCS3bhCN password=29q4wSumbF3QG2wEaAtynKNtvcEf4nfc database=CoreService sslmode=disable",
	//}), &gorm.Config{})

	var db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: "Q2TXf8UyR36Jy4CmRK4c6sGx9rKnEkuj:wd8HSPCX7j9LXFdzUvYZL8zqg3xgV2ET@tcp(localhost:3306)/CoreService?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	_ = sqlDB.Ping()

	sqlDB.Stats()
}