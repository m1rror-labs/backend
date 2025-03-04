package rpcengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mirror-backend/pkg"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

type rpcEngine struct {
	url string
}

func New(url string) pkg.RpcEngine {
	return &rpcEngine{url}
}

func (e *rpcEngine) CreateBlockchain(ctx context.Context, apiKey uuid.UUID) (uuid.UUID, error) {
	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/blockchains", e.url), nil)
	if err != nil {
		log.Println("error creating request", err)
		return uuid.Nil, pkg.ErrHttpRequest
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("api_key", apiKey.String())

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("error sending request", err)
		return uuid.Nil, pkg.ErrHttpRequest
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return uuid.Nil, pkg.ErrHttpRequest
	}

	type Error struct {
		Message string `json:"message"`
	}

	if resp.StatusCode != http.StatusOK {
		var res Error
		if err := json.Unmarshal(body, &res); err != nil {
			log.Println("error unmarshalling error", err, string(body))
			return uuid.Nil, pkg.ErrHttpRequest
		}

		return uuid.Nil, errors.New(res.Message)
	}

	type Response struct {
		Url string `json:"url"`
	}

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println("error unmarshalling response", err, string(body))
		return uuid.Nil, pkg.ErrHttpRequest
	}

	id := removeMirrorRPC(res.Url)
	log.Println(id)
	return uuid.Parse(id)
}

func removeMirrorRPC(url string) string {
	re := regexp.MustCompile(`https://rpc\.mirror\.ad/rpc/`)
	return re.ReplaceAllString(url, "")
}

func (e *rpcEngine) DeleteBlockchain(ctx context.Context, apiKey uuid.UUID, blockchainID uuid.UUID) error {
	r, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/rpc/%s", e.url, blockchainID.String()), nil)
	if err != nil {
		log.Println("error creating request", err)
		return pkg.ErrHttpRequest
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("api_key", apiKey.String())

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

	type Error struct {
		Message string `json:"message"`
	}

	if resp.StatusCode != http.StatusOK {
		var res Error
		if err := json.Unmarshal(body, &res); err != nil {
			log.Println("error unmarshalling error", err, string(body))
			return pkg.ErrHttpRequest
		}

		return errors.New(res.Message)
	}

	return nil
}
