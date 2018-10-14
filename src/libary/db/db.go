package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"logger"
	dbConfig "config/db"
)


var err error
var DB  *sql.DB


func init(){
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=%s",dbConfig.DB_USER,dbConfig.DB_PASSWORD,dbConfig.DB_DATABASE_NAME,dbConfig.DB_CHARSET)
	DB,err = sql.Open(dbConfig.DB_DRIVER_NAME, dataSourceName)
	if err != nil{
          logger.Error.Fatal("mysql connect error",err)
	}
}

func Insert(sql string)(id int64, err error) {
	result, e := DB.Exec(sql)
	if e != nil{
		logger.Info.Println("sql:",sql,"err:",e)
		logger.Error.Println("insert data error:",e)
        return 0,e
	}
	id, err = result.LastInsertId()
	return
}

