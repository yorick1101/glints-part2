package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"glints-part2/config"
	"glints-part2/etl"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	if len(os.Args) != 2 {
		log.Panic("etl [source_file]")
	}

	filePath := os.Args[1]
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.IsDir() {
		log.Panic("file not found or not a file", filePath, err)
	}

	config.Init()
	etl.NewService().Process(filePath)
}
