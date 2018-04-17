package parserlexer

import (
	"strings"
)

func New(confRdr *strings.Reader) []*Token {
	scnr := NewScanner(confRdr)
	var directives []*Token
	for {
		if tok := scnr.Scan(); tok.Type != EOF {
			directives = append(directives, tok)
		} else {
			break
		}
	}
	return directives
}
