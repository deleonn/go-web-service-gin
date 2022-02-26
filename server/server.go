package server

import "web-service-gin/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run("localhost:" + config.GetString("server.port"))
}
