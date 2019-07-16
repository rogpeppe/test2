package test2

import (
	"testing"

	"github.com/rogpeppe/test"
)

func TestTest(t *testing.T) {
	if r := test.Test(); r != "test v1.0.0" {
		t.Fatalf("unexpected value %q", r)
	}
}
