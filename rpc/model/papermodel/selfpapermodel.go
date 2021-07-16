package papermodel

import (
	"context"

	"github.com/globalsign/mgo/bson"
	cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

var prefixSelfPaperCacheKey = "cache:SelfPaper:"

type SelfPaperModel interface {
	Insert(ctx context.Context, data *SelfPaper) error
	FindOne(ctx context.Context, id string) (*SelfPaper, error)
	Update(ctx context.Context, data *SelfPaper) error
	Delete(ctx context.Context, id string) error
	DeleteSoft(ctx context.Context, id string) error
}

type defaultSelfPaperModel struct {
	*mongoc.Model
}

func NewSelfPaperModel(url, collection string, c cachec.CacheConf) SelfPaperModel {
	return &defaultSelfPaperModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultSelfPaperModel) Insert(ctx context.Context, data *SelfPaper) error {
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

func (m *defaultSelfPaperModel) FindOne(ctx context.Context, id string) (*SelfPaper, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data SelfPaper
	key := prefixSelfPaperCacheKey + id
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

func (m *defaultSelfPaperModel) Update(ctx context.Context, data *SelfPaper) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixSelfPaperCacheKey + data.ID.Hex()
	return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *defaultSelfPaperModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixSelfPaperCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}

func (m *defaultSelfPaperModel) DeleteSoft(ctx context.Context, id string) error {
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
