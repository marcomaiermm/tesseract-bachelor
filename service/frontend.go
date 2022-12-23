package service

import (
	"log"
	"os"
	"os/exec"
)

// BuildFrontendService builds the admin svelte app from the ./admin folder and copies the content to the ./public/admin folder
func BuildFrontendService() {
	// get path of base directory
	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	isYarn := true
	cmd := exec.Command("yarn", "install")
	if _, err := os.Stat(baseDir + "/frontend/yarn.lock"); err != nil {
		isYarn = false
	}

	// if no node_modules folder is found, install the dependencies
	if _, err := os.Stat(baseDir + "/frontend/node_modules"); os.IsNotExist(err) {
		log.Println("Installing frontend dependencies...")
		if !isYarn {
			cmd = exec.Command("npm", "install")
		}
		cmd.Dir = baseDir + "/frontend"
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	// build the frontend app
	// if yarn is used, use yarn build command. Else use npm run build
	log.Println("Building frontend app...")

	cmd = exec.Command("yarn", "build")
	if !isYarn {
		cmd = exec.Command("npm", "run", "build")
	}
	cmd.Dir = baseDir + "/frontend"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
