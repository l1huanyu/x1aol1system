package client

import "sync"

type DB interface {
}

var (
	dbClient DB
	dbOnce   sync.Once
)

func InitDB(client DB) {
	dbOnce.Do(func() {
		dbClient = client
	})
}

func GetDB() DB {
	return dbClient
}
