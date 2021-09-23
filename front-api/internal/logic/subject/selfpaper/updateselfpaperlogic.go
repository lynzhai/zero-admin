package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSelfPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateSelfPaperLogic {
	return UpdateSelfPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSelfPaperLogic) UpdateSelfPaper(req types.UpdateSelfPaperReq) (*types.UpdateSelfPaperResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.Su.UpdateSelfPaper(l.ctx, &suclient.UpdateSelfPaperReq{
		Id:             req.Id,
		PaperName:      req.PaperName,
		CreaterId:      req.CreaterId,
		Status:         req.Status,

		StartTime:      req.StartTime,
		StopTime:       req.StopTime,
		Version:        req.Version,
		PaperItems:     paperItemsToProto(req.PaperItems),
		RandomSettings: randomSettingsToProto(req.RandomSettings),
	})

	return &types.UpdateSelfPaperResp{
		Code:    0,
		Message: "success",
		Data: types.UpdateSelfPaperRespData{
			Result: resp.Result,
		},
	}, nil
}
