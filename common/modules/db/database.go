package db

import "github.com/topfreegames/pitaya/v2/interfaces"

type DataBase interface {
	interfaces.Module
	Connect()
	Close()
	TestPing() error
}
