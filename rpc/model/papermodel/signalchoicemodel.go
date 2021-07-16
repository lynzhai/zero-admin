package papermodel

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
	cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

var prefixSignalChoiceCacheKey = "cache:SignalChoice:"

type SignalChoiceModel interface {
	Insert(ctx context.Context, data *SignalChoice) error
	FindOne(ctx context.Context, id string) (*SignalChoice, error)
	Update(ctx context.Context, data *SignalChoice) error
	Delete(ctx context.Context, id string) error
	DeleteSoft(ctx context.Context, id string) error
}

type defaultSignalChoiceModel struct {
	*mongoc.Model
}

func NewSignalChoiceModel(url, collection string, c cachec.CacheConf) SignalChoiceModel {
	return &defaultSignalChoiceModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultSignalChoiceModel) Insert(ctx context.Context, data *SignalChoice) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	createTime := time.Now()
	updateTime := time.Now()
	data.CreateTime = createTime
	data.UpdateTime = updateTime

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultSignalChoiceModel) FindOne(ctx context.Context, id string) (*SignalChoice, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data SignalChoice
	key := prefixSignalChoiceCacheKey + id
	//err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	err = m.GetCollection(session).FindOne(&data, key, bson.M{"_id": bson.ObjectIdHex(id), "deleted": false})
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSignalChoiceModel) Update(ctx context.Context, data *SignalChoice) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixSignalChoiceCacheKey + data.ID.Hex()
	return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *defaultSignalChoiceModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixSignalChoiceCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}

func (m *defaultSignalChoiceModel) DeleteSoft(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}
	hexId := bson.ObjectIdHex(id)
	defer m.PutSession(session)
	key := prefixSignalChoiceCacheKey + hexId.Hex()
	return m.GetCollection(session).Update(bson.M{"_id": hexId},
		bson.M{"$set": bson.M{
			"deleted": true,
		}}, key)

}
