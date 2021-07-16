package papermodel

import (
	"context"

	"github.com/globalsign/mgo/bson"
	cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

var prefixAnswerPaperCacheKey = "cache:AnswerPaper:"

type AnswerPaperModel interface {
	Insert(ctx context.Context, data *AnswerPaper) error
	FindOne(ctx context.Context, id string) (*AnswerPaper, error)
	Update(ctx context.Context, data *AnswerPaper) error
	Delete(ctx context.Context, id string) error
	DeleteSoft(ctx context.Context, id string) error
}

type defaultAnswerPaperModel struct {
	*mongoc.Model
}

func NewAnswerPaperModel(url, collection string, c cachec.CacheConf) AnswerPaperModel {
	return &defaultAnswerPaperModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultAnswerPaperModel) Insert(ctx context.Context, data *AnswerPaper) error {
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

func (m *defaultAnswerPaperModel) FindOne(ctx context.Context, id string) (*AnswerPaper, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data AnswerPaper
	key := prefixAnswerPaperCacheKey + id
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

func (m *defaultAnswerPaperModel) Update(ctx context.Context, data *AnswerPaper) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixAnswerPaperCacheKey + data.ID.Hex()
	return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *defaultAnswerPaperModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixAnswerPaperCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}

func (m *defaultAnswerPaperModel) DeleteSoft(ctx context.Context, id string) error {
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
