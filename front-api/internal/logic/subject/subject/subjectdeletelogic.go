package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubjectDeleteLogic {
	return SubjectDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectDeleteLogic) SubjectDelete(req types.SubjectDeleteReq) (*types.SubjectDeleteResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Errorf("err:" + err.Error())
		return nil, err
	}

	resp, err := l.svcCtx.Su.SubjectDelete(l.ctx, &suclient.SubjectDeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return &types.SubjectDeleteResp{}, err
	}

	return &types.SubjectDeleteResp{
		Code:    0,
		Message: "success",
		Data: types.SubjectDeleteRespData{
			Result: resp.Result,
		},
	}, nil
}
