package logic

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go-zero-admin/rpc/model/papermodel"
	"time"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSelfPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSelfPaperLogic {
	return &UpdateSelfPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSelfPaperLogic) UpdateSelfPaper(in *su.UpdateSelfPaperReq) (*su.UpdateSelfPaperResp, error) {
	// todo: add your logic here and delete this line

	//客户端只能传utc时间
	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	stopTime, _ := time.Parse("2006-01-02 15:04:05", in.StopTime)

	selfPaper := papermodel.SelfPaper{
		Id:             bson.ObjectIdHex(in.Id),
		PaperName:      in.PaperName,
		CreaterId:      in.CreaterId,
		Status:         in.Status,
		StartTime:      startTime,
		StopTime:       stopTime,
		Version:        in.Version,
		PaperItems:     paperItems_ProtoToModel(in.PaperItems),
		RandomSettings: randomSetting_ProtoToModel(in.RandomSettings),
	}

	err := l.svcCtx.SelfPaperModel.Update(l.ctx, &selfPaper)

	if err != nil {
		return nil, err
	}

	return &su.UpdateSelfPaperResp{
		Result: true,
	}, nil
}
