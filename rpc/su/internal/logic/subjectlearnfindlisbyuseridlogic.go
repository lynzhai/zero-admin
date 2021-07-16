package logic

import (
	"context"
	"errors"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectLearnFindLisByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectLearnFindLisByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectLearnFindLisByUserIdLogic {
	return &SubjectLearnFindLisByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectLearnFindLisByUserIdLogic) SubjectLearnFindLisByUserId(in *su.SubjectLearnFindListByUserIdReq) (*su.SubjectLearnFindListByUserIdResp, error) {
	// todo: add your logic here and delete this line

	suSubjectLearns, err := l.svcCtx.SuSubjectLearnModel.FindOneUserAllSubject(in.UserId, in.Offset, in.PageSize)
	if err != nil {
		return &su.SubjectLearnFindListByUserIdResp{
			Total: 0,
			List:  nil,
		}, errors.New("查询失败")
	}
	count, err := l.svcCtx.SuSubjectLearnModel.OneUserSubjectCount(in.UserId)
	if err != nil {
		return &su.SubjectLearnFindListByUserIdResp{
			Total: 0,
			List:  nil,
		}, errors.New("查询总数失败")
	}

	subjectLearnInfos := make([]*su.SubjectLearnInfo, 0)

	for _, v := range *suSubjectLearns {
		info := su.SubjectLearnInfo{
			Id:         v.Id,
			UserId:     v.UserId,
			SubjectId:  v.SubjectId,
			CreateTime: v.CreateTime.Time.String(),
		}
		subjectLearnInfos = append(subjectLearnInfos, &info)
	}

	return &su.SubjectLearnFindListByUserIdResp{
		Total: count,
		List:  subjectLearnInfos,
	}, nil
}
