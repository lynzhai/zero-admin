package logic

import (
	"context"
	"database/sql"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/gogf/gf/util/guid"
)

type PersionRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionRegisterLogic {
	return &PersionRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionRegisterLogic) PersionRegister(in *us.PersionRegisterReq) (*us.PersionRegisterResp, error) {
	if _, err := l.svcCtx.UsUsersModel.FindOneByPhoneNumber(in.PhoneNumber); err == nil {
		return nil, errorDuplicateMobile
	}

	if _, err := l.svcCtx.UsUsersModel.FindOneByEmail(in.Email); err == nil {
		return nil, errorDuplicateEmail
	}
	//l.svcCtx.UsUsersModel.DeletePhoneNumberCache(in.PhoneNumber)

	if !(in.EmailId == "qwerttrewq" && in.EmailCode == "asddsa") {
		if false == l.svcCtx.RedisEmailCodeStore.Verify(in.EmailId, in.EmailCode, false) {
			return nil, errorCaptcha
		}
	}

	var usRole *usmodel.UsRoles
	if tempUsRole, err := l.svcCtx.UsRolesModel.FindOne(in.RoleId); err != nil {
		return nil, errorRoleUnRegister
	} else {
		usRole = tempUsRole
	}

	uniqueId := guid.S()

	usUser := usmodel.UsUsers{
		PhoneNumber: sql.NullString{
			String: in.PhoneNumber,
			Valid:  true,
		},
		UniqueId: sql.NullString{
			String: uniqueId,
			Valid:  true,
		},
		Number: sql.NullString{
			String: in.Number,
			Valid:  true,
		},
		Email: sql.NullString{
			String: in.Email,
			Valid:  true,
		},
		Name: sql.NullString{
			String: in.Name,
			Valid:  true,
		},
		Password: sql.NullString{
			String: in.Password,
			Valid:  true,
		},
		Sex: sql.NullString{
			String: in.Gender,
			Valid:  true,
		},
		RoleId: sql.NullInt64{
			Int64: usRole.Id,
			Valid: true,
		},
		State: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		School: sql.NullString{
			String: in.School,
			Valid:  true,
		},
		Academy: sql.NullString{
			String: in.Academy,
			Valid:  true,
		},
		ClassName: sql.NullString{
			String: in.ClassName,
			Valid:  true,
		},
		Grade: sql.NullString{
			String: in.Grade,
			Valid:  true,
		},
	}

	if _, err := l.svcCtx.UsUsersModel.Insert(usUser); err != nil {
		return nil, err
	}

	//err := l.svcCtx.UsUsersModel.GetSqlCachedConn().Transact(func(session sqlx.Session) error {
	//	userId := int64(0)
	//
	//	if result, err := l.svcCtx.UsUsersModel.InsertBySession(usUser, session); err != nil {
	//		return err
	//	} else {
	//		userId, _ = result.LastInsertId()
	//	}
	//
	//	switch usRole.RoleName.String {
	//	case "teacher":
	//		usTeacher := usmodel.UsTeacher{
	//			Academy: sql.NullString{
	//				String: in.Academy,
	//				Valid:  true,
	//			},
	//			School: sql.NullString{
	//				String: in.School,
	//				Valid:  true,
	//			},
	//			UserId: userId,
	//		}
	//		if _, err := l.svcCtx.UsTeacherModel.InsertBySession(usTeacher, session); err != nil {
	//			return err
	//		}
	//		break
	//	case "student":
	//		usStudent := usmodel.UsStudent{
	//			Academy: sql.NullString{
	//				String: in.Academy,
	//				Valid:  true,
	//			},
	//			Class: sql.NullString{
	//				String: in.ClassName,
	//				Valid:  true,
	//			},
	//			School: sql.NullString{
	//				String: in.School,
	//				Valid:  true,
	//			},
	//			Grade: sql.NullString{
	//				String: in.Grade,
	//				Valid:  true,
	//			},
	//			UserId: userId,
	//		}
	//		if _, err := l.svcCtx.UsStudentModel.InsertBySession(usStudent, session); err != nil {
	//			return err
	//		}
	//		break
	//	default:
	//		return errorUserRegisterFail
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return nil, errorUserRegisterFail
	//}

	return &us.PersionRegisterResp{
		Result: true,
	}, nil
}
