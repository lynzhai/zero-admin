package mongomodel

import (
	"context"

	"github.com/globalsign/mgo/bson"
	cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

type UserModel interface {
	Insert(ctx context.Context, data *User) error
	FindOne(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, data *User) error
	Delete(ctx context.Context, id string) error
}

type defaultUserModel struct {
	*mongoc.Model
}

func NewUserModel(url, collection string, c cachec.CacheConf) UserModel {
	return &defaultUserModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data User

	err = m.GetCollection(session).FindOneIdNoCache(&data, bson.ObjectIdHex(id))
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateIdNoCache(data.ID, data)
}

func (m *defaultUserModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveIdNoCache(bson.ObjectIdHex(id))
}
