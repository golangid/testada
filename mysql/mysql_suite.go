package mysql

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	driverSql "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const mysql = "mysql"

// MysqlSuite struct for MySQL Suite
type MysqlSuite struct {
	suite.Suite
	DSN                     string
	DBConn                  *sql.DB
	Migration               *migration
	MigrationLocationFolder string
	DBName                  string
}

// SetupSuite setup at the beginning of test
func (s *MysqlSuite) SetupSuite() {
	DisableLogging()

	var err error

	s.DBConn, err = sql.Open(mysql, s.DSN)
	err = s.DBConn.Ping()
	require.NoError(s.T(), err)
	_, err = s.DBConn.Exec("set global sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';")
	require.NoError(s.T(), err)
	_, err = s.DBConn.Exec("set session sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';")
	require.NoError(s.T(), err)

	s.Migration, err = runMigration(s.DBName, s.MigrationLocationFolder, s.DSN)
	require.NoError(s.T(), err)
}

// TearDownSuite teardown at the end of test
func (s *MysqlSuite) TearDownSuite() {
	s.DBConn.Close()
}

func DisableLogging() {
	nopLogger := NopLogger{}
	driverSql.SetLogger(nopLogger)
}

type NopLogger struct {
}

func (l NopLogger) Print(v ...interface{}) {
}
