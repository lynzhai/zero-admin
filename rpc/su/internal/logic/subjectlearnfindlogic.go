package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/stores/sqlc"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectLearnFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectLearnFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectLearnFindLogic {
	return &SubjectLearnFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectLearnFindLogic) SubjectLearnFind(in *su.SubjectLearnFindReq) (*su.SubjectLearnFindResp, error) {
	// todo: add your logic here and delete this line

	suSubjectLearn, err := l.svcCtx.SuSubjectLearnModel.FindOne(in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errorSubjectNotFound
	default:
		return nil, err
	}

	return &su.SubjectLearnFindResp{
		SubjectLearnInfoData: &su.SubjectLearnInfo{
			Id:         suSubjectLearn.Id,
			UserId:     suSubjectLearn.UserId,
			SubjectId:  suSubjectLearn.SubjectId,
			CreateTime: suSubjectLearn.CreateTime.Time.String(),
		},
	}, nil
}
