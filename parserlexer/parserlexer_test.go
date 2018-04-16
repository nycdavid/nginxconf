package parserlexer

import (
	"fmt"
	"testing"
)

func TestParseLexingHttpIdent(t *testing.T) {
	snippet := `
  http {
  }
  `
	plxr := New(snippet)
	directive := plxr.Directives[0]

	if directive.Type != "http" {
		errMsg := fmt.Sprintf("Expected type %s, got %s", "http", directive.Type)
		t.Error(errMsg)
	}
}
