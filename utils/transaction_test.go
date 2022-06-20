package shrd_utils

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func loadConfigAndSetUpDb() *sql.DB {
	config := LoadBaseConfig("../app", "test")

	return ConnectDB(config.DBDriver, config.DBSource)
}

func TestExecTx(t *testing.T) {
	db := loadConfigAndSetUpDb()

	t.Run("Success commit Tx", func(t *testing.T) {
		err := ExecTx(context.Background(), db, func(tx *sql.Tx) error {
			return nil
		}, 6)

		assert.NoError(t, err)
	})

	t.Run("Failed when creating Tx", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		cancel()
		err := ExecTx(ctx, db, func(tx *sql.Tx) error {
			return nil
		})
		assert.Error(t, err)
	})

	t.Run("Success roll back Tx", func(t *testing.T) {
		err := ExecTx(context.Background(), db, func(tx *sql.Tx) error {
			return errors.New("Roll Back")
		})
		assert.Error(t, err)
	})

	t.Run("Failed roll back Tx", func(t *testing.T) {
		err := ExecTx(context.Background(), db, func(tx *sql.Tx) error {
			db.Close()
			tx.Rollback()
			return errors.New("Roll back error")
		})
		assert.Error(t, err)
	})
}

func TestExecTxWithRetry(t *testing.T) {
	db := loadConfigAndSetUpDb()

	t.Run("Success commit Tx", func(t *testing.T) {
		err := ExecTxWithRetry(context.Background(), db, func(tx *sql.Tx) error {
			return nil
		}, 6)

		assert.NoError(t, err)
	})

	t.Run("Failed and should retry tx (bad connection error)", func(t *testing.T) {
		err := ExecTxWithRetry(context.Background(), db, func(tx *sql.Tx) error {
			tx.Rollback()
			return errors.New(badConnectionErrMsg)
		})
		assert.Error(t, err)
	})

	t.Run("Failed when creating Tx", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		cancel()
		err := ExecTxWithRetry(ctx, db, func(tx *sql.Tx) error {
			return nil
		})
		assert.Error(t, err)
	})

	t.Run("Success roll back Tx", func(t *testing.T) {
		err := ExecTxWithRetry(context.Background(), db, func(tx *sql.Tx) error {
			return errors.New("Roll Back")
		})
		assert.Error(t, err)
	})

	t.Run("Failed roll back Tx", func(t *testing.T) {
		err := ExecTxWithRetry(context.Background(), db, func(tx *sql.Tx) error {
			db.Close()
			tx.Rollback()
			return errors.New("Roll back error")
		})
		assert.Error(t, err)
	})
}
