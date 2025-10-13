package main

import (
	"log"

	"github.com/metux/mpbt/core"
//	"github.com/metux/go-nebulon/webapi/servers"
)

func main() {
	_, err := core.LoadComponent("../cf/xlibre/components/system-libs/font-util.yaml")
	if err != nil {
		log.Println("err %s", err)
	}

//	server, err := servers.BootServer(flag_conffile, flag_serverid)
//	if err != nil {
//		panic(err)
//	}
//	server.Serve()
}
