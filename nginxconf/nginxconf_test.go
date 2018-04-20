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
	nginxConf := New(routesRdr)

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
	nginxConf := New(routesRdr)

	nginxConf.WriteTo(&buf)
	written := buf.String()

	rdr := strings.NewReader(written)
	scnr := parserlexer.NewScanner(rdr)
	var tokens []*parserlexer.Token
	for {
		if tok := scnr.Scan(); tok.Type != parserlexer.EOF {
			tokens = append(tokens, tok)
		} else {
			break
		}
	}

	directives := []string{"server", "location", "proxy_pass"}
	endpts := []string{"location /google", "location /elastic"}
	proxies := []string{
		"proxy_pass http://www.elastic.co;",
		"proxy_pass http://www.google.com;",
	}
	assertions := append(directives, endpts...)
	assertions = append(assertions, proxies...)

	for _, assertion := range assertions {
		if !strings.Contains(written, assertion) {
			errMsg := fmt.Sprintf("Expected string %s", assertion)
			t.Error(errMsg)
		}
	}
}

func TestConditionalRewrite(t *testing.T) {
	routes := strings.NewReader(`{
		"routes": [
			{
				"host_endpoint": "/google",
				"proxy_to": "http://www.google.com",
				"rewrite": false
			}
		]
	}`)
	conf := New(routes)
	var buf bytes.Buffer
	conf.WriteTo(&buf)
	written := buf.String()

	if strings.Contains(written, "rewrite") {
		t.Error("Unexpected \"rewrite\" directive")
	}
}
