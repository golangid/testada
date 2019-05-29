package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	MysqlSuite
}

func TestMysqlSuite(t *testing.T) {
	s := SuiteTest{
		MysqlSuite{
			DSN: "root:root@tcp(mysql:3306)/mysql_test",
			MigrationLocationFolder: "migrations",
			DBName:                  "feed",
		},
	}

	suite.Run(t, &s)
}

func (s *SuiteTest) TestRunMigrationUp() {
	errs, ok := s.Migration.Up()
	assert.True(s.T(), ok)
	assert.Len(s.T(), errs, 0)
}

func (s *SuiteTest) TestRunMigrationDown() {
	errs, ok := s.Migration.Up()
	assert.True(s.T(), ok)
	assert.Len(s.T(), errs, 0)
}
