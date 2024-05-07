package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiart-tantasi/env"
)

func main() {
	// env
	environment := os.Getenv("ENV")
	projectRoot := os.Getenv("PROJECT_ROOT")
	env.LoadEnvFile(environment, projectRoot)
	// test github.com/kiart-tantasi/env
	fmt.Println("...FOO env:", os.Getenv("FOO"), "...")
	//routing
	r := gin.Default()
	r.GET("/api/health", func(c *gin.Context) {
		time.Sleep(3000 * time.Millisecond)
		c.Data(200, "", nil)
	})
	// run
	r.Run()
}
