package parserlexer

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseLexingHttpIdent(t *testing.T) {
	confRdr := strings.NewReader(`http {
  }`)
	directives := New(confRdr)
	directive := directives[0]

	if directive.String != "http" || directive.Type != HTTP {
		errMsg := fmt.Sprintf("Expected type %s, got %s", "http", directive.String)
		t.Error(errMsg)
	}
}
