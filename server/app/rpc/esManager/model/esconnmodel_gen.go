// Code generated by goctl. DO NOT EDIT!

package model

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
	esConnFieldNames          = builder.RawFieldNames(&EsConn{})
	esConnRows                = strings.Join(esConnFieldNames, ",")
	esConnRowsExpectAutoSet   = strings.Join(stringx.Remove(esConnFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	esConnRowsWithPlaceHolder = strings.Join(stringx.Remove(esConnFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheChaosEsEsConnIdPrefix = "cache:chaosEs:esConn:id:"
)

type (
	esConnModel interface {
		Insert(ctx context.Context, data *EsConn) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*EsConn, error)
		Update(ctx context.Context, data *EsConn) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultEsConnModel struct {
		sqlc.CachedConn
		table string
	}

	EsConn struct {
		Id       uint64         `db:"id"`
		EsConn   string         `db:"es_conn"`
		Version  int64          `db:"version"`
		User     sql.NullString `db:"user"`
		Password sql.NullString `db:"password"`
		Describe string         `db:"describe"`
	}
)

func newEsConnModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultEsConnModel {
	return &defaultEsConnModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`es_conn`",
	}
}

func (m *defaultEsConnModel) Delete(ctx context.Context, id uint64) error {
	chaosEsEsConnIdKey := fmt.Sprintf("%s%v", cacheChaosEsEsConnIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, chaosEsEsConnIdKey)
	return err
}

func (m *defaultEsConnModel) FindOne(ctx context.Context, id uint64) (*EsConn, error) {
	chaosEsEsConnIdKey := fmt.Sprintf("%s%v", cacheChaosEsEsConnIdPrefix, id)
	var resp EsConn
	err := m.QueryRowCtx(ctx, &resp, chaosEsEsConnIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", esConnRows, m.table)
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

func (m *defaultEsConnModel) Insert(ctx context.Context, data *EsConn) (sql.Result, error) {
	chaosEsEsConnIdKey := fmt.Sprintf("%s%v", cacheChaosEsEsConnIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, esConnRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EsConn, data.Version, data.User, data.Password, data.Describe)
	}, chaosEsEsConnIdKey)
	return ret, err
}

func (m *defaultEsConnModel) Update(ctx context.Context, data *EsConn) error {
	chaosEsEsConnIdKey := fmt.Sprintf("%s%v", cacheChaosEsEsConnIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, esConnRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.EsConn, data.Version, data.User, data.Password, data.Describe, data.Id)
	}, chaosEsEsConnIdKey)
	return err
}

func (m *defaultEsConnModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheChaosEsEsConnIdPrefix, primary)
}

func (m *defaultEsConnModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", esConnRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultEsConnModel) tableName() string {
	return m.table
}
