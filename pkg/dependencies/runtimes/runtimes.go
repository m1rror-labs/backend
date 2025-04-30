package runtimes

import "github.com/google/uuid"

type request struct {
	Code string `json:"code" binding:"required"`
}

type response struct {
	Error  string `json:"error"`
	Output string `json:"output"`
}

type buildRequest struct {
	Code         string    `json:"code" binding:"required"`
	ProgramID    string    `json:"program_id" binding:"required"`
	BlockchainID uuid.UUID `json:"blockchain_id" binding:"required"`
}

type buildResponse struct {
	Error string `json:"error"`
}

type testRequest struct {
	Code         string    `json:"code" binding:"required"`
	ProgramID    string    `json:"program_id" binding:"required"`
	BlockchainID uuid.UUID `json:"blockchain_id" binding:"required"`
	TestCode     string    `json:"test_code" binding:"required"`
}

type testResponse struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
