package shrd_utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	sql_db "github.com/StevanoZ/dv-shared/db"
)

const rowCloseErrorMsg = "pq: unexpected Parse response 'C'"
const deadLockErrorMsg = "pq: unexpected Parse response 'D'"
const badConnectionErrMsg = "driver: bad connection"

func ExecTx(ctx context.Context, DB sql_db.DBInterface, fn func(tx *sql.Tx) error) error {
	tx, err := DB.BeginTx(ctx, nil)
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

func ExecTxWithRetry(ctx context.Context, DB sql_db.DBInterface, fn func(tx *sql.Tx) error) error {
	var retryFunc = func() error {
		tx, err := DB.BeginTx(ctx, nil)
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
			strings.Contains(err.Error(), rowCloseErrorMsg) {
			time.Sleep(1 * time.Second)
			log.Printf("retry transaction %d times \n", i+1)
			err = retryFunc()
		} else {
			// DON'T NEED TO RETRY THIS ERROR
			break
		}
	}

	return err
}
