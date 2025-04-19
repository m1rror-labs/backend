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

	"github.com/google/uuid"
)

type AnchorRuntime struct {
	url string
}

func NewAnchor(runtimeUrl string) pkg.ProgramBuilder {
	return &AnchorRuntime{url: runtimeUrl}
}

func (a *AnchorRuntime) BuildAndDeployProgram(code string, programID string, blockchainID uuid.UUID) error {
	reqBody := buildRequest{Code: code, ProgramID: programID, BlockchainID: blockchainID}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("error marshalling request", err)
		return pkg.ErrHttpRequest
	}

	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/code-exec/programs/anchor", a.url), bytes.NewReader(reqBytes))
	if err != nil {
		log.Println("error creating request", err)
		return pkg.ErrHttpRequest
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("error sending request", err)
		return pkg.ErrHttpRequest
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return pkg.ErrHttpRequest
	}

	var res buildResponse
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println("error unmarshalling response", err, string(body))
		return pkg.ErrHttpRequest
	}

	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}
