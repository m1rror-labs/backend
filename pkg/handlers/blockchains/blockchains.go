package blockchains

import (
	"context"
	"mirror-backend/pkg"
	"mirror-backend/pkg/dependencies/randomtext"

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

	label := randomtext.GenerateRandomText()
	if err := blockchainRepo.UpdateBlockchain(blockchainID).Label(&label).Execute(ctx); err != nil {
		return pkg.Blockchain{}, err
	}

	return blockchainRepo.ReadBlockchain().ID(blockchainID).ExecuteOne(ctx)
}

func UpdateBlockchain(ctx context.Context, blockchainRepo pkg.BlockchainRepo, user pkg.User, newBlockchain pkg.Blockchain) error {
	authorized := false
	for _, key := range user.Team.Blockchains {
		if key.ID == newBlockchain.ID { // Ensure the key ID is not changed
			newBlockchain.TeamID = key.TeamID                 // Ensure the team ID is not changed
			newBlockchain.CreatedAt = key.CreatedAt           // Ensure the created at time is not changed
			newBlockchain.AirdropKeypair = key.AirdropKeypair // Ensure the airdrop keypair is not changed
			authorized = true
		}
	}
	if !authorized {
		return pkg.ErrUnauthorized
	}

	return blockchainRepo.UpdateBlockchain(newBlockchain.ID).Label(newBlockchain.Label).Execute(ctx)
}

func DeleteBlockchain(ctx context.Context, rpcEngine pkg.RpcEngine, blockchainRepo pkg.BlockchainRepo, user pkg.User, blockchainID uuid.UUID) error {
	if len(user.Team.ApiKeys) == 0 {
		return pkg.ErrNoApiKey
	}
	apiKey := user.Team.ApiKeys[0].ID
	return rpcEngine.DeleteBlockchain(ctx, apiKey, blockchainID)
}
