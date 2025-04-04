package main

import (
	"mirror-backend/pkg/app"
	"mirror-backend/pkg/dependencies/jwt"
	"mirror-backend/pkg/dependencies/postgres"
	"mirror-backend/pkg/dependencies/rpcengine"
	"mirror-backend/pkg/dependencies/solana"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	auth := jwt.NewAuthMiddleware(os.Getenv("SUPABASE_JWT_SECRET"))
	repo := postgres.NewRepository(os.Getenv("DATABASE_URL"))
	rpcEngine := rpcengine.New(os.Getenv("ENGINE_URL"))
	accountRetriever := solana.NewAccountRetriever(os.Getenv("SOLANA_RPC_URL"))

	env := os.Getenv("ENV")

	app := app.NewApp(env, auth, repo, rpcEngine, accountRetriever)

	app.Run()
}
