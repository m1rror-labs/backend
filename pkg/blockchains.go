package pkg

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Blockchain struct {
	bun.BaseModel  `bun:"table:blockchains"`
	ID             uuid.UUID  `json:"id,omitempty" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	CreatedAt      time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	AirdropKeypair []byte     `json:"-" bun:",nullzero,notnull"`
	TeamID         uuid.UUID  `json:"team_id" bun:"type:uuid,notnull"`
	Label          *string    `json:"label" bun:",notnull"`
	Expiry         *time.Time `json:"expiry"`
}

type RpcEngine interface {
	CreateBlockchain(ctx context.Context, apiKey uuid.UUID, user_id *string, config *uuid.UUID) (uuid.UUID, error)
	DeleteBlockchain(ctx context.Context, apiKey uuid.UUID, id uuid.UUID) error
	ExpireBlockchains(ctx context.Context) error
}

type BlockchainRepo interface {
	ReadBlockchain() BlockchainReader

	UpdateBlockchain(id uuid.UUID) BlockchainUpdater
}

type BlockchainReader interface {
	Execute(ctx context.Context) ([]Blockchain, error)
	ExecuteOne(ctx context.Context) (Blockchain, error)

	ID(id uuid.UUID) BlockchainReader
	TeamID(teamID uuid.UUID) BlockchainReader
	Label(label *string) BlockchainReader
}

type BlockchainUpdater interface {
	Execute(ctx context.Context) error

	Label(label *string) BlockchainUpdater
}
