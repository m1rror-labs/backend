package users

import (
	"context"
	"mirror-backend/pkg"
	"mirror-backend/pkg/dependencies/randomtext"
)

func GetUserSelf(ctx context.Context, userRepo pkg.UserRepo, email string) (pkg.User, error) {
	user, err := userRepo.ReadUser().Email(email).WithApiKeys().WithBlockchains().ExecuteOne(ctx)
	if err != nil {
		if err == pkg.ErrNotFound {
			u, err := CreateUserAndTeam(ctx, userRepo, email)
			if err != nil {
				return pkg.User{}, err
			}
			return u, nil
		}
		return pkg.User{}, err
	}
	return user, nil
}

func CreateUserAndTeam(ctx context.Context, userRepo pkg.UserRepo, email string) (pkg.User, error) {
	team := pkg.Team{
		Name: randomtext.GenerateRandomText(),
	}
	if err := userRepo.CreateTeam(ctx, &team); err != nil {
		return pkg.User{}, err
	}

	user := pkg.User{
		Email:  email,
		TeamID: team.ID,
	}
	if err := userRepo.CreateUser(ctx, &user); err != nil {
	}

	return user, nil
}
