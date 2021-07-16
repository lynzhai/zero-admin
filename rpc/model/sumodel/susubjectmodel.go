package sumodel

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"strings"
	"time"
)

var (
	suSubjectFieldNames          = builderx.RawFieldNames(&SuSubject{})
	suSubjectRows                = strings.Join(suSubjectFieldNames, ",")
	suSubjectRowsExpectAutoSet   = strings.Join(stringx.Remove(suSubjectFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	suSubjectRowsWithPlaceHolder = strings.Join(stringx.Remove(suSubjectFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	//加两个 ,？  ,？
	usSubjectRowsForInsert = strings.Join(stringx.Remove(suSubjectFieldNames, "`id`"), ",")

	// 加一个 ,?
	usSubjectRowsForUpdate     = strings.Join(stringx.Remove(suSubjectFieldNames, "`id`", "`create_time`"), "=?,") + "=?"
	usSubjectRowsForDeleteSoft = strings.Join([]string{"delete_time"}, "=?,") + "=?"

	cacheSuSubjectIdPrefix   = "cache#suSubject#id#"
	cacheSuSubjectCodePrefix = "cache#suSubject#code#"
)

type (
	SuSubjectModel interface {
		Insert(data SuSubject) (sql.Result, error)
		FindOne(id int64) (*SuSubject, error)
		FindOneByCode(code string) (*SuSubject, error)
		FindOneUserAll(userId int64, Current int64, PageSize int64) (*[]SuSubject, error)
		Count(userId int64, ) (int64, error)
		Update(data SuSubject) error
		Delete(id int64) error
		DeleteSoft(id int64) error
		DeleteIdCache(id int64) error
		DeleteCodeCache(code string) error
	}

	defaultSuSubjectModel struct {
		sqlc.CachedConn
		table string
	}

	SuSubject struct {
		Id                 int64          `db:"id"`
		Uuid               string         `db:"uuid"`                 // uuid
		Name               sql.NullString `db:"name"`                 // 课程名称
		Status             sql.NullInt64  `db:"status"`               // 课程状态
		Code               string         `db:"code"`                 // 课程码
		MaxPersion         sql.NullInt64  `db:"max_persion"`          // 最大人数
		MainTeacherId      int64          `db:"main_teacher_id"`      // 主讲老师id
		AssistantTeacherId sql.NullInt64  `db:"assistant_teacher_id"` // 助教老师id
		Introduce          sql.NullString `db:"introduce"`            // 介绍
		Backup             sql.NullString `db:"backup"`               // 备注
		CreateTime         sql.NullTime   `db:"create_time"`
		UpdateTime         sql.NullTime   `db:"update_time"`
		DeleteTime         sql.NullTime   `db:"delete_time"`
	}
)

func NewSuSubjectModel(conn sqlx.SqlConn, c cache.CacheConf) SuSubjectModel {
	return &defaultSuSubjectModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`su_subject`",
	}
}

func (m *defaultSuSubjectModel) Insert(data SuSubject) (sql.Result, error) {
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, data.Code)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usSubjectRowsForInsert)
		return conn.Exec(query, data.Uuid, data.Name, data.Status, data.Code, data.MaxPersion, data.MainTeacherId, data.AssistantTeacherId, data.Introduce, data.Backup, data.CreateTime, data.UpdateTime, data.DeleteTime)
	}, suSubjectCodeKey)
	return ret, err
}

func (m *defaultSuSubjectModel) FindOne(id int64) (*SuSubject, error) {
	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, id)
	var resp SuSubject
	err := m.QueryRow(&resp, suSubjectIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", suSubjectRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSuSubjectModel) FindOneByCode(code string) (*SuSubject, error) {
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, code)
	var resp SuSubject
	err := m.QueryRowIndex(&resp, suSubjectCodeKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", suSubjectRows, m.table)
		if err := conn.QueryRow(&resp, query, code); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 倒叙
func (m *defaultSuSubjectModel) FindOneUserAll(userId int64, Current int64, PageSize int64) (*[]SuSubject, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s where `main_teacher_id` = ? and `delete_time` is null order by id desc limit ?, ? ", suSubjectRows, m.table)
	var resp []SuSubject
	err := m.QueryRowsNoCache(&resp, query, userId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultSuSubjectModel) Count(userId int64, ) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `main_teacher_id` = ? and `delete_time` is null", m.table)

	var count int64
	err := m.QueryRowNoCache(&count, query, userId)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultSuSubjectModel) Update(data SuSubject) error {
	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usSubjectRowsForUpdate)
		data.UpdateTime.Time = time.Now()
		data.UpdateTime.Valid = true
		return conn.Exec(query, data.Uuid, data.Name, data.Status, data.Code, data.MaxPersion, data.MainTeacherId, data.AssistantTeacherId, data.Introduce, data.Backup, data.UpdateTime, data.DeleteTime, data.Id)
	}, suSubjectIdKey)
	return err
}

func (m *defaultSuSubjectModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, id)
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, data.Code)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, suSubjectIdKey, suSubjectCodeKey)
	return err
}

func (m *defaultSuSubjectModel) DeleteSoft(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, id)
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, data.Code)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		//query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usSubjectRowsForDeleteSoft)
		currTime := time.Now()
		return conn.Exec(query, currTime, id)
	}, suSubjectIdKey, suSubjectCodeKey)
	return err
}

func (m *defaultSuSubjectModel) DeleteIdCache(id int64) error {
	usSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, id)
	return m.DelCache(usSubjectIdKey)
}

func (m *defaultSuSubjectModel) DeleteCodeCache(code string) error {
	usSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, code)
	return m.DelCache(usSubjectCodeKey)
}

func (m *defaultSuSubjectModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, primary)
}

func (m *defaultSuSubjectModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", suSubjectRows, m.table)
	return conn.QueryRow(v, query, primary)
}
