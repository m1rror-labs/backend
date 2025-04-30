package codeexec

import (
	"context"
	"mirror-backend/pkg"

	"github.com/google/uuid"
)

type BuildProgramRequest struct {
	Code         string    `json:"code" binding:"required"`
	ProgramID    string    `json:"program_id" binding:"required"`
	BlockchainID uuid.UUID `json:"blockchain_id" binding:"required"`
}

type BuildAndTestRequest struct {
	Code         string    `json:"code" binding:"required"`
	ProgramID    string    `json:"program_id" binding:"required"`
	BlockchainID uuid.UUID `json:"blockchain_id" binding:"required"`
	TestCode     string    `json:"test_code" binding:"required"`
}

func BuildAndLoadProgram(
	ctx context.Context,
	code string,
	programID string,
	blockchainID uuid.UUID,
	programBuilder pkg.ProgramBuilder,
	rpcEngine pkg.RpcEngine,
) error {
	return programBuilder.BuildAndDeployProgram(code, programID, blockchainID)
}

func BuildAndTestProgram(
	ctx context.Context,
	code string,
	programID string,
	blockchainID uuid.UUID,
	testCode string,
	programBuilder pkg.ProgramBuilder,
	rpcEngine pkg.RpcEngine,
) (string, error) {
	return programBuilder.BuildAndTestProgram(code, programID, blockchainID, testCode)
}
