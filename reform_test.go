package main_test

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/bench"
	"github.com/efectn/go-orm-benchmarks/bench/reform"
	"github.com/gkampitakis/go-snaps/snaps"
	_ "github.com/jackc/pgx/v4/stdlib"
	reformware "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

func TestReform_ReadSlice(t *testing.T) {
	var err error
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	conn := reformware.NewDB(db, postgresql.Dialect, nil)

	m := bench.NewReformModel()
	for i := 0; i < 100; i++ {
		err := conn.Save(m)
		if err != nil {
			t.Fatal(err)
		}
	}

	data, err := conn.SelectAllFrom(reform.ReformModelsTable, "WHERE id > 0 LIMIT 100")
	if err != nil {
		t.Fatal(err)
	}

	snaps.MatchSnapshot(t, data)
}
