package logic

import (
	"context"
	"errors"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectLearnFindLisBySubjectIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectLearnFindLisBySubjectIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectLearnFindLisBySubjectIdLogic {
	return &SubjectLearnFindLisBySubjectIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectLearnFindLisBySubjectIdLogic) SubjectLearnFindLisBySubjectId(in *su.SubjectLearnFindListBySubjectIdReq) (*su.SubjectLearnFindListBySubjectIdResp, error) {
	// todo: add your logic here and delete this line

	suSubjectLearns, err := l.svcCtx.SuSubjectLearnModel.FindOneSubjectAllUser(in.SubjectId, in.Offset, in.PageSize)
	if err != nil {
		return &su.SubjectLearnFindListBySubjectIdResp{
			Total: 0,
			List:  nil,
		}, errors.New("查询失败")
	}
	count, err := l.svcCtx.SuSubjectLearnModel.OneSubjectUserCount(in.SubjectId)
	if err != nil {
		return &su.SubjectLearnFindListBySubjectIdResp{
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

	return &su.SubjectLearnFindListBySubjectIdResp{
		Total: count,
		List:  subjectLearnInfos,
	}, nil

	return &su.SubjectLearnFindListBySubjectIdResp{}, nil
}
