package logic

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go-zero-admin/rpc/model/papermodel"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSignalChoiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSignalChoiceLogic {
	return &UpdateSignalChoiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSignalChoiceLogic) UpdateSignalChoice(in *su.UpdateSignalChoiceReq) (*su.UpdateSignalChoiceResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.SignalChoiceModel.Update(l.ctx, &papermodel.SignalChoice{
		ID:            bson.ObjectIdHex(in.Id),
		Title:         in.Title,
		AAnswer:       in.AAnswer,
		BAnswer:       in.BAnswer,
		CAnswer:       in.CAnswer,
		DAnswer:       in.DAnswer,
		EAnswer:       in.EAnswer,
		FAnswer:       in.FAnswer,
		CorrectAnswer: in.CorrectAnswer,
		Version:       in.Version,
	})
	if err != nil {
		return nil, err
	}

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

	return &su.UpdateSignalChoiceResp{
		Result:       true,
		SignalChoice: &info,
	}, nil
}
