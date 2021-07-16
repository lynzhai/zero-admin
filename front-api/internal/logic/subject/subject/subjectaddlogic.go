package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubjectAddLogic {
	return SubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectAddLogic) SubjectAdd(req types.SubjectAddReq) (*types.SubjectAddResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}

	resp, err := l.svcCtx.Su.SubjectAdd(l.ctx, &suclient.SubjectAddReq{
		Name:               req.Name,
		Status:             req.Status,
		Code:               req.Code,
		MaxPersion:         req.MaxPersion,
		MainTeacherId:      req.MainTeacherId,
		AssistantTeacherId: req.AssistantTeacherId,
		Introduce:          req.Introduce,
		Backup:             req.Backup,
	})

	if err != nil {
		logx.Errorf("SubjectAdd error:" + err.Error())
		return &types.SubjectAddResp{}, err
	}
	return &types.SubjectAddResp{
		Code:    0,
		Message: "success",
		Data: types.SubjectInfo{
			Id:                 resp.SubjectInfoData.Id,
			Uuid:               resp.SubjectInfoData.Uuid,
			Name:               resp.SubjectInfoData.Name,
			Status:             resp.SubjectInfoData.Status,
			Code:               resp.SubjectInfoData.Code,
			MaxPersion:         resp.SubjectInfoData.MaxPersion,
			MainTeacherId:      resp.SubjectInfoData.MainTeacherId,
			AssistantTeacherId: resp.SubjectInfoData.AssistantTeacherId,
			Introduce:          resp.SubjectInfoData.Introduce,
			Backup:             resp.SubjectInfoData.Backup,
			CreateTime:         resp.SubjectInfoData.CreateTime,
			UpdateTime:         resp.SubjectInfoData.UpdateTime,
			DeleteTime:         resp.SubjectInfoData.DeleteTime,
		},
	}, nil
}
