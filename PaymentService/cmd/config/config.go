package config

import(
	"evn"
)

type DatabaseConfig struct{
	ConnUrl string
	MaxOpenTime int
	MaxIdealConn int
	MaxIdealTime int
}

var Dbname = evn.GetString("DATABASE_NAME")
var CollectionName = evn.GetString("COLLECTION_NAME")



