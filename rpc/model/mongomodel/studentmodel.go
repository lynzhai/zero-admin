package mongomodel

import (
	"context"

	"github.com/globalsign/mgo/bson"
	cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

var prefixStudentCacheKey = "cache:Student:"

type StudentModel interface {
	Insert(ctx context.Context, data *Student) error
	FindOne(ctx context.Context, id string) (*Student, error)
	Update(ctx context.Context, data *Student) error
	Delete(ctx context.Context, id string) error
}

type defaultStudentModel struct {
	*mongoc.Model
}

func NewStudentModel(url, collection string, c cachec.CacheConf) StudentModel {
	return &defaultStudentModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultStudentModel) Insert(ctx context.Context, data *Student) error {
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

func (m *defaultStudentModel) FindOne(ctx context.Context, id string) (*Student, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data Student
	key := prefixStudentCacheKey + id
	err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStudentModel) Update(ctx context.Context, data *Student) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixStudentCacheKey + data.ID.Hex()
	return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *defaultStudentModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixStudentCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}
