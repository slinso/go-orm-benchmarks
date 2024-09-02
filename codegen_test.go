package main_test

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/bench/codegen"
	"github.com/efectn/go-orm-benchmarks/bench/codegen/db/mysql"
	"github.com/gkampitakis/go-snaps/snaps"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestCodegen_ReadSlice(t *testing.T) {
	var err error
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	m := &codegen.Models{}
	ctx := codegen.NewContext()
	store := mysql.NewModelsStore(ctx, db)

	for i := 0; i < 100; i++ {
		m.ID = 0
		err := store.Insert(m)
		if err != nil {
			t.Fatal(err)
		}
	}

	data, err := store.Limit(100).Query()
	if err != nil {
		t.Fatal(err)
	}

	snaps.MatchSnapshot(t, data)
}
