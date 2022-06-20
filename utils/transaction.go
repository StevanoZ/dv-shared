package shrd_utils

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sql_db "github.com/StevanoZ/dv-shared/db"
)

// pq: could not serialize access due to concurrent update
// The transaction might succeed if retried.
const rowCloseErrorMsg = "pq: unexpected Parse response 'C'"
const deadLockErrorMsg = "pq: unexpected Parse response 'D'"
const badConnectionErrMsg = "driver: bad connection"
const txAbortingErrMsg = "pq: Could not complete operation in a failed transaction"

func ExecTx(ctx context.Context, DB sql_db.DBInterface, fn func(tx *sql.Tx) error, level ...int) error {
	isolationLevel := 0
	if len(level) > 0 {
		isolationLevel = level[0]
	}

	tx, err := DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.IsolationLevel(isolationLevel)})
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func ExecTxWithRetry(ctx context.Context, DB sql_db.DBInterface, fn func(tx *sql.Tx) error, level ...int) error {
	var retryFunc = func() error {
		isolationLevel := 0
		if len(level) > 0 {
			isolationLevel = level[0]
		}

		tx, err := DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.IsolationLevel(isolationLevel)})
		if err != nil {
			return err
		}

		err = fn(tx)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
			return err
		}

		return tx.Commit()
	}

	err := retryFunc()
	for i := 0; i < 3; i++ {
		if err == nil {
			break
		} else if strings.Contains(err.Error(), badConnectionErrMsg) ||
			strings.Contains(err.Error(), deadLockErrorMsg) ||
			strings.Contains(err.Error(), rowCloseErrorMsg) ||
			strings.Contains(err.Error(), txAbortingErrMsg) {
			// immediately RETRY??
			//	time.Sleep(500 * time.Millisecond)
			LogInfo(fmt.Sprintf("retry transaction %d times \n", i+1))
			err = retryFunc()
		} else {
			// DON'T NEED TO RETRY THIS ERROR
			break
		}
	}

	return err
}
