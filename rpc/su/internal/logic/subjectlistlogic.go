package logic

import (
	"context"
	"errors"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"
	"strconv"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func dbFindSubjectList(userId, offset, pageSize, totalCount int64, svcCtx *svc.ServiceContext) (*su.MainTeacherSubjectListReap, error) {
	suSubjects, err := svcCtx.SuSubjectModel.FindOneUserAll(userId, offset, pageSize)

	if err != nil {
		return &su.MainTeacherSubjectListReap{
			Total: 0,
			List:  nil,
		}, errors.New("查询失败")
	}

	subjectInfos := make([]*su.SubjectInfo, 0)
	for _, v := range *suSubjects {
		info := su.SubjectInfo{
			Id:                 v.Id,
			Uuid:               v.Uuid,
			Name:               v.Name.String,
			Status:             v.Status.Int64,
			Code:               v.Code,
			MaxPersion:         v.MaxPersion.Int64,
			MainTeacherId:      v.MainTeacherId,
			AssistantTeacherId: v.AssistantTeacherId.Int64,
			Introduce:          v.Introduce.String,
			Backup:             v.Backup.String,
			CreateTime:         v.CreateTime.Time.String(),
			UpdateTime:         v.UpdateTime.Time.String(),
		}
		subjectInfos = append(subjectInfos, &info)
	}

	return &su.MainTeacherSubjectListReap{
		Total: totalCount,
		List:  subjectInfos,
	}, nil
}

func bizCacheFindSubjectList(userId, offset, pageSize, totalCount int64, svcCtx *svc.ServiceContext) (*su.MainTeacherSubjectListReap, error) {
	endTime := time.Now()
	startTime := endTime.AddDate(-2, 0, 0)
	subjects, err := ListSubjectByRangeTime(userId, startTime, endTime, offset, pageSize, svcCtx.RedisConn, svcCtx)
	if err != nil {
		return &su.MainTeacherSubjectListReap{
			Total: 0,
			List:  nil,
		}, errors.New("查询缓存失败")
	}

	subjectInfos := make([]*su.SubjectInfo, 0)

	for _, v := range subjects {
		info := su.SubjectInfo{
			Id:                 v.Id,
			Uuid:               v.Uuid,
			Name:               v.Name.String,
			Status:             v.Status.Int64,
			Code:               v.Code,
			MaxPersion:         v.MaxPersion.Int64,
			MainTeacherId:      v.MainTeacherId,
			AssistantTeacherId: v.AssistantTeacherId.Int64,
			Introduce:          v.Introduce.String,
			Backup:             v.Backup.String,
			CreateTime:         v.CreateTime.Time.String(),
			UpdateTime:         v.UpdateTime.Time.String(),
		}
		subjectInfos = append(subjectInfos, &info)
	}
	return &su.MainTeacherSubjectListReap{
		Total: totalCount,
		List:  subjectInfos,
	}, nil
}

func (l *SubjectListLogic) SubjectList(in *su.MainTeacherSubjectListReq) (*su.MainTeacherSubjectListReap, error) {
	// todo: add your logic here and delete this line

	// 缓存最新的maxRedisSubjectNumber个subject
	realCacheLen := 0
	subjectListCacheExists, _ := SubjectRedisCacheExists(in.Id, l.svcCtx.RedisConn)
	if !subjectListCacheExists {
		suSubjects, err := l.svcCtx.SuSubjectModel.FindOneUserAll(in.Id, 0, maxRedisSubjectNumber)
		if err != nil {
			return &su.MainTeacherSubjectListReap{
				Total: 0,
				List:  nil,
			}, errors.New("查询失败")
		}

		for _, v := range *suSubjects {
			// 添加subject 缓存
			err := AddSubjectId(v.MainTeacherId, v.Id, v.CreateTime.Time, l.svcCtx.RedisConn)
			if err != nil {
				logx.Errorf("err:" + err.Error())
			}
		}

		if len(*suSubjects) <= maxRedisSubjectNumber {
			realCacheLen = len(*suSubjects)
		} else {
			realCacheLen = maxRedisSubjectNumber
		}
		subjectListCacheExists, _ = SubjectRedisCacheExists(in.Id, l.svcCtx.RedisConn)
	} else {
		if value, err := GetSubjectRedisCacheLen(in.Id, l.svcCtx.RedisConn); err == nil {
			logx.Info("value:" + strconv.Itoa(value))
			realCacheLen = value
		}
	}

	subjectTotalCount := int64(0)
	subjectCountCacheExists, _ := SubjectTotalNumberCacheExists(in.Id, l.svcCtx.RedisConn)
	if !subjectCountCacheExists {
		count, err := l.svcCtx.SuSubjectModel.Count(in.Id)
		if err == nil {
			subjectTotalCount = count
			// 设置最大条目数 缓存
			err := SetSubjectTotalNumberCache(in.Id, count, l.svcCtx.RedisConn)
			if err != nil {
				logx.Errorf("err:" + err.Error())
			}
		}
		subjectCountCacheExists, _ = SubjectTotalNumberCacheExists(in.Id, l.svcCtx.RedisConn)
	} else {
		count, err := GetSubjectTotalNumberCache(in.Id, l.svcCtx.RedisConn)
		if err == nil {
			subjectTotalCount = count
		}
	}

	// 总条目数或者 要求页码数大于总条目数
	if subjectTotalCount == 0 || (in.Offset-1)*in.PageSize > subjectTotalCount {
		return &su.MainTeacherSubjectListReap{
			Total: 0,
			List:  nil,
		}, nil
	}

	logx.Info("subjectTotalCount:" + strconv.Itoa(int(subjectTotalCount)))
	logx.Info("realCacheLen:" + strconv.Itoa(int(realCacheLen)))
	logx.Info("subjectListCacheExists:" + strconv.FormatBool(subjectListCacheExists))

	if subjectTotalCount == int64(realCacheLen) && subjectListCacheExists { //从缓存找
		return bizCacheFindSubjectList(in.Id, in.Offset, in.PageSize, subjectTotalCount, l.svcCtx)
	} else if (in.Offset*in.PageSize) > int64(realCacheLen) || !subjectListCacheExists { //从数据库找
		return dbFindSubjectList(in.Id, in.Offset, in.PageSize, subjectTotalCount, l.svcCtx)
	} else {
		return bizCacheFindSubjectList(in.Id, in.Offset, in.PageSize, subjectTotalCount, l.svcCtx)
	}
}
