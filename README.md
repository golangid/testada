# testada
This package is just a boilerplate for doing integration testing in Golang. We already separate a few code that may same when doing integration testing in every test-case. 

## Index

* [Support](#support)
* [How To Use](#how-to-use) 
* [Contribution](#contribution)


## Support

You can file an [Issue](https://github.com/golangid/testada/issues/new).
See documentation in [Godoc](https://godoc.org/github.com/golangid/testada)

## How To Use

### Prerequisite Library and Tools
- A live DB on your project, you could use docker to spawn a live DB/services
- `github.com/stretchr/testify`. A testing package, really usefull. And all of this package use this package. Make sure understand how to use it, before using our boilerplate test-suite. 
- DB drivers, depend on what driver you use. 

#### List DB driver and libraries that already supported here
|services| driver and libraries |testada-package|
|--------|--------|---------------|
| Mysql  | <ul> <li>github.com/go-sql-driver/mysql </li><li> sql/db</li></ul> | github.com/testada/mysql|
| Redis  | github.com/go-redis/redis | github.com/golangid/testada/go-redis |

## Usage In MYSQL
Complete file can be seen in: [example-testada-mysql](https://github.com/golangid/testada-example/blob/master/mysql/repository_test.go)

```go
import (
	// ... other imports
	"github.com/stretchr/testify/suite"
	"github.com/golangid/testada/mysql"
)

type mysqlCategorySuiteTest struct {
	mysql.MysqlSuite
}

func TestCategorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip category mysql repository test")
	}
	dsn := os.Getenv("MYSQL_TEST_URL")
	if dsn == "" {
		dsn = "root:root-pass@tcp(localhost:33060)/testing?parseTime=1&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	}
	categorySuite := &mysqlCategorySuiteTest{
		mysql.MysqlSuite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}

	suite.Run(t, categorySuite)
}

func (s *mysqlCategorySuiteTest) SetupTest() {
	log.Println("Starting a Test. Migrating the Database")
	err, _ := s.Migration.Up()
	require.NoError(s.T(), err)
	log.Println("Database Migrated Successfully")
}

func (s *mysqlCategorySuiteTest) TearDownTest() {
	log.Println("Finishing Test. Dropping The Database")
	err, _ := s.Migration.Down()
	require.NoError(s.T(), err)
	log.Println("Database Dropped Successfully")
}

func (m *mysqlCategorySuiteTest) TestStore() {
  // Your test code will be placed here
  // This function will do the integration-test for Store function.
  // Your Store function will test directly with a real DB by this TestFunction
}
func (m *mysqlCategorySuiteTest) TestOtherFunction() {
  // Your test code will be placed here
  // This function will do the integration-test for your defined function as you want.
  // Your Store function will test directly with a real DB by this TestFunction
}
// ... Add more test according to your cases

```
## Usage In Redis with GO-REDIS package
This example use this driver: github.com/go-redis/redis 
Complete file can be seen in: [example-testada-redis](https://github.com/golangid/testada-example/blob/master/redis/cache_test.go)

```go
import(
	// ... other imports
	"github.com/stretchr/testify/suite"
	goRedisSuite "github.com/golangid/testada/go-redis"
)
type redisHandlerSuite struct {
	goRedisSuite.RedisSuite
}

func TestRedisSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for redis repository")
	}
	redisHostTest := os.Getenv("REDIS_TEST_URL")
	if redisHostTest == "" {
		redisHostTest = "localhost:6379"
	}
	redisHandlerSuiteTest := &redisHandlerSuite{
		goRedisSuite.RedisSuite{
			Host: redisHostTest,
		},
	}
	suite.Run(t, redisHandlerSuiteTest)
}

func getItemByKey(client *redis.Client, key string) ([]byte, error) {
	return client.Get(key).Bytes()
}
func seedItem(client *redis.Client, key string, value interface{}) error {
	jybt, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(key, jybt, time.Second*30).Err()
}
func (r *redisHandlerSuite) TestSet() {
  // Your test code will be placed here
  // This function will do the integration-test for your defined function as you want.
  // Your Store function will test directly with a real DB by this TestFunction
}
func (r *redisHandlerSuite) TestGet() {
  // Your test code will be placed here
  // This function will do the integration-test for your defined function as you want.
  // Your Store function will test directly with a real DB by this TestFunction
}
// ... Add more test according to your cases
```
