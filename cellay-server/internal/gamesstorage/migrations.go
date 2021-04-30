package gamesstorage

import (
	"context"
	"fmt"

	"github.com/genjidb/genji"
)

const migrationScript string = `
CREATE TABLE IF NOT EXISTS games;
`

func applyMigrations(ctx context.Context, db *genji.DB) error {
	return inTx(ctx, db, func(tx *genji.Tx) error {
		if err := tx.Exec(migrationScript); err != nil {
			return fmt.Errorf("can't exec migration script: %w", err)
		}
		return nil
	})
}
