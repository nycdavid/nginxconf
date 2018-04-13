package parserlexer

import (
	"bytes"
	"log"
	"strings"
)

type Scanner struct {
	rdr *strings.Reader
}

func NewScanner(confRdr *strings.Reader) *Scanner {
	return &Scanner{rdr: confRdr}
}

func (scnr *Scanner) Scan() (Token, string) {
	ch := scnr.read()
	// // detect if the next rune in the buffer is a whitespace or non-whitespace char
	// if isWhitespace(ch) {
	// 	// consume all whitespace and tokenize
	// 	scnr.unread()
	// 	return scnr.scanWhitespace()
	// }
	if isAlpha(ch) {
		scnr.unread()
		return scnr.scanIdent()
	}
	switch ch {
	case rune(0):
		return EOF, ""
	}
	return ILLEGAL, string(ch)
}

// Private
func (scnr *Scanner) read() rune {
	ch, _, err := scnr.rdr.ReadRune()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	return ch
}

func (scnr *Scanner) unread() {
	scnr.rdr.UnreadRune()
}

func (scnr *Scanner) scanIdent() (Token, string) {
	var buf bytes.Buffer
	for {
		ch := scnr.read()
		if ch != rune(0) || !isWhitespace(ch) {
			buf.WriteRune(ch)
		} else {
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

// Utility Fns
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
