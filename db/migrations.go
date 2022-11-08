package db

import (
	"context"
	"database/sql"
	"embed"
	"io"
	"path/filepath"
	"sort"

	"github.com/rs/zerolog"
)

//go:embed sql/migrations/*.sql
var migrationFiles embed.FS

const migrationsPath = "sql/migrations"

const (
	createMigrationsTableSQL = `
		CREATE TABLE IF NOT EXISTS "migration" (
			"name"        TEXT UNIQUE,
			"time"        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`
	getAllRanMigrationNamesSQL = `SELECT name FROM migration`
	markMigrationAsRanSQL      = `INSERT INTO "migration" ("name") VALUES (?)`
)

func getAllRanMigration(ctx context.Context, tx *sql.Tx) (map[string]struct{}, error) {
	res, err := tx.QueryContext(ctx, getAllRanMigrationNamesSQL)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	out := make(map[string]struct{})

	for res.Next() {
		var name string
		err = res.Scan(&name)
		if err != nil {
			return nil, err
		}

		out[name] = struct{}{}
	}
	if err = res.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func markMigrationAsRan(ctx context.Context, tx *sql.Tx, name string) error {
	_, err := tx.ExecContext(ctx, markMigrationAsRanSQL, name)
	return err
}

type migrationData struct {
	Name    string
	Content string
}

func getAllMigrations() ([]migrationData, error) {
	migrations, err := migrationFiles.ReadDir(migrationsPath)
	if err != nil {
		return nil, err
	}

	out := make([]migrationData, 0, len(migrations))

	for _, migration := range migrations {
		name := migration.Name()

		fp, err := migrationFiles.Open(filepath.Join(migrationsPath, name))
		if err != nil {
			return nil, err
		}

		content, err := io.ReadAll(fp)
		if err != nil {
			return nil, err
		}

		out = append(out, migrationData{name, string(content)})
	}

	// make sure migrations are in correct order
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})

	return out, nil
}

func (db Database) runMigrations(ctx context.Context) error {
	return withTransaction2(ctx, db.db, func(ctx context.Context, tx *sql.Tx) error {
		logger := zerolog.Ctx(ctx)

		_, err := tx.ExecContext(ctx, createMigrationsTableSQL)
		if err != nil {
			logger.Err(err).Msg("failed to create migrations table")
			return err
		}

		allMigrationNames, err := getAllRanMigration(ctx, tx)
		if err != nil {
			logger.Err(err).Msg("failed to get all ran migrations")
			return err
		}

		migrations, err := getAllMigrations()
		if err != nil {
			logger.Err(err).Msg("failed to get all migrations")
			return err
		}

		for _, migration := range migrations {
			l := logger.With().Str("migration", migration.Name).Logger()

			l.Info().Msg("running migration")

			_, ok := allMigrationNames[migration.Name]
			if ok {
				l.Info().Msg("skipping migration because it has already been ran")
				continue
			}

			_, err = tx.ExecContext(ctx, migration.Content)
			if err != nil {
				l.Err(err).Msg("failed to execute migration")
				return err
			}

			err = markMigrationAsRan(ctx, tx, migration.Name)
			if err != nil {
				l.Err(err).Msg("failed to mark migration as ran")
				return err
			}
		}

		return nil
	})
}
