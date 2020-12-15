package env

import (
	"log"
	"os"
)

func get(key, fallback string) string {
	if value := os.Getenv(key); len(value) > 0 {
		return value
	}
	log.Printf("%s not found. Falling back to \"%s\".\n", key, fallback)
	return fallback
}

func mustGet(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panic("failed to get variable: " + key)
	}
	return value
}
