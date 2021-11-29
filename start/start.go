package start

import (
	"github.com/l1huanyu/x1aol1system/client"
	"github.com/l1huanyu/x1aol1system/infrastructure/db"
)

func Init() {
	client.InitDB(db.NewEnt())
}
