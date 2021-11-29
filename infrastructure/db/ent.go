package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/l1huanyu/x1aol1system/ent"
)

type Ent struct {
	client *ent.Client
}

func NewEnt() *Ent {
	client, err := ent.Open("mysql", "root:Li960127!@tcp(sh-cynosdbmysql-grp-8wz4ch3s.sql.tencentcdb.com:26230)/x1aol1system?parseTime=True")
	if err != nil {
		panic(err)
	}
	return &Ent{
		client: client,
	}
}
