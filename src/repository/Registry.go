package repository

import (
	"CoreService/src/util"
	"fmt"
	"time"
)

type Registry struct {
	Base
	Name 		string
	LastUsed	time.Time
}

func CreateRegistry()  {
	registry := Registry{Name: "NGINX"}
	db := GetDatabase()

	result := db.Create(&registry)

	util.Logger().Info(fmt.Sprintf("Rows affected: %d", result.RowsAffected))
}