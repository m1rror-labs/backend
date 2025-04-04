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
	account string,
	label *string,
) error {
	// TODO: Billing stuff
	accountData, err := accountRetriever.GetAccount(ctx, account)
	if err != nil {
		return err
	}

	if err := rpcEngine.SetAccount(ctx, blockchainID, accountData, label); err != nil {
		return pkg.ErrSettingAccount
	}
	return nil
}
