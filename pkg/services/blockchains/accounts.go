package blockchains

import (
	"context"
	"mirror-backend/pkg"
	"slices"

	"github.com/google/uuid"
)

func SetMainnetAccountState(
	ctx context.Context,
	rpcEngine pkg.RpcEngine,
	accountRetriever pkg.AccountRetriever,
	blockchainID uuid.UUID,
	accounts []string,
	label *string,
	token_mint_auth *string,
) error {
	// TODO: Billing stuff
	accountsData, err := accountRetriever.GetMultipleAccounts(ctx, accounts)
	if err != nil {
		return err
	}

	if err := rpcEngine.SetAccounts(ctx, blockchainID, accountsData, label, token_mint_auth); err != nil {
		return pkg.ErrSettingAccount
	}
	return nil
}

func SetAccountStateFromRecentTransactions(
	ctx context.Context,
	rpcEngine pkg.RpcEngine,
	accountRetriever pkg.AccountRetriever,
	blockchainID uuid.UUID,
	account string,
) error {
	signatures, err := accountRetriever.GetSignaturesForAddress(ctx, account, 25)
	if err != nil {
		return err
	}

	var accounts []string
	for _, signature := range signatures {
		aa, err := accountRetriever.GetTransactionAccountKeys(ctx, signature)
		if err != nil {
			return err
		}
		for _, account := range aa {
			if !slices.Contains(accounts, account) {
				accounts = append(accounts, account)
			}
		}
	}

	var accountsData []pkg.SolanaAccount
	var secondOrderAccounts []string
	// get accounts every 100 because solana has a limit of 100 accounts per request
	for i := 0; i < len(accounts); i += 99 {
		end := min(i+99, len(accounts))
		accountsDataBatch, err := accountRetriever.GetMultipleAccounts(ctx, accounts[i:end])
		if err != nil {
			return err
		}
		accountsData = append(accountsData, accountsDataBatch...)
		for _, account := range accountsDataBatch {
			if !slices.Contains(accounts, account.Owner) && !slices.Contains(secondOrderAccounts, account.Owner) {
				secondOrderAccounts = append(secondOrderAccounts, account.Owner)
			}
		}
	}
	for i := 0; i < len(secondOrderAccounts); i += 99 {
		end := min(i+99, len(secondOrderAccounts))
		accountsDataBatch, err := accountRetriever.GetMultipleAccounts(ctx, secondOrderAccounts[i:end])
		if err != nil {
			return err
		}
		accountsData = append(accountsData, accountsDataBatch...)
	}
	if len(accountsData) == 0 {
		return pkg.ErrNoAccounts
	}

	if err := rpcEngine.SetAccounts(ctx, blockchainID, accountsData, nil, nil); err != nil {
		return pkg.ErrSettingAccount
	}
	return nil
}

func GetAccounts(
	ctx context.Context,
	accountRepo pkg.AccountRepo,
	blockchainID uuid.UUID,
	page,
	limit int,
) ([]pkg.Account, int, error) {
	return accountRepo.ReadAccounts().BlockchainID(blockchainID).Paginate(page, limit).OrderCreatedAt("DESC").ExecuteWithCount(ctx)
}
