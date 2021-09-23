package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	logx.Logger
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		//Store:  NewRedisCaptchaStore(svcCtx.RedisConn),
	}
}

func (l *CaptchaLogic) Captcha(in *us.CaptchaReq) (*us.CaptchaResp, error) {
	// todo: add your logic here and delete this line
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(50, 100, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, l.svcCtx.CaptchaStore)

	resp := us.CaptchaResp{}
	if id, b64s, err := cp.Generate(); err != nil {
		return nil, err
	} else {
		resp.CaptchaId = id
		resp.PicPath = b64s
	}

	return &resp, nil
}

