package architecture

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Doesnt work in architecture amd64")
	}
	t.Fail()
}