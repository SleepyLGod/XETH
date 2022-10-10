package dbDriver

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"time"
)

const MaxOpenConnections = 10
const MaxConnectionLifetime = time.Minute * 3
const MaxIdleConnections = 10

type MysqlDriver struct {
	db *gorm.DB
	t  reflect.Type
}

func NewMysqlDriverFromDsn(dsn string, schema interface{}) (*MysqlDriver, error) {
	sqlDb, err := sql.Open("mysql", dsn+"?charset=utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxOpenConns(MaxOpenConnections)
	sqlDb.SetMaxIdleConns(MaxIdleConnections)
	sqlDb.SetConnMaxLifetime(MaxConnectionLifetime)
	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = gormDb.AutoMigrate(schema); err != nil {
		return nil, err
	}

	return &MysqlDriver{
		db: gormDb,
		t:  reflect.TypeOf(schema),
	}, nil
}

func NewMysqlDriver(usrName string, psw string, addr string, dbName string, schema interface{}) (*MysqlDriver, error) {
	dsn := usrName + ":" + psw + "@(" + addr + ")/" + dbName
	return NewMysqlDriverFromDsn(dsn, schema)
}

func (d *MysqlDriver) CloseDB() {
	if d.db == nil {
		fmt.Println("DB is NIL")
		return
	}
	sqlDb, err := d.db.DB()
	if err != nil {
		panic(err)
	}
	if err = sqlDb.Close(); err != nil {
		panic(err)
	}
}

func (d *MysqlDriver) Create(v interface{}) error {
	if t := reflect.TypeOf(v); t != d.t {
		return errors.New(fmt.Sprintf("unexpected type: %s, expecting: %s", t.Name(), d.t.Name()))
	}
	d.db.Create(v)
	return nil
}

func (d *MysqlDriver) QuerySingle(output interface{}, constraints string, params ...string) {
	d.db.First(output, constraints, params)
}
