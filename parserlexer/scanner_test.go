package parserlexer

import (
	"strings"
	"testing"
)

func TestReadCharacter(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	rne := scnr.read()

	if string(rne) != "h" {
		t.Error("Rune mismatch")
	}
}

func TestUnreadCharacter(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	scnr.read()
	scnr.unread()
	rne := scnr.read()

	if string(rne) != "h" {
		t.Error("Rune mismatch")
	}
}
