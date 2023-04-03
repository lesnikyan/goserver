package server

type MysqlClient struct{}

type ServiceConf struct {
	Remote NetPoint
	Auth   AuthKey
}
