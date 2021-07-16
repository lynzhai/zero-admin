package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubjectFindLogic {
	return SubjectFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectFindLogic) SubjectFind(req types.SubjectFindReq) (*types.SubjectFindResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}


	resp, err := l.svcCtx.Su.SubjectFind(l.ctx, &suclient.SubjectFindReq{
		SubjectId: req.Id,
	})
	if err != nil {
		return &types.SubjectFindResp{}, err
	}
	return &types.SubjectFindResp{
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
