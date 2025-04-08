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

func TestGetSignaturesForAddress(t *testing.T) {
	t.Skip("Skipping test for Solana signatures retrieval")
	godotenv.Load("../../../.env")
	rpcUrl := os.Getenv("SOLANA_RPC_URL")
	accountRetriever := NewAccountRetriever(rpcUrl)

	// Test with a valid address
	address := "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
	signatures, err := accountRetriever.GetSignaturesForAddress(context.Background(), address, 25)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Fatal("Signatures:", signatures)
}

func TestGetTransactionAccountKeys(t *testing.T) {
	t.Skip("Skipping test for Solana transaction account keys retrieval")
	godotenv.Load("../../../.env")
	rpcUrl := os.Getenv("SOLANA_RPC_URL")
	accountRetriever := NewAccountRetriever(rpcUrl)

	// Test with a valid signature
	signature := "5XA7QzwEgY7BVFnUyobBa9aMXGy69neytVSYxEjSQfwe1jQhGv1fhAPYArGNkEQfrGdxy7NtJ2ULsuK4dKemFb8P"
	accountKeys, err := accountRetriever.GetTransactionAccountKeys(context.Background(), signature)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Fatal("Account Keys:", accountKeys)
}
