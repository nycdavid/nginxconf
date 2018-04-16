package nginxconf

import (
	"fmt"
	"strings"
	"testing"
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
