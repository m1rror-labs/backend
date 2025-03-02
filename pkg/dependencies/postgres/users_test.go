package postgres

import (
	"context"
	"mirror-backend/pkg"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestReadUser(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	users, err := rep.ReadUser().
		ID(uuid.MustParse("170dbec3-2983-415f-9f8b-090cb426688b")).
		Email("hunterwilliamsimmons@gmail.com").
		TeamID(uuid.MustParse("15b1eed5-6148-40ce-97dd-c0aaaa43bef0")).
		WithApiKeys().
		WithBlockchains().
		Execute(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(*users[0].Team)
}

func TestCreateUser(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	user := pkg.User{
		Email:  "hunterwilliamsimmons@gmail.com",
		TeamID: uuid.MustParse("15b1eed5-6148-40ce-97dd-c0aaaa43bef0"),
	}

	err := rep.CreateUser(context.Background(), &user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	user := pkg.User{
		ID:        uuid.MustParse("a04bca4f-0fef-434a-bb28-6c401d397ca0"),
		CreatedAt: time.Now(),
		Email:     "hunter@zapmoto.com",
		TeamID:    uuid.MustParse("15b1eed5-6148-40ce-97dd-c0aaaa43bef0"),
	}

	err := rep.UpdateUser(context.Background(), &user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	user := pkg.User{
		ID: uuid.MustParse("a04bca4f-0fef-434a-bb28-6c401d397ca0"),
	}

	err := rep.DeleteUser(context.Background(), &user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadTeam(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	teams, err := rep.ReadTeam().
		ID(uuid.MustParse("15b1eed5-6148-40ce-97dd-c0aaaa43bef0")).
		WithApiKeys().
		Execute(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(teams)
}
