package {{plugin}}

import "testing"

func TestName(t *testing.T) {
	p := New()
	if p.Name() == "" {
		t.Fatal("expected Name to be non-empty")
	}
}
