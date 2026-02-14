package repo

import (
	"context"
	"errors"

	"github.com/zjutjh/mygo/ndb"
	"gorm.io/gorm"

	"app/dao/model"
	"app/dao/query"
)

type UserRepo struct {
	query *query.Query
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		query: query.Use(ndb.Pick()),
	}
}

// FindById 根据ID查询用户
func (r *UserRepo) FindById(ctx context.Context, id int64) (*model.User, error) {
	u := r.query.User
	record, err := u.WithContext(ctx).Where(u.ID.Eq(id)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return record, nil
}
