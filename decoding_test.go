package nginxconf

import (
	"fmt"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	nginxConfStr := strings.NewReader(`
  http {
    server {
      listen 80;
      location /google {
        proxy_pass http://www.google.com;
      }
    }
  }
  `)
	var nginxConf NginxConf

	dec := NewDecoder(nginxConfStr)
	err := dec.Decode(&nginxConf)

	firstDirective := nginxConf.Directives[0]
	if err != nil {
		t.Error(err)
	}
	if firstDirective.Type != "server" {
		errMsg := fmt.Sprintf("Expected type to be \"%s\", got \"%s\"", "server", firstDirective.Type)
		t.Error(errMsg)
	}
}
