package base

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type key struct {
}

var keyDB = key{}

type Repository struct {
	Db *gorm.DB
}

type IBaseRepository interface {
	BeginTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context
	Commit(ctx context.Context) *gorm.DB
	Rollback(ctx context.Context) *gorm.DB
	Transaction(ctx context.Context, callback func(ctx context.Context) error) error
}

func (r *Repository) GetDB(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(keyDB).(*gorm.DB)
	if !ok {
		return r.Db
	}
	return val
}

func (r *Repository) BeginTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context {
	db := r.GetDB(ctx)
	return context.WithValue(ctx, keyDB, db.Begin(opts...))
}

func (r *Repository) Commit(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Commit()
}

func (r *Repository) Rollback(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Rollback()
}

func (r *Repository) Transaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return r.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, keyDB, tx)
		err := callback(txCtx)
		if err != nil {
			return err
		}
		return nil
	})
}
