package db

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

type Mysql struct {
	User   string
	Passwd string
	Host   string
	Port   int
	DB     string
}

func NewMysql(host string, port int, user string, passwd string, db string) *Mysql {
	return &Mysql{user, passwd, host, port, db}
}

func (m *Mysql) Open() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", m.User, m.Passwd, m.Host, m.Port, m.DB))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (m *Mysql) Query(dest interface{}, query string) error {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", m.User, m.Passwd, m.Host, m.Port, m.DB))
	if err != nil {
		return err
	}

	if reflectx.Deref(reflect.TypeOf(dest)).Kind() == reflect.Slice {
		err = db.Select(dest, query)
	} else {
		err = db.Get(dest, query)
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (m *Mysql) Exec(query string) (int64, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", m.User, m.Passwd, m.Host, m.Port, m.DB))
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	nRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return nRows, nil
}

// func (m *Mysql) QueryBy(conn *sql.DB, dest interface{}, query string) error {
// 	db := sqlx.NewDb(conn, "mysql")

// 	var err error
// 	if reflectx.Deref(reflect.TypeOf(dest)).Kind() == reflect.Slice {
// 		err = db.Select(dest, query)
// 	} else {
// 		err = db.Get(dest, query)
// 	}

// 	if err != nil && err != sql.ErrNoRows {
// 		return err
// 	}

// 	return nil
// }

// func (m *Mysql) Select(dest interface{}, query string) error {
// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", m.User, m.Passwd, m.Host, m.Port, m.DB))
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		rows.Columns
// 		if err := rows.Scan(); err != nil {
// 			return err
// 		}
// 	}
// }
