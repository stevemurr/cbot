package store_test

import (
	"os"
	"testing"

	"github.com/stevemurr/cbot/store"
)

func TestPostgres(t *testing.T) {
	db := store.NewPostgres(
		os.Getenv("cbot_postgres_host"),
		os.Getenv("cbot_postgres_username"),
		os.Getenv("cbot_postgres_password"),
		5432,
		"cbot")
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	if err := db.Insert("123.0", "now()", "abc"); err != nil {
		t.Fatal(err)
	}
}
