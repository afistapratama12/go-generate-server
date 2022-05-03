package optionmodel

type CreateListOpts struct {
	WithMiddlware    bool
	QueryOrderByASC  bool
	QueryOrderByDESC bool
	TemplateUser     bool
}

type InitListOpts struct {
	// --db-mysql
	// --db-postgres / --db-postgresql / --db-pg
	// --db-sqlite
	ConfigDBGenerate string `long:"db" short:"d" description:"set db connection with custom driver, list (mysql | postgres | sqlite)"`
}

// get now directory, create directory if null
// get name folder directory
// - go mod init name sesuai nama folder <nama folder>
// - go get -u github.com/gin-gonic/gin
// - go get -u gorm.io/gorm
// create config, connection with mysql (default)
// add helper custom
// db default using mysql
// --db-mysql
// --db-postgres / --db-postgresql / --db-pg
// --db-sqlite
