package rust

import "testing"

func TestExecuteCode(t *testing.T) {
	t.Skip()
	r := Runtime()
	output, err := r.ExecuteCode("fn main() { println!(\"Hello, world!\"); }")
	if err != nil {
		t.Errorf("Error executing Rust code: %s", err)
	}
	t.Fatal(output)
}
