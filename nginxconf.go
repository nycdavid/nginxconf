package nginxconf

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"text/template"
)

const confTmpl = `http {
	server {
		listen 80;
		{{range .Routes}}
		location {{.HostEndpoint}} {
			rewrite ^/google/(.*)$ /$1 break;
			proxy_pass {{.ProxyTo}};
		}
		{{end}}
	}
}`

type NginxConf struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	HostEndpoint string `json:"host_endpoint"`
	ProxyTo      string `json:"proxy_to"`
	Rewrite      bool   `json:"rewrite"`
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
	tmpl := template.New("nginxConf")
	tmpl.Parse(confTmpl)
	err := tmpl.Execute(buf, conf)
	if err != nil {
		log.Print(err)
	}
}
