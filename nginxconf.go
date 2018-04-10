package nginxconf

import (
	"encoding/json"
	"log"
	"os"
)

type NginxConf struct {
	Routes []map[string]string `json:"routes"`
}

func New(fpath string) *NginxConf {
	return parseAndDecode(fpath)
}

func parseAndDecode(fpath string) *NginxConf {
	file, err := os.Open(fpath)
	if err != nil {
		log.Print(err)
		panic(err)
	}
	var conf NginxConf
	dec := json.NewDecoder(file)
	err = dec.Decode(&conf)
	if err != nil {
		log.Print(err)
	}
	return &conf
}
