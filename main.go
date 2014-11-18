package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"gopkg.in/alecthomas/kingpin.v1"

	"frister.net/playground/synoshares/synoapi"
)

var (
	list            = kingpin.Command("list", "List shares")
	lock            = kingpin.Command("lock", "Lock an encrypted volume")
	lockShareName   = lock.Arg("share name", "Name of the share to be locked").Required().String()
	unlock          = kingpin.Command("unlock", "Unlock an encrypted volume")
	unlockShareName = unlock.Arg("share name", "Name of the share to be unlocked").Required().String()
)

func main() {
	// something like https://myds.example.net:5001
	api_base := os.Getenv("SYNO_BASE_URL")
	user := os.Getenv("SYNO_USER")
	password := os.Getenv("SYNO_PASSWORD")

	command := kingpin.Parse()

	client := synoapi.NewClient(api_base)
	err := client.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	default:
		kingpin.Usage()
	case "list":
		listShares(client)
	case "lock":
		lockShare(client, *lockShareName)
	case "unlock":
		unlockShare(client, *unlockShareName)
	}
}

func listShares(client *synoapi.Client) {
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

func lockShare(client *synoapi.Client, shareName string) {
	err := client.LockShare(shareName)
	if err != nil {
		log.Fatalf("Locking failed: %v", err)
	}
}

func unlockShare(client *synoapi.Client, shareName string) {
	fmt.Fprint(os.Stderr, "Enter password (passing via stdin is also ok):\n")
	pass := readPassword()
	err := client.UnlockShare(shareName, pass)
	if err != nil {
		log.Fatalf("Unlocking failed: %v", err)
	}
}

func readPassword() string {
	reader := bufio.NewReader(os.Stdin)

	// read until newline (enter pressed) and strip that newline
	pass, err := reader.ReadString('\n')
	trimmed_pass := strings.TrimRight(pass, "\n")
	if err != nil || trimmed_pass == "" {
		log.Fatalf("Failed to read password")
	}
	return trimmed_pass
}
