package utils

import (
	"database/sql"
	"errors"
	. "fmt"
	_ "github.com/mattn/go-sqlite3"
)


type DataBase struct {
	driverName string
	dataSourceName string
}

func NewSqlite(dataSourceName string) *DataBase {
	driverName := "sqlite3"
	return &DataBase{driverName,dataSourceName}
}

func (database *DataBase) OpenSource() (*sql.DB,error) {
	//db, err := sql.Open("sqlite3", "./foo.db")
	db, err := sql.Open(database.driverName, database.dataSourceName)
	if err != nil {
		return &sql.DB{},errors.New(Sprintf("OpenSource Failed : %s",err))
	}
	return db,nil
}

func (database *DataBase) CreateDataTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE `sys_info` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,`CPU_Temperature` float,`CPU_Use` float,`RAM_Used` float,`strftime` DATA);")
	if err != nil {
		return errors.New(Sprintf("CreateDataTable Failed : %s",err))
	}
	return nil
}

func (database *DataBase) InsertData(db *sql.DB,temp,used,ram float64) int64 {
	stmt, err := db.Prepare("INSERT INTO sys_info(CPU_Temperature,CPU_Use,RAM_Used,strftime) VALUES (ROUND(?, 2),ROUND(?, 2),ROUND(?, 2),datetime(CURRENT_TIMESTAMP,'localtime'));")
	checkErr(err)
	res, err := stmt.Exec(temp, used, ram)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}

func (database *DataBase) QueryData(db *sql.DB,SQL string) []map[string]interface{} {
	rows, err := db.Query(SQL)
	checkErr(err)

	var res []map[string]interface{}
	for rows.Next() {
		var id int64
		var temp float64
		var used float64
		var ram float64
		var strftime string
		err = rows.Scan(&id, &temp, &used, &ram,&strftime)
		m := map[string]interface{}{}
		m["id"]=id
		m["temp"]=temp
		m["used"]=used
		m["ram"]=ram
		m["strftime"]=strftime
		checkErr(err)
		res=append(res,m)
	}
	return res
}

func (database *DataBase) UpdateData(db *sql.DB,id int64,temp,used,ram float64) int64 {
	stmt, err := db.Prepare("UPDATE sys_info SET CPU_Temperature=ROUND(?, 2),CPU_Use=ROUND(?, 2),RAM_Used=ROUND(?, 2),strftime=datetime(CURRENT_TIMESTAMP,'localtime') WHERE id=?;")
	checkErr(err)

	res, err := stmt.Exec(temp, used, ram,id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}

func (database *DataBase) DeleteData(db *sql.DB,id int64) int64 {
	stmt, err := db.Prepare("DELETE FROM sys_info WHERE id=?;")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}

func (database *DataBase) DeleteAllData(db *sql.DB) int64 {
	stmt, err := db.Prepare("DELETE FROM sys_info;")
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}

func (database *DataBase) CloseConn(db *sql.DB) {
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}