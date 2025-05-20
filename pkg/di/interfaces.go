package di

import "go/adv-api/internal/user"

type IStatRepository interface {
	AddClick(linkid uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
