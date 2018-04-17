package nginxconf

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/velvetreactor/nginxconf/parserlexer"
)

type NginxConf struct {
	Routes []map[string]string `json:"routes"`
}

func NewNginxConf(routes *strings.Reader) *NginxConf {
	var conf NginxConf
	dec := json.NewDecoder(routes)
	err := dec.Decode(&conf)
	if err != nil {
		log.Print(err)
	}
	return &conf
}

func (conf *NginxConf) WriteTo(buf *bytes.Buffer) {
	httpTok := &parserlexer.Token{Type: parserlexer.HTTP, String: "http"}
	buf.WriteString(httpTok.String)
	buf.WriteString(" {")

	for _, route := range conf.Routes {
		buf.WriteString(" location ")
		buf.WriteString(route["location"])
		buf.WriteString(" { ")
		buf.WriteString(" proxy_pass ")
		buf.WriteString(route["proxy_pass"])
		buf.WriteString(" } ")
	}

	buf.WriteString(" }")
}
