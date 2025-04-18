package main

import (
	"mirror-backend/pkg/app"
	"mirror-backend/pkg/dependencies/jwt"
	"mirror-backend/pkg/dependencies/postgres"
	"mirror-backend/pkg/dependencies/rpcengine"
	"mirror-backend/pkg/dependencies/runtimes"
	"mirror-backend/pkg/dependencies/solana"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	approvedCodeExecRaw := os.Getenv("APPROVED_CODE_EXEC")
	approvedCodeExec := strings.Split(approvedCodeExecRaw, ",")
	auth := jwt.NewAuthMiddleware(os.Getenv("SUPABASE_JWT_SECRET"), approvedCodeExec)
	repo := postgres.NewRepository(os.Getenv("DATABASE_URL"))
	rpcEngine := rpcengine.New(os.Getenv("ENGINE_URL"))
	accountRetriever := solana.NewAccountRetriever(os.Getenv("SOLANA_RPC_URL"))

	tsRuntime := runtimes.NewTypescript(os.Getenv("CODE_EXEC_URL"))
	rustRuntime := runtimes.NewRust(os.Getenv("CODE_EXEC_URL"))

	env := os.Getenv("ENV")

	app := app.NewApp(
		env,
		auth,
		repo,
		rpcEngine,
		accountRetriever,
		tsRuntime,
		rustRuntime,
	)

	app.Run()
}
