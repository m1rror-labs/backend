package pkg

type Repository interface {
	AccountRepo
	BlockchainRepo
	TransactionRepo
	UserRepo
}
