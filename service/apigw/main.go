package main

import (
	"CloudStorage/service/apigw/route"
)

func main() {
	r := route.Router()
	r.Run(":8080")
}
