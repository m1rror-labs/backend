package postgres

import (
	"context"
	"mirror-backend/pkg"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ pkg.AccountRepo = &repository{}

type accountReader struct {
	selectQuery            *bun.SelectQuery
	transactionLogMessages *[]pkg.Account
}

func (r *repository) ReadAccounts() pkg.AccountReader {
	var transactionLogMessages []pkg.Account = []pkg.Account{}
	return &accountReader{selectQuery: r.db.NewSelect().Model(&transactionLogMessages), transactionLogMessages: &transactionLogMessages}
}

func (r *accountReader) Execute(ctx context.Context) ([]pkg.Account, error) {
	err := r.selectQuery.Scan(ctx)
	return *r.transactionLogMessages, err
}

func (r *accountReader) ExecuteWithCount(ctx context.Context) ([]pkg.Account, int, error) {
	count, err := r.selectQuery.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	transactionLogMessages, err := r.Execute(ctx)
	return transactionLogMessages, count, err
}

func (r *accountReader) ExecuteOne(ctx context.Context) (pkg.Account, error) {
	err := r.selectQuery.Limit(1).Scan(ctx)
	if err != nil {
		return pkg.Account{}, err
	}
	if len(*r.transactionLogMessages) == 0 {
		return pkg.Account{}, pkg.ErrNotFound
	}
	return (*r.transactionLogMessages)[0], err
}

func (r *accountReader) BlockchainID(blockchainID uuid.UUID) pkg.AccountReader {
	r.selectQuery = r.selectQuery.
		Where("blockchain = ?", blockchainID)
	return r
}

func (r *accountReader) Paginate(page int, limit int) pkg.AccountReader {
	offset := (page - 1) * limit
	r.selectQuery = r.selectQuery.Offset(offset).Limit(limit)
	return r
}

func (r *accountReader) OrderCreatedAt(order string) pkg.AccountReader {
	r.selectQuery = r.selectQuery.Order("account.created_at " + order)
	return r
}

func (r *accountReader) Between(start, end time.Time) pkg.AccountReader {
	r.selectQuery = r.selectQuery.Where("account.created_at BETWEEN ? AND ?", start, end)
	return r
}
