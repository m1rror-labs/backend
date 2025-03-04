package pkg

type Repository interface {
	BlockchainRepo
	TransactionRepo
	UserRepo
}
