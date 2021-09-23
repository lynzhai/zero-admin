package logic

import (
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"go-zero-admin/common"
	"go-zero-admin/rpc/model/sumodel"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/su/internal/svc"
	"strconv"
	"time"
)

//type Content struct {
//	Id         string    `json:"id"`
//	Title      string    `json:"title"`
//	Content    string    `json:"content"`
//	CreateTime time.Time `json:"create_time"`
//}
const maxRedisSubjectNumber = 20

const bizSubjectCacheKey = `bizcache#subject#`

const bizSubjectTotalNumberCacheKey = `bizcache#subject#total#number#`

func SubjectTotalNumberCacheExists(userId int64, r *redis.Redis) (bool, error) {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectTotalNumberCacheKey, userId)
	return r.Exists(cacheKey)
}

func SetSubjectTotalNumberCache(userId int64, count int64, r *redis.Redis) error {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectTotalNumberCacheKey, userId)
	countStr := strconv.FormatInt(count, 10)

	exists, _ := SubjectTotalNumberCacheExists(userId, r)
	if err := r.Set(cacheKey, countStr); err != nil {
		return err
	}
	if !exists {
		SetSubjectTotalNumberCacheExpire(userId, common.GetCommonRedisExpireSeconds(), r)
	}
	return nil
}
func GetSubjectTotalNumberCache(userId int64, r *redis.Redis) (int64, error) {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectTotalNumberCacheKey, userId)
	valueStr, err := r.Get(cacheKey)
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, err
}

func DelSubjectCountCache(userId int64, r *redis.Redis) error {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectTotalNumberCacheKey, userId)
	_, err := r.Del(cacheKey)
	return err
}

func SetSubjectTotalNumberCacheExpire(userId int64, seconds int, r *redis.Redis) error {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectTotalNumberCacheKey, userId)
	return r.Expire(cacheKey, seconds)
}

func SubjectRedisCacheExists(userId int64, r *redis.Redis) (bool, error) {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	return r.Exists(cacheKey)
}

func SetSubjectRedisCacheExpire(userId int64, seconds int, r *redis.Redis) error {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	return r.Expire(cacheKey, seconds)
}

func GetSubjectRedisCacheLen(userId int64, r *redis.Redis) (val int, err error) {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	return r.Zcard(cacheKey)
}

func AddSubjectId(userId int64, subjectId int64, createTime time.Time, r *redis.Redis) error {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	//save subject id
	subjectIdStr := strconv.FormatInt(subjectId, 10)

	exists, _ := SubjectRedisCacheExists(userId, r)

	if _, err := r.Zadd(cacheKey, createTime.UnixNano()/1e6, subjectIdStr); err != nil {
		return err
	}
	// 获取key的数量
	val, _ := r.Zcard(cacheKey)

	// 超过指定数目，则删除旧的
	if val > maxRedisSubjectNumber {
		pageSize := val - maxRedisSubjectNumber
		// 获取从小到大的一页10个排列
		nowTime := time.Now()
		twoYearsAgoTime := nowTime.AddDate(-2, 0, 0)
		//获取最老的时间排序
		pairs, err := r.ZrangebyscoreWithScoresAndLimit(cacheKey, twoYearsAgoTime.UnixNano()/1e6, nowTime.UnixNano()/1e6, 1, pageSize)
		needDelKeys := make([]string, 0)
		if err == nil {
			for _, v := range pairs {
				needDelKeys = append(needDelKeys, v.Key)
			}
		}
		r.Zrem(cacheKey, needDelKeys)
	}

	if !exists {
		SetSubjectBizCacheExpire(userId, common.GetCommonRedisExpireSeconds(), r)
	}
	return nil
}

func DelSubjectId(userId int64, subjectId int64, r *redis.Redis) error {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	s := strconv.FormatInt(subjectId, 10)
	_, err := r.Zrem(cacheKey, s)
	return err
}

func SetSubjectBizCacheExpire(userId int64, seconds int, r *redis.Redis) {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	r.Expire(cacheKey, seconds)
}

func compress(c *usmodel.UsRoles) string {
	return strconv.Itoa(int(c.Id))
}

func uncompress(svcCtx *svc.ServiceContext, v string) (*sumodel.SuSubject, error) {
	// todo: do it yourself
	subjectId, err := strconv.Atoi(v)
	if err != nil {
		return nil, errors.New("转换失败")
	}
	suSubject, err := svcCtx.SuSubjectModel.FindOne(int64(subjectId))
	if err != nil {
		return nil, errors.New("查找subject 缓存失败")
	}
	return suSubject, nil
}

// ListByRangeTime提供根据时间段进行数据查询
func ListSubjectByRangeTime(userId int64, startTime, stopTime time.Time, offset int64, pageSize int64, r *redis.Redis, svcCtx *svc.ServiceContext) ([]*sumodel.SuSubject, error) {
	cacheKey := fmt.Sprintf("%s%v", bizSubjectCacheKey, userId)
	kvs, err := r.ZrevrangebyscoreWithScoresAndLimit(cacheKey, startTime.UnixNano()/1e6, stopTime.UnixNano()/1e6, int(offset-1), int(pageSize))
	if err != nil {
		return nil, err
	}
	//logx.Info("len(kvs):"+ strconv.Itoa(len(kvs)))

	var list []*sumodel.SuSubject
	for _, kv := range kvs {
		data, err := uncompress(svcCtx, kv.Key)
		if err == nil {
			list = append(list, data)
		}
	}
	return list, nil
}
