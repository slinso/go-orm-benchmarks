package bench

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	"github.com/efectn/go-orm-benchmarks/bench/newgen"
	"github.com/efectn/go-orm-benchmarks/bench/newgen/db/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type NewGen struct {
	db *sql.DB
}

func CreateNewGen() helper.ORMInterface {
	return &NewGen{}
}

func (cgen *NewGen) Name() string {
	return "NewGen"
}

func (cgen *NewGen) Init() error {
	var err error
	cgen.db, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (cgen *NewGen) Close() error {
	return cgen.db.Close()
}

func (cgen *NewGen) Insert(b *testing.B) {
	m := &newgen.Models{}
	ctx := newgen.NewContext()
	store := mysql.NewModelsStore(ctx, cgen.db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := store.Insert(m)
		if err != nil {
			helper.SetError(b, cgen.Name(), "Insert", err.Error())
		}
	}
}

func (cgen *NewGen) InsertMulti(b *testing.B) {
}

func (cgen *NewGen) Update(b *testing.B) {
	m := &newgen.Models{}
	ctx := newgen.NewContext()
	store := mysql.NewModelsStore(ctx, cgen.db)

	err := store.Insert(m)
	if err != nil {
		helper.SetError(b, cgen.Name(), "Insert", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := store.Update(m)
		if err != nil {
			helper.SetError(b, cgen.Name(), "Update", err.Error())
		}
	}
}

func (cgen *NewGen) Read(b *testing.B) {
	m := &newgen.Models{}
	ctx := newgen.NewContext()
	store := mysql.NewModelsStore(ctx, cgen.db)

	err := store.Insert(m)
	if err != nil {
		fmt.Println(err, "Insert")
		helper.SetError(b, cgen.Name(), "Insert", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// _, err := store.OneByID(1)
		// if err != nil {
		// 	fmt.Println(err, "Read")
		// 	helper.SetError(b, cgen.Name(), "Read", err.Error())
		// }
	}
}

func (cgen *NewGen) ReadSlice(b *testing.B) {
	m := &newgen.Models{}
	ctx := newgen.NewContext()
	store := mysql.NewModelsStore(ctx, cgen.db)

	for i := 0; i < 100; i++ {
		m.ID = 0
		err := store.Insert(m)
		if err != nil {
			helper.SetError(b, cgen.Name(), "Insert", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := store.Limit(100).Query()
		if err != nil {
			helper.SetError(b, cgen.Name(), "ReadSlice", err.Error())
		}
	}
}
