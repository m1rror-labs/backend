package typescript

import (
	"fmt"
	"log"
	"mirror-backend/pkg"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type runtime struct {
}

func Runtime() pkg.CodeExecutor {
	return &runtime{}
}

func (r *runtime) ExecuteCode(code string) (string, error) {
	id := uuid.NewString()
	filename := "./" + id + ".ts"
	fullFilename := "./pkg/dependencies/runtimes/typescript/"+id+".ts"
	err := os.WriteFile(fullFilename, []byte(code), 0644)
	if err != nil {
		return "", err
	}
	defer os.Remove(fullFilename)
	defer os.Remove(filename)
	defer os.Remove("./dist/" + id + ".mjs")

	
	// Compile the TypeScript file to JavaScript
	cmd := exec.Command("npx", "tsc", "-t", "es2022", "-m", "es2022", "--moduleResolution", "node", "--outDir", "dist", filename)
	cmd.Dir = "./pkg/dependencies/runtimes/typescript"
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Error compiling TypeScript:", string(output)) // Print the error logs
		return "", fmt.Errorf("error compiling TypeScript: %s", string(output))
	}

	jsFilename := "./pkg/dependencies/runtimes/typescript/dist/" + id + ".js"
	mjsFilename := "./pkg/dependencies/runtimes/typescript/dist/" + id + ".mjs"
	err = os.Rename(jsFilename, mjsFilename)
	if err != nil {
		return "", fmt.Errorf("error renaming file: %s", err)
	}
	defer os.Remove(mjsFilename)

	// Run the resulting JavaScript file
	shortMjsFilename := "./dist/"+id+".mjs"
	cmd = exec.Command("node", "--no-warnings", shortMjsFilename)
	cmd.Dir = "./pkg/dependencies/runtimes/typescript"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error running JavaScript: %s", string(output))
	}
	fmt.Println(string(output)) // Print the console logs
	return string(output), nil
}
