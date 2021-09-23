package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"
	"go-zero-admin/rpc/us/usclient"
)

type PersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) PersionInfoLogic {
	return PersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersionInfoLogic) PersionInfo(req types.PersionInfoReq) (*types.PersionInfoResp, error) {
	// todo: add your logic here and delete this line

	value, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	//logx.Infof("value: %v",value)
	//logx.Info("value type:" +reflect.TypeOf(value).String())


	resp, err := l.svcCtx.Us.PersionInfo(l.ctx, &usclient.PersionInfoReq{
		Id: value,
	})

	if err != nil {
		return nil, err
	}

	return &types.PersionInfoResp{
		Code:    0,
		Message: "success",
		Data: types.PersionInfo{
			Id:          resp.Info.Id,
			Name:        resp.Info.Name,
			PhoneNumber: resp.Info.PhoneNumber,
			Email:       resp.Info.Email,
			RoleName:    resp.Info.RoleName,
			RoleId:      resp.Info.RoleId,
			RoleTypeId:  resp.Info.RoleTypeId,
			Gender:      resp.Info.Gender,
			ClassName:   resp.Info.ClassName,
			Academy:     resp.Info.Academy,
			School:      resp.Info.School,
			Grade:       resp.Info.Grade,
		},
	}, nil

}
