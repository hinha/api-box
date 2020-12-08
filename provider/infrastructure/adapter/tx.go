package adapter

import (
	"context"
	"database/sql"
	"github.com/hinha/api-box/provider"
)

// A Tx adapter for golang sql
type Tx struct {
	tx *sql.Tx
}

// AdaptTx do adapting mysql transaction
func AdaptTx(tx *sql.Tx) *Tx {
	return &Tx{tx: tx}
}

func (t *Tx) ExecContext(ctx context.Context, queryKey, query string, args ...interface{}) (provider.Result, error) {
	var (
		result provider.Result
		err    error
	)

	_ = runWithSQLAnalyzer(ctx, "tx", func() error {
		result, err = t.tx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})

	return result, nil
}

func (t *Tx) QueryContext(ctx context.Context, queryKey, query string, args ...interface{}) (provider.Rows, error) {
	var (
		rows provider.Rows
		err  error
	)

	_ = runWithSQLAnalyzer(ctx, "tx", func() error {
		rows, err = t.tx.QueryContext(ctx, query, args...)
		if err == sql.ErrNoRows {
			err = provider.ErrDBNotFound
			return provider.ErrDBNotFound
		} else if err != nil {
			return err
		}

		return nil
	})

	return rows, err
}

func (t *Tx) QueryRowContext(ctx context.Context, queryKey, query string, args ...interface{}) provider.Row {
	var row provider.Row
	_ = runWithSQLAnalyzer(ctx, "tx", func() error {
		row = t.tx.QueryRowContext(ctx, query, args...)
		return nil
	})

	return row
}
