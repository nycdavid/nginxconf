package nginxconf

import (
	"io"
)

type Decoder struct {
	r io.Reader
}

type Location struct {
	Path string
}

func NewDecoder(conf io.Reader) *Decoder {
	return &Decoder{r: conf}
}

func (dec *Decoder) Decode(v interface{}) error {
	var drctvs []*Directive
	drctv := &Directive{Type: "location"}
	drctvs = append(drctvs, drctv)
	nginxConf := v.(*NginxConf)
	nginxConf.Directives = drctvs
	return nil
}
