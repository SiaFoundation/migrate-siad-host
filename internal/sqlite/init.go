package sqlite

import (
	_ "embed" // for init.sql

	"fmt"

	"go.sia.tech/core/types"
	"lukechampine.com/frand"
)

// init queries are run when the database is first created.
//
//go:embed init.sql
var initDatabase string

func (s *Store) initNewDatabase(target int64) error {
	return s.transaction(func(tx txn) error {
		if _, err := tx.Exec(initDatabase); err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		} else if err := setDBVersion(tx, target); err != nil {
			return fmt.Errorf("failed to set initial database version: %w", err)
		} else if err = generateHostKey(tx); err != nil {
			return fmt.Errorf("failed to generate host key: %w", err)
		}
		return nil
	})
}

func (s *Store) init() error {
	const target = 24 // hardcoded target from the v1.0.0 release

	// disable foreign key constraints during migration
	if _, err := s.db.Exec("PRAGMA foreign_keys = OFF"); err != nil {
		return fmt.Errorf("failed to disable foreign key constraints: %w", err)
	}

	version := getDBVersion(s.db)
	switch {
	case version == 0:
		return s.initNewDatabase(target)
	case version < target:
		return fmt.Errorf("database version %v is older than expected %v. database upgrades are not supported", version, target)
	case version > target:
		return fmt.Errorf("database version %v is newer than expected %v. database downgrades are not supported", version, target)
	}
	// nothing to do
	return nil
}

func generateHostKey(tx txn) (err error) {
	key := types.NewPrivateKeyFromSeed(frand.Bytes(32))
	var dbID int64
	err = tx.QueryRow(`UPDATE global_settings SET host_key=? RETURNING id`, key).Scan(&dbID)
	return
}
