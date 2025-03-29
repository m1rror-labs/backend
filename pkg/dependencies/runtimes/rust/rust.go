package rust

import (
	"fmt"
	"log"
	"mirror-backend/pkg"
	"mirror-backend/pkg/dependencies/multisync"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

type runtime struct {
	mu *multisync.Mutex
}

func Runtime(mu *multisync.Mutex) pkg.CodeExecutor {
	return &runtime{mu}
}

func (r *runtime) ExecuteCode(code string) (string, error) {
	awaiting := r.mu.Acquire()
	defer r.mu.Release()
	<-awaiting

	now := time.Now()
	id := uuid.NewString()
	fullFilename := "./pkg/dependencies/runtimes/rust/src/bin/" + id + ".rs"
	err := os.WriteFile(fullFilename, []byte(code), 0644)
	if err != nil {
		return "", err
	}
	defer os.Remove(fullFilename)

	cmd := exec.Command("cargo", "run", "--locked", "--bin", id)
	cmd.Dir = "./pkg/dependencies/runtimes/rust"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error running Rust: %s", string(output))
	}
	fmt.Println("time taken to run Rust:", time.Since(now)) // Log the time taken to run the JavaScript
	return string(output), nil
}
