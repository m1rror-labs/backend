package pkg

type Dependencies struct {
	Auth             Auth
	Repo             Repository
	RpcEngine        RpcEngine
	TsRuntime        CodeExecutor
	RustRuntime      CodeExecutor
	AccountRetriever AccountRetriever
}

type Repository interface {
	AccountRepo
	BlockchainRepo
	TransactionRepo
	UserRepo
}
