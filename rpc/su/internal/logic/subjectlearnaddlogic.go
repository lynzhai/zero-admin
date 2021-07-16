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

type SubjectLearnAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectLearnAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectLearnAddLogic {
	return &SubjectLearnAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectLearnAddLogic) SubjectLearnAdd(in *su.SubjectLearnAddReq) (*su.SubjectLearnAddResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.SuSubjectLearnModel.Insert(sumodel.SuSubjectLearn{
		UserId:    in.UserId,
		SubjectId: in.SubjectId,
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return &su.SubjectLearnAddResp{}, err
	}

	lastId, err := resp.LastInsertId()
	if err != nil {
		return &su.SubjectLearnAddResp{}, err
	}

	suSubjectLearn, err := l.svcCtx.SuSubjectLearnModel.FindOne(lastId)
	if err != nil {
		return &su.SubjectLearnAddResp{}, err
	}

	createTimeStr := ""
	if suSubjectLearn.CreateTime.Valid {
		createTimeStr = suSubjectLearn.CreateTime.Time.String()
	}

	return &su.SubjectLearnAddResp{
		Result: true,
		SubjectLearnInfoData: &su.SubjectLearnInfo{
			Id:         suSubjectLearn.Id,
			UserId:     suSubjectLearn.UserId,
			SubjectId:  suSubjectLearn.SubjectId,
			CreateTime: createTimeStr,
		},
	}, nil
}
