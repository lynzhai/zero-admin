package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteSelfPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteSelfPaperLogic {
	return DeleteSelfPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSelfPaperLogic) DeleteSelfPaper(req types.DeleteSelfPaperReq) (*types.DeleteSelfPaperResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.Su.DeleteSelfPaper(l.ctx, &suclient.DeleteSelfPaperReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteSelfPaperResp{
		Code:    0,
		Message: "success",
		Data: types.DeleteSelfPaperRespData{
			Result: resp.Result,
		},
	}, nil
}
