package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestReadAccounts(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	accounts, err := rep.ReadAccounts().BlockchainID(uuid.MustParse("924e3393-10c9-4044-8e0b-1b17ae032c3a")).Paginate(1, 10).Execute(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) == 0 {
		t.Fatal("No accounts found")
	}

	for _, account := range accounts {
		if account.BlockchainID == uuid.Nil {
			t.Fatal("Blockchain ID is nil")
		}
	}
	t.Fatal("Accounts found: ", accounts)
}
