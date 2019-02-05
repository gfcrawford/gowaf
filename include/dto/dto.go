package dto

import (
	"log"
	"os"
)

type dto struct {
	id     int
	status string
}

func init() {
	f, err := os.OpenFile("/var/log/gowaf/dto.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

}
