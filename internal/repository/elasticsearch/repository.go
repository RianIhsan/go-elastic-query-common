package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/RianIhsan/go-elastic-query-common/internal/domain"
	"github.com/olivere/elastic/v7"
)

type UserRepository struct {
	client *elastic.Client
}

func NewUserRepository(client *elastic.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Save(user domain.User) error {
	_, err := r.client.Index().
		Index("users").
		Id(user.ID).
		BodyJson(user).
		Do(context.Background())

	return err
}

func (r *UserRepository) FindByID(id string) (domain.User, error) {
	res, err := r.client.Get().
		Index("users").
		Id(id).
		Do(context.Background())
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = json.Unmarshal(res.Source, &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
