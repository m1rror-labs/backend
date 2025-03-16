package codeexec

import (
	"context"
	"errors"
	"fmt"
	"mirror-backend/pkg"
	"net/url"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type ExecuteCodeRequest struct {
	Code string `json:"code" binding:"required"`
}

type LogWithUrl struct {
	ID                   uuid.UUID       `json:"id"`
	CreatedAt            time.Time       `json:"created_at"`
	Url                  string          `json:"url"`
	TransactionSignature string          `json:"transaction_signature"`
	Log                  string          `json:"log"`
	Index                int             `json:"index"`
	Transaction          pkg.Transaction `json:"transaction"`
}

func RunCode(ctx context.Context, code string, codeExecutor pkg.CodeExecutor, transactionRepo pkg.TransactionRepo) (string, []LogWithUrl, error) {
	engineID, err := parseEngineUrl(code)
	if err != nil {
		return "", nil, errors.New("code must contain a valid engine URL")
	}

	start := time.Now()
	output, err := codeExecutor.ExecuteCode(code)
	if err != nil {
		return "", nil, err
	}
	end := time.Now()

	logs, err := transactionRepo.ReadTransactionLogMessages().BlockchainID(engineID).Between(start, end).Execute(ctx)
	if err != nil {
		return "", nil, err
	}

	var logsWithUrl []LogWithUrl = []LogWithUrl{}
	for _, log := range logs {
		blockchainUrl := fmt.Sprintf("https://engine.mirror.ad/rpc/%s", engineID)

		logsWithUrl = append(logsWithUrl, LogWithUrl{
			ID:                   log.ID,
			CreatedAt:            log.CreatedAt,
			Url:                  fmt.Sprintf("https://explorer.solana.com/tx/%s?cluster=custom&customUrl=%s", log.TransactionSignature, url.QueryEscape(blockchainUrl)),
			TransactionSignature: log.TransactionSignature,
			Log:                  log.Log,
			Index:                log.Index,
			Transaction:          log.Transaction,
		})
	}

	return output, logsWithUrl, nil
}

func parseEngineUrl(code string) (uuid.UUID, error) {
	re := regexp.MustCompile(`https://engine\.mirror\.ad/rpc/([0-9a-fA-F-]{36})`)

	match := re.FindStringSubmatch(code)
	if len(match) < 2 {
		return uuid.Nil, fmt.Errorf("UUID not found in the code string")
	}

	uuidStr := match[1]
	engineID, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse UUID: %v", err)
	}
	return engineID, nil
}
