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

type SubjectLearnUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectLearnUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectLearnUpdateLogic {
	return &SubjectLearnUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectLearnUpdateLogic) SubjectLearnUpdate(in *su.SubjectLearnUpdateReq) (*su.SubjectLearnUpdateResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.SuSubjectLearnModel.Update(sumodel.SuSubjectLearn{
		Id:        in.Id,
		UserId:    in.UserId,
		SubjectId: in.SubjectId,
		CreateTime: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	})
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return &su.SubjectLearnUpdateResp{}, err
	}
	suSubjectLearn, err := l.svcCtx.SuSubjectLearnModel.FindOne(in.Id)

	if err != nil {
		logx.Errorf("err:" + err.Error())
		return &su.SubjectLearnUpdateResp{}, err
	}

	return &su.SubjectLearnUpdateResp{
		Result: true,
		SubjectLearnInfoData: &su.SubjectLearnInfo{
			Id:         suSubjectLearn.Id,
			UserId:     suSubjectLearn.UserId,
			SubjectId:  suSubjectLearn.SubjectId,
			CreateTime: suSubjectLearn.CreateTime.Time.String(),
		},
	}, nil
}
