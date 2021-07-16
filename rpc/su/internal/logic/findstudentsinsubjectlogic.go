package logic

import (
	"context"
	"go-zero-admin/rpc/us/usclient"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindStudentsInSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindStudentsInSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindStudentsInSubjectLogic {
	return &FindStudentsInSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询某个Subject中的学生内容
func (l *FindStudentsInSubjectLogic) FindStudentsInSubject(in *su.FindStudentsInSubjectReq) (*su.FindStudentsInSubjectResp, error) {
	// todo: add your logic here and delete this line

	suSubjectLearns, err := l.svcCtx.SuSubjectLearnModel.FindOneSubjectAllUser(in.SubjectId, in.Offset, in.PageSize)
	if err != nil {
		logx.Errorf("err:"+ err.Error())
		return &su.FindStudentsInSubjectResp{}, errorFindStudentInSubject
	}
	count, _ := l.svcCtx.SuSubjectLearnModel.OneSubjectUserCount(in.SubjectId)

	persionInfos := make([]*su.PersionInfo, 0)
	for _, v := range *suSubjectLearns {
		userId := v.UserId
		persionInfoResp, err := l.svcCtx.Us.PersionInfo(l.ctx, &usclient.PersionInfoReq{
			Id: userId,
		})
		if err != nil {
			logx.Errorf("err:" + err.Error())
			continue
		}

		persionInfos = append(persionInfos, &su.PersionInfo{
			Id:          persionInfoResp.Info.Id,
			PhoneNumber: persionInfoResp.Info.PhoneNumber,
			UniqueId:    persionInfoResp.Info.UniqueId,
			Number:      persionInfoResp.Info.Number,
			Email:       persionInfoResp.Info.Email,
			Gender:      persionInfoResp.Info.Gender,
			Name:        persionInfoResp.Info.Name,
			RoleId:      persionInfoResp.Info.RoleId,
			RoleType:    persionInfoResp.Info.RoleName,
			State:       persionInfoResp.Info.State,
			CreateTime:  persionInfoResp.Info.CreateTime,
			UpdateTime:  persionInfoResp.Info.UpdateTime,
			Class:       persionInfoResp.Info.Class,
			Academy:     persionInfoResp.Info.Academy,
			School:      persionInfoResp.Info.School,
			Grade:       persionInfoResp.Info.Grade,
		})
	}

	return &su.FindStudentsInSubjectResp{
		Total: count,
		List:  persionInfos,
	}, nil
}
