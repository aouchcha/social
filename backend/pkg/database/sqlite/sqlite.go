package sqlite

import (
	"database/sql"
	"path/filepath"
	"socialNetwork/pkg/config"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb(c *config.Database) (*sql.DB, error) {
	migrationsPath, _ := filepath.Abs("./pkg/database/migrations/sqlite")

	// when running on linux remove this line
	migrationsPath = strings.ReplaceAll(migrationsPath, "\\", "/")

	m, err := migrate.New("file://"+migrationsPath, "sqlite://database.db")
	if err != nil {
		return nil, err
	}
	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	database, err := sql.Open(c.Driver, c.FileName)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
