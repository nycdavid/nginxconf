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
	scnr.read()        // Read 1st char
	scnr.unread()      // Put it back
	rne := scnr.read() // Read it again

	if string(rne) != "h" {
		t.Error("Rune mismatch")
	}
}

func TestReadFirstToken(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	tok, _ := scnr.Scan()
	if tok != HTTP {
		t.Error("Token mismatch")
	}
}

func TestReadWhitespaceToken(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	scnr.Scan()
	tok, _ := scnr.Scan()

	if tok != WS {
		t.Error("Token mismatch")
	}
}

func TestReadOpenBraceToken(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	scnr.Scan()
	scnr.Scan()
	tok, _ := scnr.Scan()

	if tok != OPEN_BRACE {
		t.Error("Token mismatch")
	}
}

func TestReadNewlineToken(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	for i := 0; i < 3; i++ {
		scnr.Scan()
	}
	tok, _ := scnr.Scan()

	if tok != WS {
		t.Error("Token mismatch")
	}
}

func TestReadCloseBraceToken(t *testing.T) {
	rdr := strings.NewReader(`http {
  }`)
	scnr := NewScanner(rdr)
	for i := 0; i < 4; i++ {
		scnr.Scan()
	}
	tok, _ := scnr.Scan()

	if tok != CLOSE_BRACE {
		t.Error("Token mismatch")
	}
}

func TestReadLocationToken(t *testing.T) {
	rdr := strings.NewReader(`http {
    location / {
    }
  }`)
	scnr := NewScanner(rdr)
	for i := 0; i < 4; i++ {
		scnr.Scan()
	}
	tok, _ := scnr.Scan()

	if tok != LOCATION {
		t.Error("Token mismatch")
	}
}
