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
