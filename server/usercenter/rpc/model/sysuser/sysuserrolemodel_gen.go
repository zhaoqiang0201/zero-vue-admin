// Code generated by goctl. DO NOT EDIT!

package sysuser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserRoleFieldNames          = builder.RawFieldNames(&SysUserRole{})
	sysUserRoleRows                = strings.Join(sysUserRoleFieldNames, ",")
	sysUserRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysUserRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheUsercenterSysUserRoleIdPrefix = "cache:usercenter:sysUserRole:id:"
)

type (
	sysUserRoleModel interface {
		Insert(ctx context.Context, data *SysUserRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUserRole, error)
		Update(ctx context.Context, data *SysUserRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysUserRoleModel struct {
		sqlc.CachedConn
		table string
	}

	SysUserRole struct {
		Id     int64 `db:"id"`
		UserId int64 `db:"user_id"`
		RoleId int64 `db:"role_id"`
	}
)

func newSysUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysUserRoleModel {
	return &defaultSysUserRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_user_role`",
	}
}

func (m *defaultSysUserRoleModel) Delete(ctx context.Context, id int64) error {
	usercenterSysUserRoleIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserRoleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterSysUserRoleIdKey)
	return err
}

func (m *defaultSysUserRoleModel) FindOne(ctx context.Context, id int64) (*SysUserRole, error) {
	usercenterSysUserRoleIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserRoleIdPrefix, id)
	var resp SysUserRole
	err := m.QueryRowCtx(ctx, &resp, usercenterSysUserRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRoleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultSysUserRoleModel) Insert(ctx context.Context, data *SysUserRole) (sql.Result, error) {
	usercenterSysUserRoleIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserRoleIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserRoleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.RoleId)
	}, usercenterSysUserRoleIdKey)
	return ret, err
}

func (m *defaultSysUserRoleModel) Update(ctx context.Context, data *SysUserRole) error {
	usercenterSysUserRoleIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserRoleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRoleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.RoleId, data.Id)
	}, usercenterSysUserRoleIdKey)
	return err
}

func (m *defaultSysUserRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsercenterSysUserRoleIdPrefix, primary)
}

func (m *defaultSysUserRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRoleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysUserRoleModel) tableName() string {
	return m.table
}
