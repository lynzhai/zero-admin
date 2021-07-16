package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindSubjectsByStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSubjectsByStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindSubjectsByStudentLogic {
	return FindSubjectsByStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSubjectsByStudentLogic) FindSubjectsByStudent(req types.FindSubjectsByStudentReq) (*types.FindSubjectsByStudentResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}

	resp, err := l.svcCtx.Su.FindSubjectByStudent(l.ctx, &suclient.FindSubjectByStudentReq{
		UserId:   req.UserId,
		Offset:   req.Offset,
		PageSize: req.PageSize,
	})
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return &types.FindSubjectsByStudentResp{}, err
	}

	subjectInfos := make([]*types.SubjectInfo, 0)
	for _, v := range resp.List {
		info := types.SubjectInfo{
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
		}
		subjectInfos = append(subjectInfos, &info)
	}

	return &types.FindSubjectsByStudentResp{
		Code:    0,
		Message: "success",
		Data: types.FindSubjectsByStudentRespListData{
			SubjectInfos: subjectInfos,
			Total:        resp.Total,
		},
	}, nil
}
