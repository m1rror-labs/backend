package blockchains

import (
	"context"
	"mirror-backend/pkg"

	"github.com/google/uuid"
)

func CreateBlockchain(ctx context.Context, rpcEngine pkg.RpcEngine, blockchainRepo pkg.BlockchainRepo, user pkg.User) (pkg.Blockchain, error) {
	if len(user.Team.ApiKeys) == 0 {
		return pkg.Blockchain{}, pkg.ErrNoApiKey
	}
	apiKey := user.Team.ApiKeys[0].ID
	blockchainID, err := rpcEngine.CreateBlockchain(ctx, apiKey)
	if err != nil {
		return pkg.Blockchain{}, err
	}

	return blockchainRepo.ReadBlockchain().ID(blockchainID).ExecuteOne(ctx)
}

func DeleteBlockchain(ctx context.Context, rpcEngine pkg.RpcEngine, blockchainRepo pkg.BlockchainRepo, user pkg.User, blockchainID uuid.UUID) error {
	if len(user.Team.ApiKeys) == 0 {
		return pkg.ErrNoApiKey
	}
	apiKey := user.Team.ApiKeys[0].ID
	return rpcEngine.DeleteBlockchain(ctx, apiKey, blockchainID)
}
