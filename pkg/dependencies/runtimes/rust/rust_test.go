package rust

import (
	"mirror-backend/pkg/dependencies/multisync"
	"testing"
)

func TestExecuteCode(t *testing.T) {
	t.Skip()
	r := NewRuntime(multisync.NewMutex(1))
	output, err := r.ExecuteCode("fn main() { println!(\"Hello, world!\"); }")
	if err != nil {
		t.Errorf("Error executing Rust code: %s", err)
	}
	t.Fatal(output)
}
