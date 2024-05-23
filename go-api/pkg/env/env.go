package env

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvFile(environment, projectRoot string) {
	path := ""
	if environment == "production" {
		// production
		path = filepath.Join(projectRoot, ".env.production")
		// [why we need to set PROJECT_ROOT in production and not in development]
		// production app is running on systemd and current directory will be "/"
		// so we need to tell the app where env file is located
		// on the other hand, running app at `go-api` in development is already in correct directory
	} else {
		// development and others
		path = filepath.Join(projectRoot, ".env")
		environment = "development"
	}
	err := godotenv.Load(path)
	if err != nil {
		panic(err)
	}
	log.Printf("loaded env file for environment \"%s\"\n", environment)
}
