package config

type DatabaseConfig struct{
	ConnUrl string
	MaxOpenTime int
	MaxIdealConn int
	MaxIdealTime int
}
