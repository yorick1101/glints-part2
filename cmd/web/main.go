package main

import (
	"glints-part2/config"
	"glints-part2/web"
)

func main() {
	config.Init()
	engine := web.InitRouters()
	engine.Run(":8080")
}
