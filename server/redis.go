package server

type RedisClient struct {
	Server   string
	Port     int
	User     string
	Password string
}

func NewRedis(serverAddress NetPoint, auth AuthKey) *RedisClient {
	red := &RedisClient{}

	return red
}
