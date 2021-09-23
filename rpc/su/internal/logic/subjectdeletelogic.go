package logic

import (
	"context"
	"errors"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectDeleteLogic {
	return &SubjectDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectDeleteLogic) SubjectDelete(in *su.SubjectDeleteReq) (*su.SubjectDeleteResp, error) {
	// todo: add your logic here and delete this line

	subject, err := l.svcCtx.SuSubjectModel.FindOne(in.Id)
	if err != nil {
		return &su.SubjectDeleteResp{
			Result: false,
		}, errors.New("科目不存在")

	}
	if err := DelSubjectId(subject.MainTeacherId, subject.Id, l.svcCtx.RedisConn); err != nil {
		logx.Errorf("DelSubjectId error,MainTeacherId:" + strconv.Itoa(int(subject.MainTeacherId)) + ",err:" + err.Error())
	}

	if err := DelSubjectCountCache(subject.MainTeacherId, l.svcCtx.RedisConn); err != nil {
		logx.Errorf("DelSubjectCountCache error,MainTeacherId:" + strconv.Itoa(int(subject.MainTeacherId)) + ",err:" + err.Error())
	}

	err = l.svcCtx.SuSubjectModel.DeleteSoft(in.Id)

	if err != nil {
		logx.Errorf("DeleteSoft err:" + err.Error())
		return &su.SubjectDeleteResp{
			Result: false,
		}, err

	}
	return &su.SubjectDeleteResp{
		Result: true,
	}, nil
}
