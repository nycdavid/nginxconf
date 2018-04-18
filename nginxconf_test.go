package nginxconf

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/velvetreactor/nginxconf/parserlexer"
)

func TestIngestRoutesFile(t *testing.T) {
	routesRdr := strings.NewReader(`{
    "routes": [
      { "location": "/google", "proxy_pass": "http://www.google.com" },
      { "location": "/elasticsearch", "proxy_pass": "http://www.elastic.co" }
    ]
  }`)
	nginxConf := NewNginxConf(routesRdr)

	routesCt := len(nginxConf.Routes)
	if routesCt != 2 {
		errMsg := fmt.Sprintf("Expected %d routes, got %d", 2, routesCt)
		t.Error(errMsg)
	}
}

func TestWriteTo(t *testing.T) {
	var buf bytes.Buffer
	routesRdr := strings.NewReader(`{
    "routes": [
			{
				"host_endpoint": "/google",
				"proxy_to": "http://www.google.com",
				"rewrite": true
			},
			{
				"host_endpoint": "/elastic",
				"proxy_to": "http://www.elastic.co",
				"rewrite": true
			}
    ]
  }`)
	nginxConf := NewNginxConf(routesRdr)

	nginxConf.WriteTo(&buf)

	rdr := strings.NewReader(buf.String())
	scnr := parserlexer.NewScanner(rdr)
	var tokens []*parserlexer.Token
	for {
		if tok := scnr.Scan(); tok.Type != parserlexer.EOF {
			tokens = append(tokens, tok)
		} else {
			break
		}
	}

	if len(tokens) != 43 {
		errMsg := fmt.Sprintf("Expected %d tokens, got %d", 31, len(tokens))
		t.Error(errMsg)
	}
}
