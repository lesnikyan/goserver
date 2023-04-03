package server

/*
1. Use grpc for client-server
2. use http for web
3. use RabbitMQ for request queue
4. use Mysql as long-term storage
5. use Redis as fast store and cache

Demo:
1. Listen requests
2. Get params from resuest
3. Try read Redis cache
4. Read DB
5. Generate result and put in cache
6. Respond
*/

func NewSBServer(
	grpc NetPoint,
	http NetPoint,
	msql ServiceConf,
	redis ServiceConf) *SBServer {
	server := &SBServer{}
	return server
}
