package pkg

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID           uuid.UUID `bun:"id,pk,type:uuid"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	BlockchainID uuid.UUID `bun:"blockchain,type:uuid,notnull"`
	Address      string    `json:"address"`
	Lamports     uint      `json:"lamports"`
	Data         []byte    `json:"data"`
	Owner        string    `json:"owner"`
	Executable   bool      `json:"executable"`
	RentEpoch    uint      `json:"rentEpoch"`
}

func (a Account) ToAccount() SolanaAccount {
	return SolanaAccount{
		Address:    a.Address,
		Lamports:   a.Lamports,
		Data:       a.Data,
		Owner:      a.Owner,
		Executable: a.Executable,
		RentEpoch:  a.RentEpoch,
	}
}

type SolanaAccount struct {
	Address    string `json:"address"`
	Lamports   uint   `json:"lamports"`
	Data       []byte `json:"data"`
	Owner      string `json:"owner"`
	Executable bool   `json:"executable"`
	RentEpoch  uint   `json:"rentEpoch"`
}

func (a SolanaAccount) ToAccount(blockchainID uuid.UUID) Account {
	return Account{
		ID:           uuid.New(),
		CreatedAt:    time.Now(),
		BlockchainID: blockchainID,
		Address:      a.Address,
		Lamports:     a.Lamports,
		Data:         a.Data,
		Owner:        a.Owner,
		Executable:   a.Executable,
		RentEpoch:    a.RentEpoch,
	}
}

type AccountRetriever interface {
	GetAccount(ctx context.Context, address string) (SolanaAccount, error)
	GetMultipleAccounts(ctx context.Context, addresses []string) ([]SolanaAccount, error)
	GetProgramOwnedAccounts(ctx context.Context, programID string) ([]SolanaAccount, error)
}

type AccountRepo interface {
	ReadAccounts() AccountReader
}

type AccountReader interface {
	Execute(ctx context.Context) ([]Account, error)
	ExecuteWithCount(ctx context.Context) ([]Account, int, error)
	ExecuteOne(ctx context.Context) (Account, error)

	BlockchainID(blockchainID uuid.UUID) AccountReader
	Paginate(page int, limit int) AccountReader
	Between(start time.Time, end time.Time) AccountReader

	OrderCreatedAt(order string) AccountReader
}
