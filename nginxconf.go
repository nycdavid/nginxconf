package nginxconf

type NginxConf struct {
	Directives []*Directive
}

type Directive struct {
	Type string
}
