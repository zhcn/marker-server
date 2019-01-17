package main

import (
	"market-server/server/routes"
)

func main() {
	routes.Init()
	routes.Router.Run(":8000")
}
