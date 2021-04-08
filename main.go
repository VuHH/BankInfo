package main

import (
	"./routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
