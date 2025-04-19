package pkg

type Dependencies struct {
	Auth             Auth
	Repo             Repository
	RpcEngine        RpcEngine
	TsRuntime        CodeExecutor
	RustRuntime      CodeExecutor
	AnchorRuntime    ProgramBuilder
	AccountRetriever AccountRetriever
}

type Repository interface {
	AccountRepo
	BlockchainRepo
	TransactionRepo
	UserRepo
}
