package main

import (
	"mirror-backend/pkg/app"
	"mirror-backend/pkg/dependencies/jwt"
	"os"
)

func main() {
	supabaseJwtKey := os.Getenv("SUPABASE_JWT_SECRET")
	auth := jwt.NewAuthMiddleware(supabaseJwtKey)

	env := os.Getenv("ENV")

	app := app.NewApp(env, auth)

	app.Run()
}
