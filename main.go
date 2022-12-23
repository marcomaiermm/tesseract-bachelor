package main

import (
	"log"
	"os"
	"strconv"

	Core "github.com/marcomaiermm/tesseract-bachelor/core"
	Handler "github.com/marcomaiermm/tesseract-bachelor/handler"
	Service "github.com/marcomaiermm/tesseract-bachelor/service"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// main function to boot up the go server. it takes care of the routing and middleware
func main() {
	// get first arg from command line if present
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	if arg == "build" {
		Service.BuildFrontendService()
	}

	config := Core.Config

	log.Println("\nTrying to start the server on " + config.Host + ":" + config.Port + "...\n")
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
	// routes for the api

	router.GET("/api/complaint", Handler.GetAllComplaints)
	router.GET("/api/complaint/:id", Handler.GetComplaint)

	router.GET("/api/defect", Handler.GetDefectsByWeek)

	// every route which is not /api/ will be redirected to the frontend
	router.NoRoute(Handler.FrontendHandler)

	// run the server with the default port 8080 and on localhost if no port or host is provided
	// if the port is already in use, up the port number by one and try again
	for {
		err := router.Run(config.Host + ":" + config.Port)
		if err != nil {
			// convert the port string to int and increment it by one. Port can be a string because of the .env file
			port, _ := strconv.Atoi(config.Port)
			log.Println("\nPort " + config.Port + " is already in use. Trying port " + strconv.Itoa(port+1) + "\n")
			port++
			config.Port = strconv.Itoa(port)
			// try again with the new port
			continue
		} else {
			break
		}
	}
	router.Run(config.Host + ":" + config.Port)
}
