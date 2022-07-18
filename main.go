package main

import (
	"flag"
	"github.com/benjankowski/raspio/server"
)

func main() {
    enable_server := flag.Bool("server", false, "Enables server mode (for the controlled device)")
    server_bind := flag.String("bind", "127.0.0.1:3030", "The IP and Port to bind to")
    server_key := flag.String("key", "", "Server authentication key")
    
    flag.Parse()

    if *enable_server {
        server.Start(*server_bind, *server_key)
    }
}
