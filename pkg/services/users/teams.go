package users

import (
	"context"
	"mirror-backend/pkg"
	"mirror-backend/pkg/dependencies/randomtext"

	"github.com/google/uuid"
)

func CreateApiKey(ctx context.Context, userRepo pkg.UserRepo, user pkg.User) error {
	if len(user.Team.ApiKeys) >= 2 {
		return pkg.ErrTooManyApiKeys
	}

	key := pkg.ApiKey{
		TeamID: user.Team.ID,
		Label:  randomtext.GenerateRandomText(),
	}

	return userRepo.CreateApiKey(ctx, &key)
}

func UpdateApiKey(ctx context.Context, userRepo pkg.UserRepo, user pkg.User, newKey pkg.ApiKey) error {
	authorized := false
	for _, key := range user.Team.ApiKeys {
		if key.ID == newKey.ID { // Ensure the key ID is not changed
			newKey.TeamID = key.TeamID       // Ensure the team ID is not changed
			newKey.CreatedAt = key.CreatedAt // Ensure the created at time is not changed
			authorized = true
		}
	}
	if !authorized {
		return pkg.ErrUnauthorized
	}

	return userRepo.UpdateApiKey(ctx, &newKey)
}

func DeleteApiKey(ctx context.Context, userRepo pkg.UserRepo, user pkg.User, id uuid.UUID) error {
	authorized := false
	for _, key := range user.Team.ApiKeys {
		if key.ID == id {
			authorized = true
		}
	}
	if !authorized {
		return pkg.ErrUnauthorized
	}

	key := pkg.ApiKey{ID: id}
	return userRepo.DeleteApiKey(ctx, &key)
}
