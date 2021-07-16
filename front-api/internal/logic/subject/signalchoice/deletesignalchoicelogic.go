package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteSignalChoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteSignalChoiceLogic {
	return DeleteSignalChoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSignalChoiceLogic) DeleteSignalChoice(req types.DeleteSignalChoiceReq) (*types.DeleteSignalChoiceResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.Su.DeleteSignalChoice(l.ctx, &suclient.DeleteSignalChoiceReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteSignalChoiceResp{
		Code:    0,
		Message: "success",
		Data: types.DeleteSignalChoiceRespData{
			Result: resp.Result,
		},
	}, nil
}
