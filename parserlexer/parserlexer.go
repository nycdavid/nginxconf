package parserlexer

import ()

const (
	HTTP Token = iota // http directive
	IDENT
	ILLEGAL
	EOF
)

type Token int

type Directive struct {
	Type string
}

type NginxAST struct {
	Directives []*Directive
}

func New(conf string) *NginxAST {
	return &NginxAST{}
}
