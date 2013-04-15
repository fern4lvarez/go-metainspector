package metainspector

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v := Version()
	if v != "0.1.1" {
		t.Errorf(msgFail, "Version", "0.1.1", v)
	}
}
