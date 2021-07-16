package logic

import (
	"context"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteSignalChoiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSignalChoiceLogic {
	return &DeleteSignalChoiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSignalChoiceLogic) DeleteSignalChoice(in *su.DeleteSignalChoiceReq) (*su.DeleteSignalChoiceResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.SignalChoiceModel.DeleteSoft(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &su.DeleteSignalChoiceResp{
		Result: true,
	}, nil
}
