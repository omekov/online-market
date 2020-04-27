package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = "host=localhost port=5431 user=azamat password=azamat dbname=marketdb_test sslmode=disable"
	os.Exit(m.Run())
}
