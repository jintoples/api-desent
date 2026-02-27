package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if r := recover(); r != nil {
		err := tx.Rollback()
		PanicIfError(err)
		panic(r)
	} else {
		err := tx.Commit()
		PanicIfError(err)
	}
}
