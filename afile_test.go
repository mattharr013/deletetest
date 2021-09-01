package deletetest

import (
	"testing"
)

func TestSomePublicFunc(t *testing.T) {
	value := SomePublicFunc()
	if value != "stuffy" {
		t.Errorf("should have been stuff was %s", value)
	}
}
