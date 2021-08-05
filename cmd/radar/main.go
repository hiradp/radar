package main

import (
	"log"
	"os"

	"github.com/hiradp/radar/pkg/cmd"
)

func main() {
	f, err := os.OpenFile("logs/radar.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	cmd.Execute()
}
