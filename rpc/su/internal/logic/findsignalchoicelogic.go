package logic

import (
	"context"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindSignalChoiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSignalChoiceLogic {
	return &FindSignalChoiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindSignalChoiceLogic) FindSignalChoice(in *su.FindSignalChoiceReq) (*su.FindSignalChoiceResp, error) {
	// todo: add your logic here and delete this line

	signalChoice, err := l.svcCtx.SignalChoiceModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	info := su.SignalChoiceInfo{
		Id:            signalChoice.ID.Hex(),
		Title:         signalChoice.Title,
		AAnswer:       signalChoice.AAnswer,
		BAnswer:       signalChoice.BAnswer,
		CAnswer:       signalChoice.CAnswer,
		DAnswer:       signalChoice.DAnswer,
		EAnswer:       signalChoice.EAnswer,
		FAnswer:       signalChoice.FAnswer,
		CorrectAnswer: signalChoice.CorrectAnswer,
		Version:       signalChoice.Version,
		CreateTime:    signalChoice.CreateTime.String(),
		UpdateTime:    signalChoice.UpdateTime.String(),
	}

	return &su.FindSignalChoiceResp{
		SignalChoice: &info,
	}, nil
}
