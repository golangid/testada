package mysql

import (
	"database/sql"
	"strings"

	"github.com/golang-migrate/migrate"
	_mysql "github.com/golang-migrate/migrate/database/mysql"
)

type migration struct {
	Migrate *migrate.Migrate
}

func (this *migration) Up() ([]error, bool) {
	err := this.Migrate.Up()
	if err != nil {
		return []error{err}, false
	}

	return []error{}, true
}

func (this *migration) Down() ([]error, bool) {
	err := this.Migrate.Down()
	if err != nil {
		return []error{err}, false
	}

	return []error{}, true
}

func runMigration(databaseName, migrationsFolderLocation, dbURI string) (*migration, error) {
	dataPath := []string{}
	dataPath = append(dataPath, "file://")
	dataPath = append(dataPath, migrationsFolderLocation)

	pathToMigrate := strings.Join(dataPath, "")

	db, err := sql.Open(mysql, dbURI)
	if err != nil {
		return nil, err
	}

	driver, err := _mysql.WithInstance(db, &_mysql.Config{DatabaseName: databaseName})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		pathToMigrate,
		mysql,
		driver,
	)

	if err != nil {
		return nil, err
	}

	return &migration{
		Migrate: m,
	}, nil
}
