package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddStudentInSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStudentInSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddStudentInSubjectLogic {
	return AddStudentInSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStudentInSubjectLogic) AddStudentInSubject(req types.AddStudentInSubjectReq) (*types.AddStudentInSubjectResp, error) {
	// todo: add your logic here and delete this line

	resp, err := l.svcCtx.Su.AddStudentToSubject(l.ctx, &suclient.AddStudentToSubjectReq{
		UserId:    req.UserId,
		SubjectId: req.SubejctId,
	})

	if err != nil {
		return &types.AddStudentInSubjectResp{}, err
	}

	return &types.AddStudentInSubjectResp{
		Code:    0,
		Message: "success",
		Data: types.AddStudentInSubjectRespData{
			Result: resp.Result,
		},
	}, nil
}
