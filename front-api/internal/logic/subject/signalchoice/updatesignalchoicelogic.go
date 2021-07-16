package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSignalChoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateSignalChoiceLogic {
	return UpdateSignalChoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSignalChoiceLogic) UpdateSignalChoice(req types.UpdateSignalChoiceReq) (*types.UpdateSignalChoiceResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.Su.UpdateSignalChoice(l.ctx, &suclient.UpdateSignalChoiceReq{
		Id:            req.Id,
		Title:         req.Title,
		AAnswer:       req.AAnswer,
		BAnswer:       req.BAnswer,
		CAnswer:       req.CAnswer,
		DAnswer:       req.DAnswer,
		EAnswer:       req.EAnswer,
		FAnswer:       req.FAnswer,
		CorrectAnswer: req.CorrectAnswer,
		Version:       req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateSignalChoiceResp{
		Code:    0,
		Message: "success",
		Data: types.SignalChoiceInfo{
			Id:            resp.SignalChoice.Id,
			Title:         resp.SignalChoice.Title,
			AAnswer:       resp.SignalChoice.AAnswer,
			BAnswer:       resp.SignalChoice.BAnswer,
			CAnswer:       resp.SignalChoice.CAnswer,
			DAnswer:       resp.SignalChoice.DAnswer,
			EAnswer:       resp.SignalChoice.EAnswer,
			FAnswer:       resp.SignalChoice.FAnswer,
			CorrectAnswer: resp.SignalChoice.CorrectAnswer,
			Version:       resp.SignalChoice.Version,
			CreateTime:    resp.SignalChoice.CreateTime,
			UpdateTime:    resp.SignalChoice.UpdateTime,
		},
	}, nil
}
