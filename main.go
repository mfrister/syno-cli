package main

import (
	"log"
	"os"

	"frister.net/playground/synoshares/synoapi"
)

func main() {
	// something like https://myds.example.net:5001
	api_base := os.Getenv("SYNO_BASE_URL")
	user := os.Getenv("SYNO_USER")
	password := os.Getenv("SYNO_PASSWORD")

	client := synoapi.NewClient(api_base)
	err := client.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}
}
