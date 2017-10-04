package db

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	Master *Mysql
)

// Mysql 包含连接数据库必备的参数，后续开发连接池
type Mysql struct {
	User   string
	Passwd string
	Host   string
	Port   int
	DB     string
}

// Init 初始化mysql模块
func Init() {
	fmt.Println("[modeul mysql]")
	Master = NewMysql("127.0.0.1", 3306, "root", "123456", "auth")

	fmt.Print("test connection to mysql... ")
	db, err := Master.Open()
	if err != nil {
		fmt.Println("[FAIL]")
		panic(err)
	}
	fmt.Println("[OK]")
	db.Close()
}

// NewMysql 获得结构体
func NewMysql(host string, port int, user string, passwd string, db string) *Mysql {
	return &Mysql{user, passwd, host, port, db}
}

// Open 连接数据库，原生连接
func (m *Mysql) Open() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", m.User, m.Passwd, m.Host, m.Port, m.DB))
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec("SET collation_connection = utf8mb4_unicode_ci")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Get 使用sqlx扩展以将数据传入设定的结构体切片，用于查询
func (m *Mysql) Get(dest interface{}, query string) error {
	conn, err := m.Open()
	if err != nil {
		return err
	}
	db := sqlx.NewDb(conn, "mysql")
	defer db.Close()

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

// Set 使用原生sql操作，用于插入、更新
func (m *Mysql) Set(query string) (int64, error) {
	db, err := m.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()

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
