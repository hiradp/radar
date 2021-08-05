package main

import (
	"log"
	"os"

	"github.com/hiradp/radar/pkg/cmd"
)

func main() {
	f, err := os.OpenFile("radar.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	cmd.Execute()
}
