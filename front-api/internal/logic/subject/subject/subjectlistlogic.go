package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubjectListLogic {
	return SubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectListLogic) SubjectList(req types.MainTeacherSubjectListReq) (*types.MainTeacherSubjectListResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}


	resp, err := l.svcCtx.Su.SubjectList(l.ctx, &suclient.MainTeacherSubjectListReq{
		Id:       req.Id,
		Offset:   req.Offset,
		PageSize: req.PageSize,
	})

	if err != nil {
		return &types.MainTeacherSubjectListResp{}, err
	}

	data := types.MainTeacherSubjectListData{
		Total: resp.Total,
	}

	for _, v := range resp.List {
		subjectInfo := types.SubjectInfo{
			Id:                 v.Id,
			Uuid:               v.Uuid,
			Name:               v.Name,
			Status:             v.Status,
			Code:               v.Code,
			MaxPersion:         v.MaxPersion,
			MainTeacherId:      v.MainTeacherId,
			AssistantTeacherId: v.AssistantTeacherId,
			Introduce:          v.Introduce,
			Backup:             v.Backup,
			CreateTime:         v.CreateTime,
			UpdateTime:         v.UpdateTime,
			DeleteTime:         v.DeleteTime,
		}
		data.SubjectInfos = append(data.SubjectInfos, &subjectInfo)
	}

	return &types.MainTeacherSubjectListResp{
		Code:    0,
		Message: "success",
		Data:    data,
	}, nil
}
