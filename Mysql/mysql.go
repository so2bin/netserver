package Mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"reflect"
)

type Mysql struct {
	username string
	passwd   string
	dbname   string
	db       *sql.DB
}

func NewMysql(un string, pswd string, dbn string) (rsql *Mysql, err error) {
	rsql = &Mysql{username: un, passwd: pswd, dbname: dbn}
	var instr string
	instr = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4,utf8", un, pswd, dbn)
	rsql.db, err = sql.Open("mysql", instr) //第一个参数为驱动名
	return rsql, err
}

func (qry *Mysql) Query(qstr string) (qre *sql.Rows) {
	qre, err := qry.db.Query(qstr)
	if err != nil {
		fmt.Printf("query string(%s) wrong, err:%s!\n", qstr, err)
	}
	return
}

func (qry *Mysql) Close() {
	qry.db.Close()
}

func (qur *Mysql) PrintReslut(query *sql.Rows) {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	for k, v := range results { //查询出来的数组
		fmt.Println(k, v)
	}
}
