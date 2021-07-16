package logic

import (
	"context"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"
	"go-zero-admin/rpc/su/suclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteStudentInSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStudentInSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteStudentInSubjectLogic {
	return DeleteStudentInSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStudentInSubjectLogic) DeleteStudentInSubject(req types.DeleteStudentInSubjectReq) (*types.DeleteStudentInSubjectResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}

	resp, err := l.svcCtx.Su.DeleteStudentInSubject(l.ctx, &suclient.DeleteStudentInSubjectReq{
		UserId:    req.UserId,
		SubjectId: req.SubjectId,
	})
	if err != nil {
		logx.Error("err:" + err.Error())
		return &types.DeleteStudentInSubjectResp{}, err
	}
	return &types.DeleteStudentInSubjectResp{
		Code:    0,
		Message: "success",
		Data: types.DeleteStudentInSubjectRespData{
			Result: resp.Result,
		},
	}, nil
}
