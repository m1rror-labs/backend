package codeexec

import (
	"context"
	"errors"
	"fmt"
	"mirror-backend/pkg"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type ExecuteCodeRequest struct {
	Code string `json:"code" binding:"required"`
}

func RunCode(ctx context.Context, code string, codeExecutor pkg.CodeExecutor, transactionRepo pkg.TransactionRepo) (string, []pkg.TransactionLogMessage, error) {
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

	return output, logs, nil
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
