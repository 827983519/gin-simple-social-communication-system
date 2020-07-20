package main

import (
	"sbs-entrytask-template/apps/example/routers"
)

func main() {
	r := routers.SetupRouter()

	r.Run()
}
