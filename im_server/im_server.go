package main

type serverConfig struct {
	socket_io_address string
	tls_address string
	cert_file string
	key_file string
}

var config serverConfig
var server_summary *ServerSummary


func init() {
	//app_route = NewAppRoute()
	server_summary = NewServerSummary()
	//sync_c = make(chan *SyncHistory, 100)
	//group_sync_c = make(chan *SyncGroupHistory, 100)
}

func main() {
	//config.socket_io_address = ":18888"
	//StartSocketIO(config.socket_io_address, config.tls_address,
		//config.cert_file, config.key_file)

	ListenClient()
}




