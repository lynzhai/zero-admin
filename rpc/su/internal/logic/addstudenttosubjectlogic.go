package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/sumodel"
	"time"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddStudentToSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStudentToSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentToSubjectLogic {
	return &AddStudentToSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  添加学生到课程
func (l *AddStudentToSubjectLogic) AddStudentToSubject(in *su.AddStudentToSubjectReq) (*su.AddStudentToSubjectResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.SuSubjectLearnModel.Insert(sumodel.SuSubjectLearn{
		UserId:    in.UserId,
		SubjectId: in.SubjectId,
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		logx.Errorf("err:"+ err.Error())
		return &su.AddStudentToSubjectResp{}, errorAddStudentToSubject
	}
	return &su.AddStudentToSubjectResp{
		Result: true,
	}, nil
}
