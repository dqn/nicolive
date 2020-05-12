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
