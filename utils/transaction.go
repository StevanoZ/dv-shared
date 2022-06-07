package shrd_utils

import (
	"context"
	"database/sql"
	"fmt"

	sql_db "github.com/StevanoZ/dv-shared/db"
)

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
