package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dqn/nicolive"
)

func run() error {
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage:\n  nicolive <mail> <password> <live-id>")
	}

	if flag.NArg() != 3 {
		flag.Usage()
		os.Exit(2)
	}

	var (
		mail     = flag.Arg(0)
		password = flag.Arg(1)
		liveID   = flag.Arg(2)
	)

	n, err := nicolive.New(mail, password)
	if err != nil {
		return err
	}

	n.Listen(liveID, func(c *nicolive.Chat) error {
		fmt.Println(c.Text)
		return nil
	})

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
