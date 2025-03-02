package users

import (
	"context"
	"mirror-backend/pkg"
)

func GetUserSelf(ctx context.Context, userRepo pkg.UserRepo, email string) (pkg.User, error) {
	user, err := userRepo.ReadUser().Email(email).WithApiKeys().WithBlockchains().ExecuteOne(ctx)
	if err != nil {
		return pkg.User{}, err
	}
	return user, nil
}
