package helpers

import (
	"log"
)

// has to be CapitalCase to be publicly available
func RaiseError(err error, message string) {
	// Raise error with message and display error
	if err != nil && message != "" {
		log.Fatal("Failed to create table:", err)
	} else if err != nil {
		panic(err)
	}
}
