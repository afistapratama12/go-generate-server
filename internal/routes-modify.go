package internal

import (
	"log"
	"regexp"
)

func ModifyRoutesRoot() {
	read := OpenFile("routes/root.go")

	regex, err := regexp.Compile("DB = config[.]ConnectDB[(][)]")
	if err != nil {
		log.Fatal(err)
		return
	}

	newRead := regex.ReplaceAllString(read, "DB = config.ConnectDB()\n\tauthService = auth.NewAuthService()\n\tMainMiddleware = middleware.Middleware(authService)")

	WriteNewMain("routes/root.go", newRead, "routes/root.go")
}
