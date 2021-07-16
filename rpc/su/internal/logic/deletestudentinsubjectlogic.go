package logic

import (
	"context"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteStudentInSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStudentInSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentInSubjectLogic {
	return &DeleteStudentInSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  从课程中删除学生
func (l *DeleteStudentInSubjectLogic) DeleteStudentInSubject(in *su.DeleteStudentInSubjectReq) (*su.DeleteStudentInSubjectResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.SuSubjectLearnModel.DeleteBySubjectAndUserId(in.UserId, in.SubjectId)
	if err != nil {
		logx.Errorf("err" +err.Error())
		return &su.DeleteStudentInSubjectResp{}, errorDeleteStudentInSubject
	}

	return &su.DeleteStudentInSubjectResp{
		Result: true,
	}, nil
}
