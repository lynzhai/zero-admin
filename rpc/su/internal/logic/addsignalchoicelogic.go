package logic

import (
	"context"
	"go-zero-admin/rpc/model/papermodel"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddSignalChoiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSignalChoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSignalChoiceLogic {
	return &AddSignalChoiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSignalChoiceLogic) AddSignalChoice(in *su.AddSignalChoiceReq) (*su.AddSignalChoiceResp, error) {
	// todo: add your logic here and delete this line

	signalChoice := papermodel.SignalChoice{
		Id:            bson.NewObjectId(),
		Type:          SignalChoiceType,
		Title:         in.Title,
		AAnswer:       in.AAnswer,
		BAnswer:       in.BAnswer,
		CAnswer:       in.CAnswer,
		DAnswer:       in.DAnswer,
		EAnswer:       in.EAnswer,
		FAnswer:       in.FAnswer,
		CorrectAnswer: in.CorrectAnswer,
		Version:       in.Version,
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
		Deleted:       false,
		DeleteTime:    time.Time{},
	}

	err := l.svcCtx.SignalChoiceModel.Insert(l.ctx, &signalChoice)

	if err != nil {
		return nil, err
	}

	info := su.SignalChoiceInfo{
		Id:            signalChoice.Id.Hex(),
		Type:          signalChoice.Type,
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

	return &su.AddSignalChoiceResp{
		Result:       true,
		SignalChoice: &info,
	}, nil
}
