package util

import (
	"github.com/google/uuid"
	"log"
)

func NewUUID() string {
	return uuid.New().String()
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
