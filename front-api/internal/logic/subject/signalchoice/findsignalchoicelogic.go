package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindSignalChoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindSignalChoiceLogic {
	return FindSignalChoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSignalChoiceLogic) FindSignalChoice(req types.FindSignalChoiceReq) (*types.FindSignalChoiceResp, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.Validate.Struct(&req); err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.Su.FindSignalChoice(l.ctx, &suclient.FindSignalChoiceReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.FindSignalChoiceResp{
		Code:    0,
		Message: "",
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
