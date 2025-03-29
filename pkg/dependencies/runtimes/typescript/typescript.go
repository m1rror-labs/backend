package typescript

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
	filename := "./" + id + ".ts"
	fullFilename := "./pkg/dependencies/runtimes/typescript/" + id + ".ts"
	err := os.WriteFile(fullFilename, []byte(code), 0644)
	if err != nil {
		log.Println("Error writing TypeScript file:", err)
		return "", err
	}
	defer os.Remove(fullFilename)
	defer os.Remove(filename)
	defer os.Remove("./pkg/dependencies/runtimes/typescript/dist/" + id + ".js")
	defer os.Remove("./pkg/dependencies/runtimes/typescript/dist/" + id + ".mjs")

	// Compile the TypeScript file to JavaScript
	cmd := exec.Command("npx", "tsc", "-t", "es2022", "-m", "es2022", "--skipLibCheck", "--isolatedModules", "--moduleResolution", "node", "--outDir", "dist", "--allowSyntheticDefaultImports", filename)
	cmd.Dir = "./pkg/dependencies/runtimes/typescript"
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Error compiling TypeScript:", string(output)) // Print the error logs
		return "", fmt.Errorf("error compiling TypeScript: %s", string(output))
	}
	log.Println("time taken to compile TypeScript:", time.Since(now)) // Log the time taken to compile
	now = time.Now()                                                  // Reset the timer after compilation

	jsFilename := "./pkg/dependencies/runtimes/typescript/dist/" + id + ".js"
	mjsFilename := "./pkg/dependencies/runtimes/typescript/dist/" + id + ".mjs"
	err = os.Rename(jsFilename, mjsFilename)
	if err != nil {
		log.Println("Error renaming JavaScript file:", err)
		return "", fmt.Errorf("error renaming file: %s", err)
	}
	defer os.Remove(mjsFilename)

	// Run the resulting JavaScript file
	shortMjsFilename := "./dist/" + id + ".mjs"
	cmd = exec.Command("node", "--no-warnings", shortMjsFilename)
	cmd.Dir = "./pkg/dependencies/runtimes/typescript"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error running JavaScript:", err)
		return "", fmt.Errorf("error running JavaScript: %s", string(output))
	}
	fmt.Println("time taken to run JavaScript:", time.Since(now)) // Log the time taken to run the JavaScript
	return string(output), nil
}
