package parserlexer

import (
	"log"
	"strings"
)

type Scanner struct {
	rdr *strings.Reader
}

func NewScanner(confRdr *strings.Reader) *Scanner {
	return &Scanner{rdr: confRdr}
}

func (scnr *Scanner) read() rune {
	ch, _, err := scnr.rdr.ReadRune()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	return ch
}
