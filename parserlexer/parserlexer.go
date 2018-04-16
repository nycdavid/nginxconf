package parserlexer

import (
	"strings"
)

type Token int

type Directive struct {
	Tkn  Token
	Type string
}

func New(confRdr *strings.Reader) []*Directive {
	scnr := NewScanner(confRdr)
	var directives []*Directive
	for {
		if tok, tokStr := scnr.Scan(); tok != EOF {
			directives = append(directives, &Directive{Tkn: tok, Type: tokStr})
		} else {
			break
		}
	}
	return directives
}
