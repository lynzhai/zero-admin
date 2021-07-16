package logic

import (
	"context"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindSubjectByStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSubjectByStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSubjectByStudentLogic {
	return &FindSubjectByStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询课程中的 学生信息
func (l *FindSubjectByStudentLogic) FindSubjectByStudent(in *su.FindSubjectByStudentReq) (*su.FindSubjectByStudentResp, error) {
	// todo: add your logic here and delete this line

	suSubjectLearns, err := l.svcCtx.SuSubjectLearnModel.FindOneUserAllSubject(in.UserId, in.Offset, in.PageSize)
	if err != nil {
		return &su.FindSubjectByStudentResp{}, errorFindSubjectByStudent
	}
	count, err := l.svcCtx.SuSubjectLearnModel.OneUserSubjectCount(in.UserId)

	list := make([]*su.SubjectInfo, 0)
	for _, v := range *suSubjectLearns {
		subjectId := v.SubjectId
		suSubject, err := l.svcCtx.SuSubjectModel.FindOne(subjectId)
		if err != nil {
			logx.Error("err:" + err.Error())
			continue
		}
		list = append(list, &su.SubjectInfo{
			Id:                 suSubject.Id,
			Uuid:               suSubject.Uuid,
			Name:               suSubject.Name.String,
			Status:             suSubject.Status.Int64,
			Code:               suSubject.Code,
			MaxPersion:         suSubject.MaxPersion.Int64,
			MainTeacherId:      suSubject.MainTeacherId,
			AssistantTeacherId: suSubject.AssistantTeacherId.Int64,
			Introduce:          suSubject.Introduce.String,
			Backup:             suSubject.Backup.String,
			CreateTime:         suSubject.CreateTime.Time.String(),
			UpdateTime:         suSubject.UpdateTime.Time.String(),
		})

	}
	return &su.FindSubjectByStudentResp{
		Total: count,
		List:  list,
	}, nil
}
