package benchs

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

type Model struct {
	Id      int `orm:"auto" gorm:"primary_key" db:"id"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

type ModelUpper struct {
	Id      uint   `db:"id,omitempty"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func NewModel() *Model {
	m := new(Model)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

var (
	OrmMulti   int
	OrmMaxIdle int
	OrmMaxConn int
	OrmSource  string
	DebugMode  bool
)

// Convert ORMSource to DSN (dburl)
func ConvertSourceToDSN() string {
	template := "postgres://$(user):$(password)@$(host):5432/$(dbname)"

	// Parse one-by-one instead of using REGEX because of performance issues
	for _, option := range strings.Split(OrmSource, " ") {
		k := strings.Split(option, "=")[0]
		v := strings.Split(option, "=")[1]

		if strings.Contains(template, "$("+k+")") {
			template = strings.ReplaceAll(template, "$("+k+")", v)
		} else {
			template += "?" + option
		}
	}

	return template
}

func SplitSource() map[string]string {
	options := make(map[string]string)
	// Split one-by-one instead of using REGEX because of performance issues
	for _, option := range strings.Split(OrmSource, " ") {
		k := strings.Split(option, "=")[0]
		v := strings.Split(option, "=")[1]

		options[k] = v
	}

	return options
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func wrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
	b.ResetTimer()
}

func initDB() {
	sqls := [][]string{
		{
			`DROP TABLE IF EXISTS models;`,
			`CREATE TABLE models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		},
		{
			`DROP TABLE IF EXISTS models_upper;`,
			`CREATE TABLE models_upper (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_upper_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		},
		{
			`DROP TABLE IF EXISTS models_godb;`,
			`CREATE TABLE models_godb (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_godb_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		},
		{
			`DROP TABLE IF EXISTS pop_models;`,
			`CREATE TABLE pop_models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT pop_models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		},
		{
			`DROP TABLE IF EXISTS reform_models;`,
			`CREATE TABLE reform_models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT reform_models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		},
	}

	DB, err := sql.Open("postgres", OrmSource)
	checkErr(err)
	defer func() {
		err := DB.Close()
		checkErr(err)
	}()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		for _, line := range sql {
			_, err = DB.Exec(line)
			checkErr(err)
		}
	}

}
