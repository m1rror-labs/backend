package pkg

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            uuid.UUID `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	Email         string    `json:"email" bun:",notnull"`
	TeamID        uuid.UUID `json:"team_id" bun:"type:uuid,notnull"`
	Team          *Team     `json:"team" bun:"teams,rel:has-one,join:team_id=id"`
}

type Team struct {
	bun.BaseModel `bun:"table:teams"`
	ID            uuid.UUID    `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	CreatedAt     time.Time    `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	Name          string       `json:"name" bun:",notnull"`
	Users         []User       `json:"users" bun:"rel:has-many,join:id=team_id"`
	ApiKeys       []ApiKey     `json:"api_keys" bun:"rel:has-many,join:id=team_id"`
	Blockchains   []Blockchain `json:"blockchains" bun:"rel:has-many,join:id=team_id"`
}

type ApiKey struct {
	bun.BaseModel `bun:"table:api_keys"`
	ID            uuid.UUID `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	TeamID        uuid.UUID `json:"team_id" bun:"type:uuid,notnull"`
	Team          *Team     `json:"team" bun:"teams,rel:has-one,join:team_id=id"`
	Label         string    `json:"label" bun:",notnull"`
}

type Auth interface {
	User() gin.HandlerFunc
	ApiKey(userRepo UserRepo) gin.HandlerFunc
}

type UserRepo interface {
	ReadUser() UserReader
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, user *User) error

	ReadTeam() TeamReader
	CreateTeam(ctx context.Context, team *Team) error
	UpdateTeam(ctx context.Context, team *Team) error
	DeleteTeam(ctx context.Context, team *Team) error

	ReadApiKey() ApiKeyReader
	CreateApiKey(ctx context.Context, apiKey *ApiKey) error
	UpdateApiKey(ctx context.Context, apiKey *ApiKey) error
	DeleteApiKey(ctx context.Context, apiKey *ApiKey) error
}

type UserReader interface {
	Execute(ctx context.Context) ([]User, error)
	ExecuteOne(ctx context.Context) (User, error)

	ID(id uuid.UUID) UserReader
	Email(email string) UserReader
	TeamID(teamID uuid.UUID) UserReader

	WithTeam() UserReader
	WithApiKeys() UserReader
	WithBlockchains() UserReader
}

type TeamReader interface {
	Execute(ctx context.Context) ([]Team, error)
	ExecuteOne(ctx context.Context) (Team, error)

	ID(id uuid.UUID) TeamReader

	WithUsers() TeamReader
	WithApiKeys() TeamReader
}

type ApiKeyReader interface {
	Execute(ctx context.Context) ([]ApiKey, error)
	ExecuteOne(ctx context.Context) (ApiKey, error)

	ID(id uuid.UUID) ApiKeyReader
	TeamID(teamID uuid.UUID) ApiKeyReader

	WithTeam() ApiKeyReader
}

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ErrUnauthorized   = Err("Unauthorized")
	ErrTooManyApiKeys = Err("Too many api keys")
	ErrHttpRequest    = Err("HTTP request error")
	ErrNoApiKey       = Err("No api key")
	ErrNotFound       = Err("Not found")

	ErrInvalidPubkey   = Err("Invalid pubkey")
	ErrAccountNotFound = Err("Account not found")
)
