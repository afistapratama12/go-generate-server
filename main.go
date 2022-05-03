package main

import "go-gen-server/cmd"

// list cmd

// init using go-gen-server
// cmd: init
// - go mod init name sesuai nama folder
// - go get -u github.com/gin-gonic/gin
// - go get -u gorm.io/gorm
// add helper custom
// db default using mysql
// --db-mysql
// --db-postgres / --db-postgresql / --db-pg
// --db-sqlite

// init create middleware
// go get -u github.com/dgrijalva/jwt-go
// cmd: init-mw <props>
// props, ex : user_id:int,role:string

// create <entity-name> <props>
// entity-name, ex : users (jamak)
// props, ex : id:int,name:string,description:string,username:string,password:string
// default : id:int,created_at:string,updated_at:string
// can using --without-default-model (tanpa default props) | alias: --wdm
// helper custom
// --with-wm / --with-middleware  (create api with middleware, must set the middleware first)
// --template-user (add login, register, and custom save data to db)
// --query-order-by-desc / --q-order-by-desc (get all desc with default param: created_at)
// --query-order-by-asc / --q-order-by-asc (get all asc with default param: created_at)

// P1 , custom server api

func main() {
	cmd.Execute()
}
