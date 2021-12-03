package db

import (
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/l1huanyu/x1aol1system/infrastructure/db/ent"
	"github.com/l1huanyu/x1aol1system/util"
)

type Ent struct {
	client *ent.Client
}

func NewEnt() *Ent {
	client, err := ent.Open(dialect.MySQL, util.GetMySQLSource())
	if err != nil {
		panic(err)
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	return &Ent{
		client: client,
	}
}
