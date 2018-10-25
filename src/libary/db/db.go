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
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",dbConfig.DB_USER,dbConfig.DB_PASSWORD,dbConfig.DB_HOST,dbConfig.DB_PORT, dbConfig.DB_DATABASE_NAME,dbConfig.DB_CHARSET)
	DB,err = sql.Open(dbConfig.DB_DRIVER_NAME, dataSourceName)
	if err != nil{
          logger.Error.Fatal("mysql connect error",err)
	}
}

func Insert(sql string)(id int64, err error) {
	logger.Info.Println(sql)
	result, e := DB.Exec(sql)
	if e != nil{
		logger.Error.Println(sql,"insert data error:",e)
		logger.Info.Println("isertid:0","error:",err)
        return 0,e
	}
	id, err = result.LastInsertId()
	logger.Info.Println("isertid:",id,"error:",err)
	return
}

//返回单行
func GetRow(sql string) (r map[string]string, err error)  {
	logger.Info.Println(sql)
	rows, e := DB.Query(sql)
	record := make(map[string]string)
	if e != nil{
		logger.Info.Println("querydata:",e)
		return record, e
	}
	col, _ := rows.Columns()
	scanArgs := make([]interface{},len(col))
	valArgs  := make([]interface{},len(col))
	for k:= range valArgs{
		scanArgs[k] = &valArgs[k]
	}
	for rows.Next(){
		rows.Scan(scanArgs...)
		for i,cc := range valArgs{
			if cc != nil{
				record[col[i]] = string(cc.([]byte))
			}
		}
	}
	logger.Info.Println(record)
    return record,nil
}


func UpdateSql(sql string) (count int64, err error) {
	logger.Info.Println(sql)
	result, e := DB.Exec(sql)
	if e != nil{
		logger.Error.Println(sql,"update data error:",e)
		logger.Info.Println("update count:0","error:",err)
		return 0,e
	}
	count, err = result.RowsAffected()
	logger.Info.Println("num:",count,"error:",err)
	return
}



