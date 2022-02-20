package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx, err error) {
	if err := recover(); err != nil {
		err_rollback := tx.Rollback()
		PanicIfError(err_rollback)
		panic(err)
	} else {
		err_commit := tx.Commit()
		PanicIfError(err_commit)
	}
}
