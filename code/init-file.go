package code

var DotEnv = `# This file is used to configure the application.
DB_USER=
DB_PASS=
DB_HOST=
DB_NAME=
DB_PORT=

# jwt config
JWT_SECRET=
`

var ExampleEnv = `# This file is used to configure the application.
DB_USER=root
DB_PASS=admin
DB_HOST=localhost
DB_NAME=example_db
DB_PORT=3306

# jwt config
JWT_SECRET=example-secret
`

var Gitignore = `# ignore all following files
.env
/node_modules
/public
/storage
`
