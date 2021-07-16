package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindStudentsBySubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindStudentsBySubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindStudentsBySubjectLogic {
	return FindStudentsBySubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindStudentsBySubjectLogic) FindStudentsBySubject(req types.FindStudentsBySubjectReq) (*types.FindStudentsBySubjectResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}

	resp, err := l.svcCtx.Su.FindStudentsInSubject(l.ctx, &suclient.FindStudentsInSubjectReq{
		SubjectId: req.SubjectId,
		Offset:    req.Offset,
		PageSize:  req.PageSize,
	})
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return &types.FindStudentsBySubjectResp{}, err
	}

	persionInfos := make([]*types.PersionInfo, 0)
	for _, v := range resp.List {
		info := types.PersionInfo{
			Id:          v.Id,
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			RoleName:    v.RoleType,
			RoleId:      v.RoleId,
			Gender:      v.Gender,
			Class:       v.Class,
			Academy:     v.Academy,
			School:      v.School,
			Grade:       v.Grade,
		}
		persionInfos = append(persionInfos, &info)
	}

	return &types.FindStudentsBySubjectResp{
		Code:    0,
		Message: "success",
		Data: types.FindStudentsBySubjectRespListData{
			PersionInfos: persionInfos,
			Total:        resp.Total,
		},
	}, nil
}
