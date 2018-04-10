package nginxconf

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestRouteParsing(t *testing.T) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	pathToFile := fmt.Sprintf("%s/routes.json", workDir)

	conf := New(pathToFile)

	if len(conf.Routes) != 2 {
		errMsg := fmt.Sprintf("Expected Routes to be size %d, got %d", 2, len(conf.Routes))
		t.Error(errMsg)
	}
}
