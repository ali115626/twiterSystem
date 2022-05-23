package Dao

import (
	"database/sql"
	"fmt"
)

var mySqlDb *sql.DB

var MysqlConn *sql.DB


func MysqlInit()*sql.DB{
	var err error
	mySqlDb, err = sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//return nil,nil,errors.New(fmt.Sprintf("open database error,err=", err))
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}


	mySqlDb.SetMaxOpenConns(30)
	mySqlDb.SetMaxIdleConns(5)
	return mySqlDb
}
