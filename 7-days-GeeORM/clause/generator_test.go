package clause

import (
	"testing"
)

func TestGenBindVars(t *testing.T) {
	res := genBindVars(5)
	if res != "?,?,?,?,?" {
		t.Fatal("genBindVars failed res : ", res)
	}
}
