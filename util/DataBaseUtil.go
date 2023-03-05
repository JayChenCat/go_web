package util

import (
	//驱动需要进行隐式导入
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//获取数据库连接对象-mysql
func GetConnection() (*sql.DB, error) {
	dataSourceName := userName + ":" + password + tcpPort + dbName + "?" + charset + "&" + local + "&" + parseTime
	conn, err := sql.Open(driverName, dataSourceName)
	return conn, err
}

//单个物理删除
func Delete(SQL string, Id int64) error {
	return deleteOper(SQL, Id)
}

//单个逻辑删除
func FalseDelete(SQL string, Id int64) error {
	return deleteOper(SQL, Id)
}

func deleteOper(SQL string, Id int64) error {
	conn, err := GetConnection()
	//关闭数据连接
	defer Close(conn, err)
	if err != nil {
		return err
	}
	var r sql.Result
	r, err = conn.Exec(SQL,
		Id, //编号
	)
	if err != nil {
		return err
	}
	rowNum, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowNum > 0 {
		return nil
	}
	return err
}

//事务 提交或修改
func ExecTransaction(callback func(tx *sql.Tx) error) error {
	conn, err := GetConnection()
	//关闭数据连接
	defer Close(conn, err)
	if err != nil {
		return err
	}
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	err = callback(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

//关闭数据连接对象
func Close(conn *sql.DB, err error) {
	//关闭数据连接
	if conn != nil {
		_ = conn.Close()
	}
}

//批量插入sql或批量修改(拼接语句实现)
func BatchInsertOrUpdate(stmt string, valueArgs []interface{}) error {
	conn, err := GetConnection()
	//关闭数据连接
	defer Close(conn, err)
	if err != nil {
		return err
	}
	var r sql.Result
	r, err = conn.Exec(stmt, valueArgs...)
	if err != nil {
		return err
	}
	rowNum, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowNum > 0 {
		return nil
	}
	return err
}
