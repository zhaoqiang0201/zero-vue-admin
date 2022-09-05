// Code generated by goctl. DO NOT EDIT!

package sysuser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserFieldNames          = builder.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheUsercenterSysUserIdPrefix   = "cache:usercenter:sysUser:id:"
	cacheUsercenterSysUserNamePrefix = "cache:usercenter:sysUser:name:"
)

type (
	sysUserModel interface {
		Insert(ctx context.Context, data *SysUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUser, error)
		FindOneByName(ctx context.Context, name string) (*SysUser, error)
		Update(ctx context.Context, data *SysUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysUserModel struct {
		sqlc.CachedConn
		table string
	}

	SysUser struct {
		Id            int64          `db:"id"`             // 编号
		Name          string         `db:"name"`           // 用户名
		NickName      string         `db:"nick_name"`      // 昵称
		Avatar        string         `db:"avatar"`         // 头像
		Password      string         `db:"password"`       // 密码
		Email         string         `db:"email"`          // 邮箱
		Mobile        string         `db:"mobile"`         // 手机号
		Status        uint64         `db:"status"`         // 状态  0：禁用   1：正常
		DefaultRouter string         `db:"default_router"` // 默认路由
		CreateBy      string         `db:"create_by"`      // 创建人
		CreateAt      time.Time      `db:"create_at"`      // 创建时间
		UpdateBy      string         `db:"update_by"`      // 更新人
		UpdateAt      time.Time      `db:"update_at"`      // 更新时间
		DeleteBy      sql.NullString `db:"delete_by"`      // 删除人
		DeleteAt      sql.NullTime   `db:"delete_at"`
	}
)

func newSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysUserModel {
	return &defaultSysUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_user`",
	}
}

func (m *defaultSysUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usercenterSysUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserIdPrefix, id)
	usercenterSysUserNameKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterSysUserIdKey, usercenterSysUserNameKey)
	return err
}

func (m *defaultSysUserModel) FindOne(ctx context.Context, id int64) (*SysUser, error) {
	usercenterSysUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserIdPrefix, id)
	var resp SysUser
	err := m.QueryRowCtx(ctx, &resp, usercenterSysUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
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

func (m *defaultSysUserModel) FindOneByName(ctx context.Context, name string) (*SysUser, error) {
	usercenterSysUserNameKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserNamePrefix, name)
	var resp SysUser
	err := m.QueryRowIndexCtx(ctx, &resp, usercenterSysUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", sysUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
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

func (m *defaultSysUserModel) Insert(ctx context.Context, data *SysUser) (sql.Result, error) {
	usercenterSysUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserIdPrefix, data.Id)
	usercenterSysUserNameKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.NickName, data.Avatar, data.Password, data.Email, data.Mobile, data.Status, data.DefaultRouter, data.CreateBy, data.UpdateBy, data.DeleteBy, data.DeleteAt)
	}, usercenterSysUserIdKey, usercenterSysUserNameKey)
	return ret, err
}

func (m *defaultSysUserModel) Update(ctx context.Context, newData *SysUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usercenterSysUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserIdPrefix, data.Id)
	usercenterSysUserNameKey := fmt.Sprintf("%s%v", cacheUsercenterSysUserNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.NickName, newData.Avatar, newData.Password, newData.Email, newData.Mobile, newData.Status, newData.DefaultRouter, newData.CreateBy, newData.UpdateBy, newData.DeleteBy, newData.DeleteAt, newData.Id)
	}, usercenterSysUserIdKey, usercenterSysUserNameKey)
	return err
}

func (m *defaultSysUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsercenterSysUserIdPrefix, primary)
}

func (m *defaultSysUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysUserModel) tableName() string {
	return m.table
}
