package clickhouse_

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	source := "tcp://192.168.0.150:9090?debug=true&username=default&password=&database=default"
	connect, err := sqlx.Connect("clickhouse", source)
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
	}
	defer func() {
		_ = connect.Close()
	}()

	//数据预处理写入
	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into testhi.hi_test_user (id,name,age) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	for i := 10; i < 20; i++ {
		if _, err := stmt.Exec(i, "n"+strconv.Itoa(i), i+10); err != nil {
			log.Fatal(err)
		}
	}
	_ = tx.Commit()
}
