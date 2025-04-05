package blockchains

import (
	"context"
	"mirror-backend/pkg"

	"github.com/google/uuid"
)

func SetMainnetAccountState(
	ctx context.Context,
	rpcEngine pkg.RpcEngine,
	accountRetriever pkg.AccountRetriever,
	blockchainID uuid.UUID,
	accounts []string,
	label *string,
) error {
	// TODO: Billing stuff
	accountsData, err := accountRetriever.GetMultipleAccounts(ctx, accounts)
	if err != nil {
		return err
	}

	if err := rpcEngine.SetAccounts(ctx, blockchainID, accountsData, label); err != nil {
		return pkg.ErrSettingAccount
	}
	return nil
}
