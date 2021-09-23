package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/usmodel"
	"time"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionAddLogic {
	return &PersionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionAddLogic) PersionAdd(in *us.PersionAddReq) (*us.PersionAddResp, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.UsUsersModel.Insert(usmodel.UsUsers{
		PhoneNumber: sql.NullString{
			String: in.Data.PhoneNumber,
			Valid:  true,
		},
		UniqueId: sql.NullString{
			String: in.Data.UniqueId,
			Valid:  true,
		},
		Number: sql.NullString{
			String: in.Data.Number,
			Valid:  true,
		},
		Email: sql.NullString{
			String: in.Data.Email,
			Valid:  true,
		},
		Name: sql.NullString{
			String: in.Data.Name,
			Valid:  true,
		},
		Password: sql.NullString{
			String: in.Data.Password,
			Valid:  true,
		},
		Sex: sql.NullString{
			String: in.Data.Gender,
			Valid:  true,
		},
		RoleId: sql.NullInt64{
			Int64: in.Data.RoleId,
			Valid: true,
		},
		State: sql.NullInt64{
			Int64: in.Data.State,
			Valid: true,
		},
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		DeleteTime: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	})
	if err != nil {
		return &us.PersionAddResp{
			Result: false,
		}, err
	}
	return &us.PersionAddResp{
		Result: true,
	}, nil
}
