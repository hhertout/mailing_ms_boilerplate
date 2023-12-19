package database

import (
	"database/sql"
	"io"
	"os"
	"strings"
)

type Migration struct {
	dbPool   *sql.DB
	basePath string
}

func NewMigration(db *sql.DB, basePath string) *Migration {
	return &Migration{
		db,
		basePath,
	}
}

func (m Migration) Migrate() error {
	if err := m.migrateFromFile("mail.sql"); err != nil {
		return err
	}
	return nil
}

func (m Migration) migrateFromFile(filename string) error {
	workingDir, _ := os.Getwd()
	fileOpen, err := os.Open(workingDir + m.basePath + "/cmd/database/migrations/" + filename)
	if err != nil {
		return err
	}
	defer fileOpen.Close()

	content, err := io.ReadAll(fileOpen)
	if err != nil {
		return err
	}

	queries := string(content)
	queriesSplit := strings.Split(queries, "--")

	for _, query := range queriesSplit {
		if strings.TrimSpace(query) == "" {
			continue
		}

		_, err = m.dbPool.Exec(query + ";")
		if err != nil {
			if err.Error() == "trigger set_viewed_param already exists" {
				continue
			}
			return err
		}
	}
	return nil
}
