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
