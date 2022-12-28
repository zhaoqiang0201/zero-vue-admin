// Code generated by goctl. DO NOT EDIT!

package system

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	casbinRuleFieldNames          = builder.RawFieldNames(&CasbinRule{})
	casbinRuleRows                = strings.Join(casbinRuleFieldNames, ",")
	casbinRuleRowsExpectAutoSet   = strings.Join(stringx.Remove(casbinRuleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	casbinRuleRowsWithPlaceHolder = strings.Join(stringx.Remove(casbinRuleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	casbinRuleModel interface {
		Insert(ctx context.Context, data *CasbinRule) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*CasbinRule, error)
		Update(ctx context.Context, data *CasbinRule) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultCasbinRuleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CasbinRule struct {
		Id    uint64 `db:"id"`
		PType string `db:"p_type"`
		V0    string `db:"v0"`
		V1    string `db:"v1"`
		V2    string `db:"v2"`
		V3    string `db:"v3"`
		V4    string `db:"v4"`
		V5    string `db:"v5"`
	}
)

func newCasbinRuleModel(conn sqlx.SqlConn) *defaultCasbinRuleModel {
	return &defaultCasbinRuleModel{
		conn:  conn,
		table: "`casbin_rule`",
	}
}

func (m *defaultCasbinRuleModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCasbinRuleModel) FindOne(ctx context.Context, id uint64) (*CasbinRule, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", casbinRuleRows, m.table)
	var resp CasbinRule
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCasbinRuleModel) Insert(ctx context.Context, data *CasbinRule) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, casbinRuleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PType, data.V0, data.V1, data.V2, data.V3, data.V4, data.V5)
	return ret, err
}

func (m *defaultCasbinRuleModel) Update(ctx context.Context, data *CasbinRule) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, casbinRuleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PType, data.V0, data.V1, data.V2, data.V3, data.V4, data.V5, data.Id)
	return err
}

func (m *defaultCasbinRuleModel) tableName() string {
	return m.table
}