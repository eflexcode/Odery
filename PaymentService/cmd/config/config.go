package config

import(
	"github.com/cmd/evn"
)

type DatabaseConfig struct{
	ConnUrl string
	MaxOpenTime int
	MaxIdealConn int
	MaxIdealTime int
}

var Dbname string = evn.GetString("DATABASE_NAME","OrderyPayments")
var CollectionNameCards string = evn.GetString("COLLECTION_NAME","cards")
var CollectionNamePayments string = evn.GetString("COLLECTION_NAME_PAYMENTS","payments")
