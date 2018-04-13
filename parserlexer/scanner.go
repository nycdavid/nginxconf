package parserlexer

import (
	"bytes"
	"strings"
)

const (
	HTTP Token = iota // http directive
	IDENT
	ILLEGAL
	EOF
	WS
	OPEN_BRACE
)

type Scanner struct {
	rdr *strings.Reader
}

func NewScanner(confRdr *strings.Reader) *Scanner {
	return &Scanner{rdr: confRdr}
}

func (scnr *Scanner) Scan() (Token, string) {
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
		return EOF, ""
	case '{':
		return OPEN_BRACE, "{"
	}
	return ILLEGAL, string(ch)
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

func (scnr *Scanner) scanIdent() (Token, string) {
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
		return HTTP, buf.String()
	}
	return IDENT, buf.String()
}

func (scnr *Scanner) scanWhitespace() (Token, string) {
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
	return WS, buf.String()
}

// Utility Fns
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
