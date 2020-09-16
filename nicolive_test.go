package nicolive

import (
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	_, err := New(os.Getenv("MAIL"), os.Getenv("PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestListen(t *testing.T) {
	if os.Getenv("CI") != "" {
		return
	}

	n, _ := New(os.Getenv("MAIL"), os.Getenv("PASSWORD"))
	err := n.Listen(os.Getenv("LIVE_ID"), func(c *Chat) error {
		fmt.Println(c.Text)
		return fmt.Errorf("ERROR_FOR_EXIT")
	})

	if err.Error() != "ERROR_FOR_EXIT" {
		t.Fatal(err)
	}
}
