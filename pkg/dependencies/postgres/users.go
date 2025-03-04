package postgres

import (
	"context"
	"mirror-backend/pkg"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ pkg.UserRepo = &repository{}

type userReader struct {
	selectQuery *bun.SelectQuery
	users       *[]pkg.User
}

func (r *repository) ReadUser() pkg.UserReader {
	var users []pkg.User
	return &userReader{selectQuery: r.db.NewSelect().Model(&users), users: &users}
}

func (r *userReader) Execute(ctx context.Context) ([]pkg.User, error) {
	err := r.selectQuery.Scan(ctx)
	return *r.users, err
}

func (r *userReader) ExecuteOne(ctx context.Context) (pkg.User, error) {
	err := r.selectQuery.Limit(1).Scan(ctx)
	if err != nil {
		return pkg.User{}, err
	}
	if len(*r.users) == 0 {
		return pkg.User{}, pkg.ErrNotFound
	}
	user := (*r.users)[0]
	if user.Team.Blockchains == nil {
		user.Team.Blockchains = []pkg.Blockchain{}
	}
	if user.Team.ApiKeys == nil {
		user.Team.ApiKeys = []pkg.ApiKey{}
	}
	return user, err
}

func (r *userReader) ID(id uuid.UUID) pkg.UserReader {
	r.selectQuery = r.selectQuery.Where("u.id = ?", id)
	return r
}

func (r *userReader) Email(email string) pkg.UserReader {
	r.selectQuery = r.selectQuery.Where("email = ?", email)
	return r
}

func (r *userReader) TeamID(teamID uuid.UUID) pkg.UserReader {
	r.selectQuery = r.selectQuery.Where("team_id = ?", teamID)
	return r
}

func (r *userReader) WithTeam() pkg.UserReader {
	r.selectQuery = r.selectQuery.Relation("Team")
	return r
}

func (r *userReader) WithApiKeys() pkg.UserReader {
	r.selectQuery = r.selectQuery.Relation("Team.ApiKeys", func(sq *bun.SelectQuery) *bun.SelectQuery {
		return sq.OrderExpr("created_at DESC")
	})
	return r
}

func (r *userReader) WithBlockchains() pkg.UserReader {
	r.selectQuery = r.selectQuery.Relation("Team.Blockchains", func(sq *bun.SelectQuery) *bun.SelectQuery {
		return sq.OrderExpr("created_at DESC")
	})
	return r
}

func (r *repository) CreateUser(ctx context.Context, user *pkg.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *repository) UpdateUser(ctx context.Context, user *pkg.User) error {
	_, err := r.db.NewUpdate().Model(user).WherePK().Exec(ctx)
	return err
}

func (r *repository) DeleteUser(ctx context.Context, user *pkg.User) error {
	_, err := r.db.NewDelete().Model(user).WherePK().Exec(ctx)
	return err
}

type teamReader struct {
	selectQuery *bun.SelectQuery
	teams       *[]pkg.Team
}

func (r *repository) ReadTeam() pkg.TeamReader {
	var teams []pkg.Team
	return &teamReader{selectQuery: r.db.NewSelect().Model(&teams), teams: &teams}
}

func (r *teamReader) Execute(ctx context.Context) ([]pkg.Team, error) {
	err := r.selectQuery.Scan(ctx)
	return *r.teams, err
}

func (r *teamReader) ExecuteOne(ctx context.Context) (pkg.Team, error) {
	err := r.selectQuery.Limit(1).Scan(ctx)
	if err != nil {
		return pkg.Team{}, err
	}
	if len(*r.teams) == 0 {
		return pkg.Team{}, pkg.ErrNotFound
	}
	return (*r.teams)[0], nil
}

func (r *teamReader) ID(id uuid.UUID) pkg.TeamReader {
	r.selectQuery = r.selectQuery.Where("id = ?", id)
	return r
}

func (r *teamReader) WithUsers() pkg.TeamReader {
	r.selectQuery = r.selectQuery.Relation("Users")
	return r
}

func (r *teamReader) WithApiKeys() pkg.TeamReader {
	r.selectQuery = r.selectQuery.Relation("ApiKeys")
	return r
}

func (r *repository) CreateTeam(ctx context.Context, team *pkg.Team) error {
	_, err := r.db.NewInsert().Model(team).Exec(ctx)
	return err
}

func (r *repository) UpdateTeam(ctx context.Context, team *pkg.Team) error {
	_, err := r.db.NewUpdate().Model(team).WherePK().Exec(ctx)
	return err
}

func (r *repository) DeleteTeam(ctx context.Context, team *pkg.Team) error {
	_, err := r.db.NewDelete().Model(team).WherePK().Exec(ctx)
	return err
}

type apiKeyReader struct {
	selectQuery *bun.SelectQuery
	apiKeys     *[]pkg.ApiKey
}

func (r *repository) ReadApiKey() pkg.ApiKeyReader {
	var apiKeys []pkg.ApiKey
	return &apiKeyReader{selectQuery: r.db.NewSelect().Model(&apiKeys), apiKeys: &apiKeys}
}

func (r *apiKeyReader) Execute(ctx context.Context) ([]pkg.ApiKey, error) {
	err := r.selectQuery.Scan(ctx)
	return *r.apiKeys, err
}

func (r *apiKeyReader) ExecuteOne(ctx context.Context) (pkg.ApiKey, error) {
	err := r.selectQuery.Limit(1).Scan(ctx)
	if err != nil {
		return pkg.ApiKey{}, err
	}
	if len(*r.apiKeys) == 0 {
		return pkg.ApiKey{}, pkg.ErrNotFound
	}
	return (*r.apiKeys)[0], err
}

func (r *apiKeyReader) ID(id uuid.UUID) pkg.ApiKeyReader {
	r.selectQuery = r.selectQuery.Where("id = ?", id)
	return r
}

func (r *apiKeyReader) TeamID(teamID uuid.UUID) pkg.ApiKeyReader {
	r.selectQuery = r.selectQuery.Where("team_id = ?", teamID)
	return r
}

func (r *repository) CreateApiKey(ctx context.Context, apiKey *pkg.ApiKey) error {
	_, err := r.db.NewInsert().Model(apiKey).Exec(ctx)
	return err
}

func (r *repository) UpdateApiKey(ctx context.Context, apiKey *pkg.ApiKey) error {
	_, err := r.db.NewUpdate().Model(apiKey).WherePK().Exec(ctx)
	return err
}

func (r *repository) DeleteApiKey(ctx context.Context, apiKey *pkg.ApiKey) error {
	_, err := r.db.NewDelete().Model(apiKey).WherePK().Exec(ctx)
	return err
}
