package postgres

import (
	"database/sql"

	driverSql "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const postgres = "postgres"

// PostgresSuite struct for MySQL Suite
type PostgresSuite struct {
	suite.Suite
	DSN                     string
	DBConn                  *sql.DB
	Migration               *migration
	MigrationLocationFolder string
	DBName                  string
}

// SetupSuite setup at the beginning of test
func (s *PostgresSuite) SetupSuite() {
	DisableLogging()

	var err error

	s.DBConn, err = sql.Open(postgres, s.DSN)
	err = s.DBConn.Ping()
	require.NoError(s.T(), err)

	s.Migration, err = runMigration(s.DBConn, s.MigrationLocationFolder)
	require.NoError(s.T(), err)
}

// TearDownSuite teardown at the end of test
func (s *PostgresSuite) TearDownSuite() {
	err := s.DBConn.Close()
	require.NoError(s.T(), err)
}

func DisableLogging() {
	nopLogger := NopLogger{}
	if err := driverSql.SetLogger(nopLogger); err != nil {
		panic(err)
	}

}

type NopLogger struct {
}

func (l NopLogger) Print(v ...interface{}) {}
