package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	env := os.Getenv("ENV")
	path := ""
	if env == "production" {
		// production
		path = filepath.Join(os.Getenv("PROJECT_ROOT"), ".env.production")
		// [why we need to use PROJECT_ROOT ?]
		// production app is running on systemd so current directory will be /
		// so we need to tell the app where project or env file is located
	} else {
		// development and others
		env = "development"
		path = ".env"
	}
	err := godotenv.Load(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("loaded env file for environment \"%s\"\n", env)
}
