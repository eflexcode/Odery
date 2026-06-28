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

var Dbname string = evn.GetString("DATABASE_NAME")
var CollectionName string = evn.GetString("COLLECTION_NAME")



