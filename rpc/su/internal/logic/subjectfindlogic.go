package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectFindLogic {
	return &SubjectFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectFindLogic) SubjectFind(in *su.SubjectFindReq) (*su.SubjectFindResp, error) {
	// todo: add your logic here and delete this line

	subject, err := l.svcCtx.SuSubjectModel.FindOne(in.SubjectId)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errorSubjectNotFound
	default:
		return nil, err
	}

	return &su.SubjectFindResp{
		SubjectInfoData: &su.SubjectInfo{
			Id:                 subject.Id,
			Uuid:               subject.Uuid,
			Name:               subject.Name.String,
			Status:             subject.Status.Int64,
			Code:               subject.Code,
			MaxPersion:         subject.MaxPersion.Int64,
			MainTeacherId:      subject.Id,
			AssistantTeacherId: subject.Id,
			Introduce:          subject.Introduce.String,
			Backup:             subject.Backup.String,
			CreateTime:         subject.CreateTime.Time.String(),
			UpdateTime:         subject.UpdateTime.Time.String(),
			DeleteTime:         subject.DeleteTime.Time.String(),
		},
	}, nil
}
