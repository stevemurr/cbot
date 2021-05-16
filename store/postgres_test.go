package store_test

import (
	"testing"

	"github.com/stevemurr/cbot/store"
)

func TestPostgres(t *testing.T) {
	db := store.NewPostgres("192.168.1.99", "murr", "flag9012", 5432, "cbot")
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	if err := db.Insert("123.0", "now()", "abc"); err != nil {
		t.Fatal(err)
	}
}
