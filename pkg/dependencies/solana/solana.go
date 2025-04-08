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
	data := account.Bytes()
	if data == nil {
		data = []byte{}
	}

	return pkg.SolanaAccount{
		Address:    address,
		Lamports:   uint(account.Value.Lamports),
		Data:       data,
		Owner:      account.Value.Owner.String(),
		Executable: account.Value.Executable,
		RentEpoch:  uint(account.Value.RentEpoch.Uint64()),
	}, nil

}

func (a *accountRetriever) GetMultipleAccounts(ctx context.Context, addresses []string) ([]pkg.SolanaAccount, error) {
	rpcClient := rpc.New(a.rpcUrl)

	var pubkeys []solana.PublicKey
	for _, address := range addresses {
		pubkey, err := solana.PublicKeyFromBase58(address)
		if err != nil {
			return nil, pkg.ErrInvalidPubkey
		}
		pubkeys = append(pubkeys, pubkey)
	}

	accounts, err := rpcClient.GetMultipleAccounts(ctx, pubkeys...)
	if err != nil {
		return nil, err
	}
	if accounts == nil {
		return nil, pkg.ErrAccountNotFound
	}
	var solanaAccounts []pkg.SolanaAccount
	for i, account := range accounts.Value {
		if account == nil {
			continue
		}
		data := account.Data.GetBinary()
		if data == nil {
			data = []byte{}
		}
		solanaAccounts = append(solanaAccounts, pkg.SolanaAccount{
			Address:    addresses[i],
			Lamports:   uint(account.Lamports),
			Data:       data,
			Owner:      account.Owner.String(),
			Executable: account.Executable,
			RentEpoch:  uint(account.RentEpoch.Uint64()),
		})
	}
	return solanaAccounts, nil
}

func (a *accountRetriever) GetProgramOwnedAccounts(ctx context.Context, programID string) ([]pkg.SolanaAccount, error) {
	rpcClient := rpc.New(a.rpcUrl)

	pubkey, err := solana.PublicKeyFromBase58(programID)
	if err != nil {
		return nil, pkg.ErrInvalidPubkey
	}

	accounts, err := rpcClient.GetProgramAccounts(ctx, pubkey)
	if err != nil {
		return nil, err
	}

	var solanaAccounts []pkg.SolanaAccount
	for _, account := range accounts {
		if account == nil {
			continue
		}
		solanaAccounts = append(solanaAccounts, pkg.SolanaAccount{
			Address:    account.Pubkey.String(),
			Lamports:   uint(account.Account.Lamports),
			Data:       account.Account.Data.GetBinary(),
			Owner:      account.Account.Owner.String(),
			Executable: account.Account.Executable,
			RentEpoch:  uint(account.Account.RentEpoch.Uint64()),
		})
	}
	return solanaAccounts, nil
}

func (a *accountRetriever) GetSignaturesForAddress(ctx context.Context, address string, limit int) ([]string, error) {
	rpcClient := rpc.New(a.rpcUrl)

	pubkey, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return nil, pkg.ErrInvalidPubkey
	}

	signatures, err := rpcClient.GetSignaturesForAddressWithOpts(ctx, pubkey, &rpc.GetSignaturesForAddressOpts{
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	var signatureList []string
	for _, sig := range signatures {
		signatureList = append(signatureList, sig.Signature.String())
	}

	return signatureList, nil
}

func (a *accountRetriever) GetTransactionAccountKeys(ctx context.Context, signature string) ([]string, error) {
	rpcClient := rpc.New(a.rpcUrl)

	sig, err := solana.SignatureFromBase58(signature)
	if err != nil {
		return nil, pkg.ErrInvalidSignature
	}

	maxVersion := uint64(1)
	res, err := rpcClient.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
		MaxSupportedTransactionVersion: &maxVersion,
	})
	if err != nil {
		return nil, err
	}

	tx, err := res.Transaction.GetTransaction()
	if err != nil {
		return nil, err
	}

	if tx == nil {
		return nil, pkg.ErrTransactionNotFound
	}

	var accountKeys []string
	for _, key := range tx.Message.AccountKeys {
		accountKeys = append(accountKeys, key.String())
	}

	return accountKeys, nil
}
