package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/sumodel"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectUpdateLogic {
	return &SubjectUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectUpdateLogic) SubjectUpdate(in *su.SubjectUpdateReq) (*su.SubjectUpdateResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.SuSubjectModel.Update(sumodel.SuSubject{
		Id:   in.Id,
		Uuid: in.Uuid,
		Name: sql.NullString{
			String: in.Name,
			Valid:  true,
		},
		Status: sql.NullInt64{
			Int64: in.Status,
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
	})

	if err != nil {
		return &su.SubjectUpdateResp{}, err
	}

	suSubject, err := l.svcCtx.SuSubjectModel.FindOne(in.Id)

	if err != nil {
		return &su.SubjectUpdateResp{}, err
	}

	return &su.SubjectUpdateResp{
		Result: true,
		SubjectInfoData: &su.SubjectInfo{
			Id:                 suSubject.Id,
			Uuid:               suSubject.Uuid,
			Name:               suSubject.Name.String,
			Status:             suSubject.Status.Int64,
			Code:               suSubject.Code,
			MaxPersion:         suSubject.MaxPersion.Int64,
			MainTeacherId:      suSubject.MainTeacherId,
			AssistantTeacherId: suSubject.AssistantTeacherId.Int64,
			Introduce:          suSubject.Introduce.String,
			Backup:             suSubject.Backup.String,
			CreateTime:         suSubject.CreateTime.Time.String(),
			UpdateTime:         suSubject.UpdateTime.Time.String(),
		},
	}, err
}
