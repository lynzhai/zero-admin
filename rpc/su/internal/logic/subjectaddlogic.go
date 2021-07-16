package logic

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/util/guid"
	"go-zero-admin/rpc/model/sumodel"
	"time"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectAddLogic {
	return &SubjectAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectAddLogic) SubjectAdd(in *su.SubjectAddReq) (*su.SubjectAddResp, error) {
	// todo: add your logic here and delete this line

	if len(in.Code) == 0 {
		return nil, errorCodeEmpty
	}
	_, err := l.svcCtx.SuSubjectModel.FindOneByCode(in.Code)
	if err == nil {
		return &su.SubjectAddResp{
			Result:          false,
			SubjectInfoData: nil,
		}, errorDuplicateCode
	}

	l.svcCtx.SuSubjectModel.DeleteCodeCache(in.Code)

	uniqueId := guid.S()
	result, err := l.svcCtx.SuSubjectModel.Insert(sumodel.SuSubject{
		Uuid: uniqueId,
		Name: sql.NullString{
			String: in.Name,
			Valid:  true,
		},
		Status: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		Code: in.Code,
		MaxPersion: sql.NullInt64{
			Int64: in.MaxPersion,
			Valid: true,
		},
		MainTeacherId: in.MainTeacherId,
		AssistantTeacherId: sql.NullInt64{
			Int64: in.AssistantTeacherId,
			Valid: true,
		},
		Introduce: sql.NullString{
			String: in.Introduce,
			Valid:  true,
		},
		Backup: sql.NullString{
			String: in.Backup,
			Valid:  true,
		},
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return &su.SubjectAddResp{
			Result:          false,
			SubjectInfoData: nil,
		}, errorSubjectRegisterFail
	}
	insertedId, _ := result.LastInsertId()

	lastSubject, err := l.svcCtx.SuSubjectModel.FindOne(insertedId)

	if err != nil {
		return &su.SubjectAddResp{
			Result:          false,
			SubjectInfoData: nil,
		}, errorSubjectRegisterFail
	}

	err = AddSubjectId(lastSubject.MainTeacherId, lastSubject.Id, lastSubject.CreateTime.Time, l.svcCtx.RedisConn)
	if err != nil {
		logx.Errorf("AddSubjectId error:" + err.Error())
	}
	count, err := l.svcCtx.SuSubjectModel.Count(lastSubject.MainTeacherId)
	if err != nil {
		DelSubjectCountCache(lastSubject.MainTeacherId, l.svcCtx.RedisConn)
	} else {
		SetSubjectTotalNumberCache(lastSubject.MainTeacherId, count, l.svcCtx.RedisConn)
		//SetSubjectTotalNumberCacheExpire(lastSubject.MainTeacherId,common.GetCommonRedisExpireSeconds(),l.svcCtx.RedisConn)
	}

	return &su.SubjectAddResp{
		Result: true,
		SubjectInfoData: &su.SubjectInfo{
			Id:                 lastSubject.Id,
			Uuid:               lastSubject.Uuid,
			Name:               lastSubject.Name.String,
			Status:             lastSubject.Status.Int64,
			Code:               lastSubject.Code,
			MaxPersion:         lastSubject.MaxPersion.Int64,
			MainTeacherId:      lastSubject.MainTeacherId,
			AssistantTeacherId: lastSubject.AssistantTeacherId.Int64,
			Introduce:          lastSubject.Introduce.String,
			Backup:             lastSubject.Backup.String,
			CreateTime:         lastSubject.CreateTime.Time.String(),
			UpdateTime:         lastSubject.UpdateTime.Time.String(),
			DeleteTime:         "",
		},
	}, nil
}
