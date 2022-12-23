package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	Service "github.com/marcomaiermm/tesseract-bachelor/service"
)

// FrontendHandler serves the index.html inside the svelte frontend/dist folder if its there
func FrontendHandler(c *gin.Context) {
	exists := checkIfIndexFileExists()
	// if the file exists, serve it
	if !exists {
		// build the frontend app
		Service.BuildFrontendService()
	}

	c.File("./frontend/dist/index.html")
}

func checkIfIndexFileExists() bool {
	// check if the file exists
	if _, err := os.Stat("./frontend/dist/index.html"); err == nil {
		return true
	}
	return false
}
