package nginxconf

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
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
	buf.WriteString("foo")
}
