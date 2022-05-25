package Dao

import (
	"database/sql"
	"fmt"
	"log"
)

//var MySqlDb *sql.DB

var MysqlConn *sql.DB


var MySqlDb=MysqlInit()

func MysqlInit()*sql.DB{
	var err error
	MySqlDb1, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//return nil,nil,errors.New(fmt.Sprintf("open database error,err=", err))
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}


	MySqlDb1.SetMaxOpenConns(20)
	MySqlDb1.SetMaxIdleConns(5)
	return MySqlDb1
}


var superDB *sql.DB
var superDataBase = "root:Aa123456@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true&timeout=1s&readTimeout=6s"

var DB *sql.DB
var dataBase = "orchestrator:orch_backend_password@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true&timeout=1s&readTimeout=6s"

func Init(dataBase string) *sql.DB {
	db, err := sql.Open("mysql", dataBase)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	// max idle
	db.SetMaxIdleConns(1000)

	err = db.Ping()
	if err != nil {
		log.Fatalln("ping db fail:", err)
	}

	return db
}

func main() {
	superDB = Init(superDataBase)
	DB = Init(dataBase)
}
