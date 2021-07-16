package logic

import (
	"context"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectLearnDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectLearnDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectLearnDeleteLogic {
	return &SubjectLearnDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectLearnDeleteLogic) SubjectLearnDelete(in *su.SubjectLearnDeleteReq) (*su.SubjectLearnDeleteResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.SuSubjectLearnModel.Delete(in.Id)
	if err != nil {
		logx.Errorf("DeleteSoft err:" + err.Error())
		return &su.SubjectLearnDeleteResp{
			Result: false,
		}, err

	}

	return &su.SubjectLearnDeleteResp{Result: true}, nil
}
