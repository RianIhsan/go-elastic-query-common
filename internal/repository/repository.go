package repository

import "github.com/RianIhsan/go-elastic-query-common/internal/domain"

type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (domain.User, error)
}
