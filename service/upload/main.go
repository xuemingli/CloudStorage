package main

import (
	cfg "CloudStorage/config"
	"CloudStorage/route"
)

func main() {
	// gin framework
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}
