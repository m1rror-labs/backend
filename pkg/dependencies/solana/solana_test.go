package solana

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetAccount(t *testing.T) {
	t.Skip("Skipping test for Solana account retrieval")
	godotenv.Load("../../../.env")
	rpcUrl := os.Getenv("SOLANA_RPC_URL")
	accountRetriever := NewAccountRetriever(rpcUrl)

	// Test with a valid address
	address := "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
	account, err := accountRetriever.GetAccount(context.Background(), address)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Fatal("Account:", account)
}

func TestGetProgramOwnedAccounts(t *testing.T) {
	t.Skip("Skipping test for Solana program owned accounts retrieval")
	godotenv.Load("../../../.env")
	rpcUrl := os.Getenv("SOLANA_RPC_URL")
	accountRetriever := NewAccountRetriever(rpcUrl)

	// Test with a valid program address
	programAddress := "L2TExMFKdjpN9kozasaurPirfHy9P8sbXoAN1qA3S95"
	accounts, err := accountRetriever.GetProgramOwnedAccounts(context.Background(), programAddress)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Fatal("Accounts:", accounts)
}
