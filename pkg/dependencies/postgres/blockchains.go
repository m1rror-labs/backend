package postgres

import (
	"context"
	"mirror-backend/pkg"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ pkg.BlockchainRepo = &repository{}

type blockchainReader struct {
	selectQuery *bun.SelectQuery
	blockchains *[]pkg.Blockchain
}

func (r *repository) ReadBlockchain() pkg.BlockchainReader {
	var blockchains []pkg.Blockchain
	return &blockchainReader{selectQuery: r.db.NewSelect().Model(&blockchains), blockchains: &blockchains}
}

func (r *blockchainReader) Execute(ctx context.Context) ([]pkg.Blockchain, error) {
	err := r.selectQuery.Scan(ctx)
	return *r.blockchains, err
}

func (r *blockchainReader) ExecuteOne(ctx context.Context) (pkg.Blockchain, error) {
	err := r.selectQuery.Limit(1).Scan(ctx)
	if err != nil {
		return pkg.Blockchain{}, err
	}
	return (*r.blockchains)[0], err
}

func (r *blockchainReader) ID(id uuid.UUID) pkg.BlockchainReader {
	r.selectQuery = r.selectQuery.Where("id = ?", id)
	return r
}

func (r *blockchainReader) TeamID(teamID uuid.UUID) pkg.BlockchainReader {
	r.selectQuery = r.selectQuery.Where("team_id = ?", teamID)
	return r
}
