package nicolive

import (
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
	n, _ := New(os.Getenv("MAIL"), os.Getenv("PASSWORD"))
	err := n.Listen(os.Getenv("LIVE_ID"))
	if err != nil {
		t.Fatal(err)
	}
}
