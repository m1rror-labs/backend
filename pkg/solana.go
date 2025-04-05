package pkg

import "context"

type SolanaAccount struct {
	Address    string `json:"address"`
	Lamports   uint   `json:"lamports"`
	Data       []byte `json:"data"`
	Owner      string `json:"owner"`
	Executable bool   `json:"executable"`
	RentEpoch  uint   `json:"rentEpoch"`
}

type AccountRetriever interface {
	GetAccount(ctx context.Context, address string) (SolanaAccount, error)
	GetMultipleAccounts(ctx context.Context, addresses []string) ([]SolanaAccount, error)
}
