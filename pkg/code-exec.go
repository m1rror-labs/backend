package pkg

import "github.com/google/uuid"

type CodeExecutor interface {
	ExecuteCode(code string) (string, error)
}

type ProgramBuilder interface {
	BuildAndDeployProgram(code string, programID string, blockchainID uuid.UUID) error
}
