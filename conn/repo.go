/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 16:27:31
 * @LastEditTime: 2024-05-22 14:10:59
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package conn

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type DBModel interface {
	TableName() string
}

type repo[T DBModel] struct {
}

func NewRepo[T DBModel]() *repo[T] {
	return &repo[T]{}
}

type DBCondition func(tx *gorm.DB) *gorm.DB

func (r *repo[T]) List(ctx context.Context, page, size int, conds ...DBCondition) ([]*T, int64, error) {
	var rows = make([]*T, 0)
	var count int64
	db = GetDB().Model(new(T))
	for _, v := range conds {
		db.Scopes(v)
	}
	db.Count(&count)
	err := db.Offset((page - 1) * size).Limit(size).Find(&rows).Error
	if len(rows) == 0 {
		err = fmt.Errorf("暂无数据")
	}
	r.handleErr(ctx, err)
	return rows, count, err
}

func (r *repo[T]) Show(ctx context.Context, conds ...DBCondition) (*T, error) {
	db = GetDB().Model(new(T))
	for _, v := range conds {
		db.Scopes(v)
	}
	var row T
	err := db.First(&row).Error
	r.handleErr(ctx, err)
	return &row, err
}

func (r *repo[T]) Create(ctx context.Context, data *T) error {
	err := GetDB().Model(new(T)).Create(data).Error
	r.handleErr(ctx, err)
	return err
}

func (r *repo[T]) Update(ctx context.Context, updates any, conds ...DBCondition) error {
	db = GetDB().Model(new(T))
	for _, v := range conds {
		db.Scopes(v)
	}
	err := db.Updates(updates).Error
	r.handleErr(ctx, err)
	return err
}

func (r *repo[T]) Delete(ctx context.Context, conds ...DBCondition) error {
	db = GetDB().Model(new(T))
	for _, v := range conds {
		db.Scopes(v)
	}
	err := db.Delete(new(T)).Error
	r.handleErr(ctx, err)
	return err
}

func (r *repo[T]) Search(ctx context.Context, conds ...DBCondition) ([]*T, error) {
	var rows = make([]*T, 0)
	db = GetDB().Model(new(T))
	for _, v := range conds {
		db.Scopes(v)
	}
	err := db.Find(&rows).Error
	r.handleErr(ctx, err)
	return rows, err
}

func (r *repo[T]) IsExist(ctx context.Context, conds ...DBCondition) bool {
	db = GetDB().Model(new(T))
	for _, v := range conds {
		db.Scopes(v)
	}
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (r *repo[T]) handleErr(ctx context.Context, err error) {
	if err == nil {
		return
	}
	slog.ErrorContext(ctx, "db error", "model", (*new(T)).TableName(), "err", err)
}
