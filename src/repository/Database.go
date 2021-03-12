package repository

import (
	"CoreService/src/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Base struct {
	ID			uuid.UUID	`gorm:"type:uuid;primary_key;unique;default:uuid_generate_v4()"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:"index"`
}

var db *gorm.DB

func GetDatabase() *gorm.DB {
	if db == nil {
		SetupConnection()
	}
	return db
}

func CreateDatabases()  {
	db.AutoMigrate(&Registry{})
}

func SetupConnection()  {
	var err error
	db, err = setup()
	if err != nil {
		util.Logger().Fatal(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		util.Logger().Fatal(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	err = sqlDB.Ping()
	if err != nil {
		util.Logger().Fatal(err.Error())
	}

	util.Logger().Info("Connection with database successful")

}

func setup() (*gorm.DB, error) {
	database := util.GetConfig().Database
	dbType := strings.ToLower(database.Type)

	if dbType == "" {
		return nil, errors.New("database: there is not database type given")
	}

	if dbType == "postgres" || dbType == "postgresql" {

		return setupPostgresql(database.Host, database.Port, database.Username, database.Password, database.Name)
	}

	if dbType == "mysql" || dbType == "mariadb" {
		return setupMysql(database.Host, database.Port, database.Username, database.Password, database.Name)
	}

	return nil, fmt.Errorf("database: there is no database with given type %s", dbType)
}

func setupPostgresql(host string, port int, username string, password string, name string) (*gorm.DB, error) {
	util.Logger().Info("Setting up connection with PostgreSQL database")

	dsn := "host=" + host +
		" port=" + strconv.Itoa(port) +
		" user=" + username +
		" password=" + password +
		" database=" + name +
		" sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: 					dsn,
		PreferSimpleProtocol: 	false,
		WithoutReturning:     	false,
	}), &gorm.Config{})

	return db, err
}


func setupMysql(host string, port int, username string, password string, name string) (*gorm.DB, error) {
	util.Logger().Info("Setting up connection with MySQL database")

	dsn := username + ":" +
		password + "@tcp(" +
		host + ":" +
		strconv.Itoa(port) + ")/" +
		name +
		"?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	return db, err
}