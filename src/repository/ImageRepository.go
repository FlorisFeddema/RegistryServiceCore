package repository

import (
	"CoreService/src/util"
	"fmt"
)

type Image struct {
	Base
	Name 	string
}

func CreateImage()  {
	image := Image{Name: "NGINX"}
	db := GetDatabase()

	result := db.Create(&image)

	util.Logger().Info(fmt.Sprintf("Rows affected: %d", result.RowsAffected))
}