package logic

import (
	"context"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteSelfPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSelfPaperLogic {
	return &DeleteSelfPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSelfPaperLogic) DeleteSelfPaper(in *su.DeleteSelfPaperReq) (*su.DeleteSelfPaperResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.SelfPaperModel.DeleteSoft(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &su.DeleteSelfPaperResp{
		Result: true,
	}, nil
}
