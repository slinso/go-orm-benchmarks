package bench

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	"github.com/efectn/go-orm-benchmarks/bench/codegen"
	"github.com/efectn/go-orm-benchmarks/bench/codegen/db/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Codegen struct {
	db *sql.DB
}

func CreateCodegen() helper.ORMInterface {
	return &Codegen{}
}

func (cgen *Codegen) Name() string {
	return "codegen"
}

func (cgen *Codegen) Init() error {
	var err error
	cgen.db, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (cgen *Codegen) Close() error {
	return cgen.db.Close()
}

func (cgen *Codegen) Insert(b *testing.B) {
	m := &codegen.Models{}
	ctx := codegen.NewContext()
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

func (cgen *Codegen) InsertMulti(b *testing.B) {
}

func (cgen *Codegen) Update(b *testing.B) {
	m := &codegen.Models{}
	ctx := codegen.NewContext()
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

func (cgen *Codegen) Read(b *testing.B) {
	m := &codegen.Models{}
	ctx := codegen.NewContext()
	store := mysql.NewModelsStore(ctx, cgen.db)

	err := store.Insert(m)
	if err != nil {
		fmt.Println(err, "Insert")
		helper.SetError(b, cgen.Name(), "Insert", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := store.OneByID(1)
		if err != nil {
			fmt.Println(err, "Read")
			helper.SetError(b, cgen.Name(), "Read", err.Error())
		}
	}
}

func (cgen *Codegen) ReadSlice(b *testing.B) {
	m := &codegen.Models{}
	ctx := codegen.NewContext()
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

	var data []*codegen.Models
	var err error
	for i := 0; i < b.N; i++ {
		data, err = store.Where("id > 0").Limit(100).Query()
		if err != nil {
			helper.SetError(b, cgen.Name(), "ReadSlice", err.Error())
		}
	}

	fmt.Println("codgen len", len(data))
}
