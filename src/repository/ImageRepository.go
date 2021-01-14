package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Image struct {
	Id		uuid.UUID
	Name 	string
}

func Connect()  {
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost port=5432 user=cbE8Eg2ukBD7u6yxuhXkqZb8vCS3bhCN password=29q4wSumbF3QG2wEaAtynKNtvcEf4nfc database=CoreService sslmode=disable",
	}), &gorm.Config{})

	fmt.Printf(db.Name())
}