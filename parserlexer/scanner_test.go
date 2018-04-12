package parserlexer

import (
	"strings"
	"testing"
)

func TestReadCharacter(t *testing.T) {
	snippet := `http {
  }`
	rdr := strings.NewReader(snippet)
	scnr := NewScanner(rdr)
	rne := scnr.read()

	if string(rne) != "h" {
		t.Error("Rune mismatch")
	}
}
