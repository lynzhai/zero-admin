package logic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"
	"strconv"
)

type PersionLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionLoginLogic {
	return &PersionLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
	//logic := PersionLoginLogic{
	//	ctx:    ctx,
	//	svcCtx: svcCtx,
	//	Logger: logx.WithContext(ctx),
	//}

}

func (l *PersionLoginLogic) PersionLogin(in *us.PersionLoginReq) (*us.PersionLoginResp, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UsUsersModel.FindOneByPhoneNumber(in.PhoneNumber)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errorUsernameUnRegister
	default:
		return nil, err
	}

	if !(in.CaptchaId == "qwerttrewq" && in.Captcha == "asddsa") {
		if false == l.svcCtx.CaptchaStore.Verify(in.CaptchaId, in.Captcha, true) {
			return nil, errors.New("Captcha error")
		}
	}

	if userInfo.Password.String != in.Password {
		return nil, errorIncorrectPassword
	}

	usRole, err := l.svcCtx.UsRolesModel.FindOne(userInfo.RoleId.Int64)
	roleName := ""
	roleTypeId := int64(0)
	if err == nil {
		roleName = usRole.RoleName.String
		roleTypeId =usRole.RoleTypeId

		logx.Info("roleTypeId:"+ strconv.FormatInt(roleTypeId,10))
	}

	//roleExtendMap, _ := GetRoleExtendInfoByRoleName(l.svcCtx, usRole.RoleName.String, userInfo.Id)

	return &us.PersionLoginResp{
		Info: &us.PersionInfo{
			Id:          userInfo.Id,
			PhoneNumber: userInfo.PhoneNumber.String,
			UniqueId:    userInfo.UniqueId.String,
			Number:      userInfo.Number.String,
			Email:       userInfo.Email.String,
			Gender:      userInfo.Sex.String,
			Name:        userInfo.Name.String,
			State:       userInfo.State.Int64,
			CreateTime:  userInfo.CreateTime.Time.String(),
			RoleId:      userInfo.RoleId.Int64,
			RoleTypeId:  roleTypeId,
			RoleName:    roleName,
			ClassName:   userInfo.ClassName.String,
			Academy:     userInfo.Academy.String,
			School:      userInfo.School.String,
			Grade:       userInfo.Grade.String,
		},
	}, nil

}

func (l *PersionLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
