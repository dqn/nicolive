package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dqn/nicolive"
)

func usage() {
	fmt.Println("Usage:\n  nicolive <mail> <password> <live-id>")
}

func run() error {
	if len(os.Args) != 4 {
		usage()
		os.Exit(2)
	}

	var (
		mail     = os.Args[1]
		password = os.Args[2]
		liveID   = os.Args[3]
	)

	n, err := nicolive.New(mail, password)
	if err != nil {
		return err
	}

	n.Listen(liveID, func(c *nicolive.Chat) {
		fmt.Println(c.Text)
	})

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
