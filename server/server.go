package server

import "github.com/giofcosta/portfolio-golang-api/config"

//Init initializes the server
func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
