package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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

	shares, err := client.ListShares()
	if err != nil {
		log.Fatal(err)
	}

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 5, 0, 2, ' ', 0)
	for _, share := range shares {
		fmt.Fprintf(w, "%s\t%s\t%s\n", share.Name, share.Encryption, share.Description)
	}
	w.Flush()
}
