package runtimes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mirror-backend/pkg"
	"net/http"
)

type TypescriptRuntime struct {
	url string
}

func NewTypescript(runtimeUrl string) pkg.CodeExecutor {
	return &TypescriptRuntime{url: runtimeUrl}
}

func (t *TypescriptRuntime) ExecuteCode(code string) (string, error) {
	reqBody := request{Code: code}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("error marshalling request", err)
		return "", pkg.ErrHttpRequest
	}

	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/code-exec/typescript", t.url), bytes.NewReader(reqBytes))
	if err != nil {
		log.Println("error creating request", err)
		return "", pkg.ErrHttpRequest
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("error sending request", err)
		return "", pkg.ErrHttpRequest
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return "", pkg.ErrHttpRequest
	}

	var res response
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println("error unmarshalling response", err, string(body))
		return "", pkg.ErrHttpRequest
	}

	if res.Error != "" {
		return res.Output, errors.New(res.Error)
	}
	return res.Output, nil
}
