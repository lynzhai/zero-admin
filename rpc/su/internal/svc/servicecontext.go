package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin/rpc/model/papermodel"
	"go-zero-admin/rpc/model/sumodel"
	"go-zero-admin/rpc/su/internal/config"
	"go-zero-admin/rpc/us/usclient"
)

type ServiceContext struct {
	Config              config.Config
	SuSubjectModel      sumodel.SuSubjectModel
	SuSubjectLearnModel sumodel.SuSubjectLearnModel
	RedisConn           *redis.Redis
	Us                  usclient.Us
	SelfPaperModel      papermodel.SelfPaperModel
	AnswerPaperModel    papermodel.AnswerPaperModel
	SignalChoiceModel   papermodel.SignalChoiceModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	suSubjectModel := sumodel.NewSuSubjectModel(conn, c.CacheRedis)
	suSubjectLearnModel := sumodel.NewSuSubjectLearnModel(conn, c.CacheRedis)
	rcon := redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type)
	selfPaperModel := papermodel.NewSelfPaperModel(c.Mongo.Url, "SelfPaper", c.CacheRedis)
	answerPaperModel := papermodel.NewAnswerPaperModel(c.Mongo.Url, "AnswerPaper", c.CacheRedis)
	signalChoiceModel := papermodel.NewSignalChoiceModel(c.Mongo.Url, "SignalChoice", c.CacheRedis)

	return &ServiceContext{
		Config:              c,
		SuSubjectModel:      suSubjectModel,
		SuSubjectLearnModel: suSubjectLearnModel,
		RedisConn:           rcon,
		Us:                  usclient.NewUs(zrpc.MustNewClient(c.UsRpc)),
		SelfPaperModel:      selfPaperModel,
		AnswerPaperModel:    answerPaperModel,
		SignalChoiceModel:   signalChoiceModel,
	}
}
