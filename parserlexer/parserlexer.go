package parserlexer

import ()

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
