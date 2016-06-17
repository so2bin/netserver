package Mysql

import (
	"fmt"
	"testing"
)

func test(t *testing.T) {
	/*DSN数据源名称
	  [username[:password]@][protocol[(address)]]/dbname[?param1=value1¶mN=valueN]
	  user@unix(/path/to/socket)/dbname
	  user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
	  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	  user:password@/dbname
	  无数据库: user:password@/
	*/
	dbn := "hbbdb"
	dbtbl := "hbbdata"
	Q, err := NewMysql("hbb", "hbb112358", "")
	//db, err := sql.Open("mysql", "hbb:hbb112358@/hbbdata") //第一个参数为驱动名
	checkErr(err)
	query := Q.Query("show databases")
	Q.PrintReslut(query)
	Q.Query(fmt.Sprintf("drop database if exists %s", dbn))
	Q.Query(fmt.Sprintf("create database  if not exists %s default charset utf8 collate utf8_general_ci", dbn))
	Q.Query(fmt.Sprintf("create table %s.%s(c1 int, c2 varchar(20), c3 varchar(20))", dbn, dbtbl))
	Q.Query(fmt.Sprintf("insert into %s.%s values(101, '贺彬彬1', 'address1'), (102, '贺彬彬2', 'address2'), (103, 'name3', 'address3'), (104, 'name4', 'address4')", dbn, dbtbl))
	query = Q.Query(fmt.Sprintf("select 101 from  %s.%s", dbn, dbtbl))
	Q.PrintReslut(query)
	Q.Close()
}

func checkErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
	}
}
