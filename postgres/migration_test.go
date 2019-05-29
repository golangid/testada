package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	PostgresSuite
}

func TestPostgresSuite(t *testing.T) {
	s := SuiteTest{
		PostgresSuite{
			DSN: "root:root@tcp(postgres:5432)/postgres_test",
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
