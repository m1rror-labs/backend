package blockchains

import (
	"context"
	"mirror-backend/pkg"

	"github.com/google/uuid"
)

func GetTransactionLogs(
	ctx context.Context,
	transactionRepo pkg.TransactionRepo,
	blockchainID uuid.UUID,
	user pkg.User,
	page,
	limit int,
) ([]pkg.TransactionLogMessage, int, error) {
	return transactionRepo.ReadTransactionLogMessages().TeamID(user.Team.ID).BlockchainID(blockchainID).Paginate(page, limit).OrderCreatedAt("DESC").ExecuteWithCount(ctx)
}
