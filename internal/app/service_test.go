package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/markLishansky/url-short/internal/store"
	migration "github.com/markLishansky/url-short/sql"
	"log"
	"testing"
)

type testSuite struct {
	suite.Suite
	store store.DataStore
}

func (t *testSuite) SetupSuite() {
	dbConnectionString := "dbname=shorter_test password=admin user=admin sslmode=disable"
	connection, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalln("Failed to init db:", err)
	}
	migration.RunMigrations(connection)
	t.store = store.CreateDbProvider(connection)
}

func TestName(t *testing.T) {
	suite.Run(t, new(testSuite))
}
