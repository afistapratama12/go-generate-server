package code

// package form connect DB
const (
	DriverMysql    = "gorm.io/driver/mysql"
	DriverPostgres = "gorm.io/driver"
	DriverSqlite   = "gorm.io/driver"
)

// format connection to driver db
const (
	DBMysql    = ""
	DBPostgres = ""
	DBSqlite   = ""
)

var InitConfig = `package config

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


func ConnectDB() *gorm.DB {

}

`
