package dbDriver

import (
	"XETH/config"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const MaxOpenConnections = 10
const MaxConnectionLifetime = time.Minute * 3
const MaxIdleConnections = 10

type MysqlDriver struct {
	db *gorm.DB
	t  reflect.Type
}

// Db 全局变量（只建立一个数据库）
var Db *gorm.DB

// GetDB getter
func GetDB() *gorm.DB {
	return Db
}

func NewMysqlDriverFromDsnWithoutInterface(dsn string) {
	var err error
	Db, err = gorm.Open(mysql.Open(dsn+"?charset=utf8&parseTime=true"), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		panic("Connecting database failed: " + err.Error())
	}
}

func NewMysqlDriverWithoutInterface() {
	databaseLogin := config.DatabaseGenerator
	dsn := databaseLogin().UsrName + ":" + databaseLogin().Psw + "@(" + databaseLogin().Addr + ")/" + databaseLogin().DbName
	NewMysqlDriverFromDsnWithoutInterface(dsn)
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

//func (d *MysqlDriver) QuerySingle(output interface{}, constraint string, param ...string) {
//	d.db.First(output, constraint, param)
//}
//
//func (d *MysqlDriver) QueryMulti(output interface{}, constraint string, param ...string) {
//	d.db.Find(output, constraint, param)
//}

type QueryConstraint struct {
	FieldName string `json:"fieldName"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

func QueryWithDb(db *gorm.DB, output interface{}, constraints []QueryConstraint) {
	if len(constraints) < 1 {
		db.Find(output)
	} else {
		tx := db.Where(constraints[0].FieldName+" "+constraints[0].Operator+" ? ", constraints[0].Value)
		size := len(constraints)
		for i := 1; i < size; i++ {
			tx = tx.Where(constraints[i].FieldName+" "+constraints[i].Operator+" ? ", constraints[i].Value)
		}
		tx.Find(output)
	}
}

func (d *MysqlDriver) Query(output interface{}, constraints []QueryConstraint) {
	if len(constraints) < 1 {
		d.db.Find(output)
	} else {
		tx := d.db.Where(constraints[0].FieldName+" "+constraints[0].Operator+" ? ", constraints[0].Value)
		size := len(constraints)
		for i := 1; i < size; i++ {
			tx = tx.Where(constraints[i].FieldName+" "+constraints[i].Operator+" ? ", constraints[i].Value)
		}
		tx.Find(output)
	}
}

// // testings
// u := &User{Name: "333"}
// if err = driver.Create(u); err != nil {
// 	fmt.Println(err)
// }
// var u2 []User
// con := make([]dbDriver.QueryConstraint, 0)

// con = append(con, dbDriver.QueryConstraint {
// 	FieldName: "name",
// 	Operator:  ">",
// 	Value:     "1",
// })
// con = append(con, dbDriver.QueryConstraint {
// 	FieldName: "name",
// 	Operator:  "!=",
// 	Value:     "test",
// })
// driver.Query(&u2, con)
// fmt.Println(len(u2))
// driver.CloseDB()
