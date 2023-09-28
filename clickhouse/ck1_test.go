package clickhouse_

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	//_ "github.com/ClickHouse/clickhouse-go/v2"
)

// 连接池的最大数量
const maxConnections = 10

func connect() (*sql.DB, error) {
	// 创建连接池
	pool, err := sql.Open("clickhouse", "tcp://192.168.0.150:9000?debug=true")
	if err != nil {
		return nil, err
	}

	// 设置连接池的最大连接数
	pool.SetMaxOpenConns(maxConnections)

	// 设置连接池中最多可以有多少空闲连接
	pool.SetMaxIdleConns(maxConnections / 2)

	// 设置连接的最大空闲时间
	pool.SetConnMaxLifetime(time.Minute * 5)

	// 测试连接是否可用
	if err = pool.Ping(); err != nil {
		return nil, err
	}

	return pool, nil
}

func TestAdd1(t *testing.T) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 构造插入语句
	stmt := `
    INSERT INTO mytable (id, name, age)
    VALUES
    `
	for i := 0; i < 1000; i++ {
		stmt += fmt.Sprintf("(%d, 'Name%d', %d),", i, i, i%100)
	}
	stmt = stmt[:len(stmt)-1]

	// 执行插入语句
	result, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	}

	// 输出插入结果
	fmt.Println(result.RowsAffected())
}
