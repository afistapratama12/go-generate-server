package code

// package form connect DB
const (
	DriverMysql    = "gorm.io/driver/mysql"
	DriverPostgres = "gorm.io/driver"
	DriverSqlite   = "gorm.io/driver"
)

// format connection to driver db
const (
	DBMysql    = "%s:%s@tcp(%s:%s)/%s"
	DBPostgres = ""
	DBSqlite   = ""
)

var initConfig = `package config

import (
	"%s"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     string
	SSL      string
}

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}


func ConnectDB() *gorm.DB {
	var cred Config

	cred.Username = os.Getenv("DB_USER")
	cred.Password = os.Getenv("DB_PASS")
	cred.Host = os.Getenv("DB_HOST")
	cred.DBName = os.Getenv("DB_NAME")
	cred.Port = os.Getenv("PORT")



}

`

func AddInitConfig() string {
	return initConfig
}
