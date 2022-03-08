package must

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"zhanyia/src/common"
)

type Mysql struct {
	Db *sql.DB
}

// 创建redis组件实例
func init() {
	common.AllGlobal["Mysql"] = &Mysql{}

	dbTemp, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/board_games?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	dbTemp.SetConnMaxLifetime(600 * time.Second)
	dbTemp.SetMaxIdleConns(0)
	common.AllGlobal["Mysql"].(*Mysql).Db = dbTemp
}

func (m *Mysql) Query() {
	data, err := m.Db.Exec("update config_games set value=? where auto_id=?", 2, 1)
	if err != nil {
		fmt.Println("Mysql Query has err", err)
	}
	fmt.Println(data.RowsAffected())
	fmt.Println(data.LastInsertId())
}
