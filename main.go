package main

import (
	"log"
	"os"

	"github.com/velvetreactor/nginxconf/nginxconf"
)

func main() {
	// reader and writer
	arg := os.Args[1]
	file, err := os.Open(arg)
	if err != nil {
		log.Print(err)
		panic(err)
	}
	oFile, err := os.Create("test.conf")
	if err != nil {
		log.Print(err)
		panic(err)
	}

	conf := nginxconf.New(file)
	conf.WriteTo(oFile)
}
