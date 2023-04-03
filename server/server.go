package server

/*
DB: MySQL, PostgreSQL, MongoDB
Cache: Redis
Queue: RabbitMQ,
*/

type SBServer struct {
	PortGRPC int
	PortHTTP int
	Cache    RedisClient
	DB       MysqlClient
}
