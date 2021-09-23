package logic

import (
	"context"
	"go-zero-admin/rpc/us/usclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) EmailCodeLogic {
	return EmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeLogic) EmailCode(req types.EmailCodeReq) (*types.EmailCodeResp, error) {
	// todo: add your logic here and delete this line

	emailCodeResp, err := l.svcCtx.Us.EmailCode(l.ctx, &usclient.EmailCodeReq{
		Email: req.Email,
	})
	if err != nil {
		logx.Errorf("err=", err)
		return nil, err
	}


	return &types.EmailCodeResp{
		EmailId: emailCodeResp.EmailId,
	}, nil
}
