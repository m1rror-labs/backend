package codeexec

import (
	"context"
	"mirror-backend/pkg"
)

type ExecuteCodeRequest struct {
	Code string `json:"code" binding:"required"`
}

func RunCode(ctx context.Context, code string, codeExecutor pkg.CodeExecutor) (string, error) {
	return codeExecutor.ExecuteCode(code)
}
