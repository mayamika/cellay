package gamesstorage

import (
	"context"

	"github.com/genjidb/genji"
)

const migrationScript string = `
CREATE TABLE IF NOT EXISTS games;
`

func applyMigrations(ctx context.Context, db *genji.DB) error {
	return inTx(ctx, db, func(tx *genji.Tx) error {
		return tx.Exec(migrationScript)
	})
}
