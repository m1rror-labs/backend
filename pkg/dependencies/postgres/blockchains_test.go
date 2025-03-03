package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestUpdateBlockchain(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	label := "new label"
	if err := rep.UpdateBlockchain(uuid.MustParse("41335c33-c715-4a07-9c55-4818ef900a97")).Label(&label).Execute(context.Background()); err != nil {
		t.Fatal(err)
	}
}
