package nginxconf

import (
	"encoding/json"
	"io"
	"log"
	"text/template"
)

const confTmpl = `worker_processes 5;
worker_rlimit_nofile 8192;

events {
	worker_connections 4096;
}

http {
	server {
		listen 80;
		{{range .Routes -}}
		location {{.HostEndpoint}} {
			{{- if .AppendPath}}
			rewrite ^/(.*)$ /$1 break;
			{{else}}
			rewrite ^/(.*)$ / break;
			{{- end}}
			proxy_pass {{.ProxyTo}};
		}
		{{- end}}
	}
}`

type NginxConf struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	HostEndpoint string `json:"host_endpoint"`
	ProxyTo      string `json:"proxy_to"`
	AppendPath   bool   `json:"append_path"`
}

func New(routes io.Reader) *NginxConf {
	var conf NginxConf
	dec := json.NewDecoder(routes)
	err := dec.Decode(&conf)
	if err != nil {
		log.Print(err)
	}
	return &conf
}

func (conf *NginxConf) WriteTo(buf io.Writer) {
	tmpl := template.New("nginxConf")
	tmpl.Parse(confTmpl)
	err := tmpl.Execute(buf, conf)
	if err != nil {
		log.Print(err)
	}
}
