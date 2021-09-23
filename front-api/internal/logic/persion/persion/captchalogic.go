package logic

import (
	"context"
	"go-zero-admin/rpc/us/usclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(req types.CaptchaReq) (*types.CaptchaResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Us.Captcha(l.ctx, &usclient.CaptchaReq{})
	if err != nil {
		logx.Errorf("err=", err)
		return nil, err
	}
	return &types.CaptchaResp{
		CaptchaId: resp.CaptchaId,
		PicPath:   resp.PicPath,
	}, nil
}
