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
