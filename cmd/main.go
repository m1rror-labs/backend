package main

import (
	"mirror-backend/pkg/app"
	"mirror-backend/pkg/dependencies/jwt"
	"mirror-backend/pkg/dependencies/postgres"
	"mirror-backend/pkg/dependencies/rpcengine"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	supabaseJwtKey := os.Getenv("SUPABASE_JWT_SECRET")
	auth := jwt.NewAuthMiddleware(supabaseJwtKey)

	dbUrl := os.Getenv("DATABASE_URL")
	repo := postgres.NewRepository(dbUrl)

	engineURL := os.Getenv("ENGINE_URL")
	rpcEngine := rpcengine.New(engineURL)

	env := os.Getenv("ENV")

	app := app.NewApp(env, auth, repo, rpcEngine)

	app.Run()
}
