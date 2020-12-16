package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
)

// ApplyAll applies all migration to the db database in dir directory
func ApplyAll(db *sql.DB, dir string) (err error) {
	migrations, err := getMigrations(dir)
	if err != nil {
		return
	}
	for _, name := range migrations {
		err = apply(db, path.Join(dir, name))
		if err != nil {
			return
		}
	}
	return
}

// apply applies a migration in a given file to the database
func apply(db *sql.DB, path string) error {
	migration, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file at \"%s\"", path)
	}

	if _, err := db.Exec(string(migration)); err != nil {
		return fmt.Errorf("error applying migration \"%s\": %s", path, err)
	}
	return nil
}

func getMigrations(dir string) (out []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, f := range files {
		out = append(out, f.Name())
	}
	return
}
