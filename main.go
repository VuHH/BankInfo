package main

import (
	"bankinfo.com/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
