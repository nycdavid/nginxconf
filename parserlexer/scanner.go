package parserlexer

import (
	"bytes"
	"strings"
)

type Construct int

const (
	HTTP Construct = iota // http directive
	IDENT
	ILLEGAL
	EOF
	WS
	OPEN_BRACE
	CLOSE_BRACE
	LOCATION
)

type Token struct {
	Type   Construct
	String string
}

type Scanner struct {
	rdr *strings.Reader
}

func NewScanner(confRdr *strings.Reader) *Scanner {
	return &Scanner{rdr: confRdr}
}

func (scnr *Scanner) Scan() *Token {
	ch := scnr.read()
	if isAlpha(ch) {
		scnr.unread()
		return scnr.scanIdent()
	}
	if isWhitespace(ch) {
		scnr.unread()
		return scnr.scanWhitespace()
	}
	switch ch {
	case rune(0):
		return &Token{Type: EOF, String: ""}
	case '{':
		return &Token{Type: OPEN_BRACE, String: "{"}
	case '}':
		return &Token{Type: CLOSE_BRACE, String: "}"}
	}
	return &Token{Type: ILLEGAL, String: string(ch)}
}

// Private
func (scnr *Scanner) read() rune {
	ch, _, err := scnr.rdr.ReadRune()
	if err != nil {
		return rune(0)
	}
	return ch
}

func (scnr *Scanner) unread() {
	scnr.rdr.UnreadRune()
}

func (scnr *Scanner) scanIdent() *Token {
	var buf bytes.Buffer
	for {
		if ch := scnr.read(); ch == rune(0) {
			break
		} else if !isWhitespace(ch) {
			buf.WriteRune(ch)
		} else if isWhitespace(ch) {
			scnr.unread()
			break
		}
	}
	switch strings.ToLower(buf.String()) {
	case "http":
		return &Token{Type: HTTP, String: buf.String()}
	case "location":
		return &Token{Type: LOCATION, String: buf.String()}
	}
	return &Token{Type: IDENT, String: buf.String()}
}

func (scnr *Scanner) scanWhitespace() *Token {
	var buf bytes.Buffer
	for {
		if ch := scnr.read(); ch == rune(0) {
			break
		} else if !isWhitespace(ch) {
			scnr.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return &Token{Type: WS, String: buf.String()}
}

// Utility Fns
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
