package postgres

import (
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

// OpenDB データベースと接続する
func OpenDB(datasource string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", datasource)
	if err != nil {
		return nil, xerrors.Errorf("failed open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, xerrors.Errorf("failed connect database: %w", err)
	}
	return db, nil
}

// MigrateDB データベースをマイグレーションする
func MigrateDB(datasource string) error {
	slog.Info("start migration")
	db, err := OpenDB(datasource)
	if err != nil {
		return xerrors.Errorf("failed database open: %w", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("failed close database", "err", err)
		}
	}()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return xerrors.Errorf("failed make postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./schema/postgres/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return xerrors.Errorf("failed read migration file: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return xerrors.Errorf("failed migration: %w", err)
	}

	// zap.L().Info("データベースのマイグレーションを終了します")
	slog.Info("finished database migration")
	return nil
}
