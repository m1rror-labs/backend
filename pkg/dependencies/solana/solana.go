package solana

import (
	"context"
	"mirror-backend/pkg"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type accountRetriever struct {
	rpcUrl string
}

func NewAccountRetriever(rpcUrl string) pkg.AccountRetriever {
	return &accountRetriever{
		rpcUrl: rpcUrl,
	}
}

func (a *accountRetriever) GetAccount(ctx context.Context, address string) (pkg.SolanaAccount, error) {
	rpcClient := rpc.New(a.rpcUrl)

	pubkey, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return pkg.SolanaAccount{}, pkg.ErrInvalidPubkey
	}

	account, err := rpcClient.GetAccountInfo(ctx, pubkey)
	if err != nil {
		return pkg.SolanaAccount{}, err
	}

	if account == nil {
		return pkg.SolanaAccount{}, pkg.ErrAccountNotFound
	}

	return pkg.SolanaAccount{
		Address:    address,
		Lamports:   uint(account.Value.Lamports),
		Data:       account.Bytes(),
		Owner:      account.Value.Owner.String(),
		Executable: account.Value.Executable,
		RentEpoch:  uint(account.Value.RentEpoch.Uint64()),
	}, nil

}
