package logic

import (
	"context"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/guid"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeLogic {
	return &EmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailCodeLogic) EmailCode(in *us.EmailCodeReq) (*us.EmailCodeResp, error) {
	// todo: add your logic here and delete this line

	uniqueId := guid.S()

	num := grand.N(100000,999999)
	emailCode := strconv.Itoa(num)

	logx.Info("email:" +in.Email)
	logx.Info("emailCode:" +emailCode)

	err := l.svcCtx.AliEmail.Send(in.Email, "注册验证码", emailCode)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.RedisEmailCodeStore.Set(uniqueId, emailCode)

	if err != nil {
		return nil, err
	}

	return &us.EmailCodeResp{
		EmailId: uniqueId,
	}, nil
}
