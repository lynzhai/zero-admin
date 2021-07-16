package sumodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	suSubjectLearnFieldNames          = builderx.RawFieldNames(&SuSubjectLearn{})
	suSubjectLearnRows                = strings.Join(suSubjectLearnFieldNames, ",")
	suSubjectLearnRowsExpectAutoSet   = strings.Join(stringx.Remove(suSubjectLearnFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	suSubjectLearnRowsWithPlaceHolder = strings.Join(stringx.Remove(suSubjectLearnFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	//加 1个 ？
	usSubjectLearnRowsForInsert = strings.Join(stringx.Remove(suSubjectLearnFieldNames, "`id`"), ",")

	usSubjectLearnRowsForUpdate     = strings.Join(stringx.Remove(suSubjectLearnFieldNames, "`id`", "`create_time`"), "=?,") + "=?"
	usSubjectLearnRowsForDeleteSoft = strings.Join([]string{"delete_time"}, "=?,") + "=?"

	cacheSuSubjectLearnIdPrefix = "cache#suSubjectLearn#id#"
)

type (
	SuSubjectLearnModel interface {
		Insert(data SuSubjectLearn) (sql.Result, error)
		FindOne(id int64) (*SuSubjectLearn, error)
		FindOneUserAllSubject(userId int64, Current int64, PageSize int64) (*[]SuSubjectLearn, error)
		FindOneSubjectAllUser(subjectId int64, Current int64, PageSize int64) (*[]SuSubjectLearn, error)
		OneUserSubjectCount(userId int64, ) (int64, error)
		OneSubjectUserCount(subjectId int64, ) (int64, error)
		Update(data SuSubjectLearn) error
		Delete(id int64) error
		DeleteBySubjectAndUserId(UserId, SubjectId int64) error
	}

	defaultSuSubjectLearnModel struct {
		sqlc.CachedConn
		table string
	}

	SuSubjectLearn struct {
		Id         int64        `db:"id"`
		UserId     int64        `db:"user_id"`
		SubjectId  int64        `db:"subject_id"`
		CreateTime sql.NullTime `db:"create_time"`
	}
)

func NewSuSubjectLearnModel(conn sqlx.SqlConn, c cache.CacheConf) SuSubjectLearnModel {
	return &defaultSuSubjectLearnModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`su_subject_learn`",
	}
}

func (m *defaultSuSubjectLearnModel) Insert(data SuSubjectLearn) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, usSubjectLearnRowsForInsert)
	ret, err := m.ExecNoCache(query, data.UserId, data.SubjectId, data.CreateTime)

	return ret, err
}

func (m *defaultSuSubjectLearnModel) FindOne(id int64) (*SuSubjectLearn, error) {
	suSubjectLearnIdKey := fmt.Sprintf("%s%v", cacheSuSubjectLearnIdPrefix, id)
	var resp SuSubjectLearn
	err := m.QueryRow(&resp, suSubjectLearnIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", suSubjectLearnRows, m.table)
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

func (m *defaultSuSubjectLearnModel) FindOneUserAllSubject(userId int64, Current int64, PageSize int64) (*[]SuSubjectLearn, error) {
	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s where `user_id` = ?  limit ?, ? ", suSubjectLearnRows, m.table)
	var resp []SuSubjectLearn
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
func (m *defaultSuSubjectLearnModel) FindOneSubjectAllUser(subjectId int64, Current int64, PageSize int64) (*[]SuSubjectLearn, error) {
	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s where `subject_id` = ?  limit ?, ? ", suSubjectLearnRows, m.table)
	var resp []SuSubjectLearn
	err := m.QueryRowsNoCache(&resp, query, subjectId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSuSubjectLearnModel) OneUserSubjectCount(userId int64, ) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `user_id` = ?", m.table)

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
func (m *defaultSuSubjectLearnModel) OneSubjectUserCount(subjectId int64, ) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `subject_id` = ?", m.table)

	var count int64
	err := m.QueryRowNoCache(&count, query, subjectId)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultSuSubjectLearnModel) Update(data SuSubjectLearn) error {
	suSubjectLearnIdKey := fmt.Sprintf("%s%v", cacheSuSubjectLearnIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usSubjectLearnRowsForUpdate)
		return conn.Exec(query, data.UserId, data.SubjectId, data.Id)
	}, suSubjectLearnIdKey)
	return err
}

func (m *defaultSuSubjectLearnModel) Delete(id int64) error {

	suSubjectLearnIdKey := fmt.Sprintf("%s%v", cacheSuSubjectLearnIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, suSubjectLearnIdKey)
	return err
}

func (m *defaultSuSubjectLearnModel) DeleteBySubjectAndUserId(userId, subjectId int64) error {
	query := fmt.Sprintf("select %s from %s where `subject_id` = ?  and `user_id` = ? limit 1", suSubjectLearnRows, m.table)
	var resp SuSubjectLearn
	err := m.QueryRowNoCache(&resp, query, subjectId, userId)

	switch err {
	case nil:
		return m.Delete(resp.Id)
	case sqlc.ErrNotFound:
		return nil
	default:
		return err
	}

}

func (m *defaultSuSubjectLearnModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSuSubjectLearnIdPrefix, primary)
}

func (m *defaultSuSubjectLearnModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", suSubjectLearnRows, m.table)
	return conn.QueryRow(v, query, primary)
}
